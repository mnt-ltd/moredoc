package model

import (
	"bytes"
	"fmt"
	"moredoc/util/sitemap"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	isCreatingSitemap bool
)

// UpdateSitemap 更新站点地图
func (m *DBModel) UpdateSitemap() (err error) {
	if isCreatingSitemap {
		return
	}
	isCreatingSitemap = true
	defer func() {
		isCreatingSitemap = false
	}()
	os.MkdirAll("sitemap", os.ModePerm)

	var (
		limit          = 10000
		page           = 1
		documents      []Document
		articles       []Article
		modelDocument  = &Document{}
		modelArticle   = &Article{}
		sitemapIndexes []sitemap.SitemapIndex
		sm             = sitemap.NewSitemap()
		domain         = strings.TrimRight(m.GetConfigOfSystem(ConfigSystemDomain).Domain, "/")
	)
	for {
		if err = m.db.Model(modelDocument).Select("id", "updated_at").Limit(limit).Offset((page - 1) * limit).Order("id asc").Find(&documents).Error; err != nil && err != gorm.ErrRecordNotFound {
			m.logger.Error("execUpdateSitemap", zap.Error(err))
			return
		}
		if len(documents) == 0 {
			break
		}
		file := fmt.Sprintf("sitemap/documents-%d.xml", page)
		var su []sitemap.SitemapUrl
		for _, doc := range documents {
			su = append(su, sitemap.SitemapUrl{
				Loc:        fmt.Sprintf("%s/document/%d", domain, doc.Id),
				Lastmod:    doc.UpdatedAt.Format("2006-01-02 15:04:05"),
				ChangeFreq: sitemap.DAILY,
				Priority:   1.0,
			})
		}
		if err = sm.CreateSitemapContent(su, file); err != nil {
			m.logger.Error("execUpdateSitemap", zap.Error(err))
			return
		}
		sitemapIndexes = append(sitemapIndexes, sitemap.SitemapIndex{
			Loc:     domain + "/" + file,
			Lastmod: time.Now().Format("2006-01-02 15:04:05"),
		})
		page++
	}

	page = 1
	for {
		if err = m.db.Model(modelArticle).Select("id", "updated_at", "identifier").Limit(limit).Offset((page - 1) * limit).Order("id asc").Find(&articles).Error; err != nil && err != gorm.ErrRecordNotFound {
			m.logger.Error("execUpdateSitemap", zap.Error(err))
			return
		}
		if len(articles) == 0 {
			break
		}
		file := fmt.Sprintf("sitemap/articles-%d.xml", page)
		var su []sitemap.SitemapUrl
		for _, article := range articles {
			su = append(su, sitemap.SitemapUrl{
				Loc:        fmt.Sprintf("%s/article/%s", domain, article.Identifier),
				Lastmod:    article.UpdatedAt.Format("2006-01-02 15:04:05"),
				ChangeFreq: sitemap.DAILY,
				Priority:   1.0,
			})
		}
		if err = sm.CreateSitemapContent(su, file); err != nil {
			m.logger.Error("execUpdateSitemap", zap.Error(err))
			return
		}
		sitemapIndexes = append(sitemapIndexes, sitemap.SitemapIndex{
			Loc:     domain + "/" + file,
			Lastmod: time.Now().Format("2006-01-02 15:04:05"),
		})
		page++
	}

	if len(sitemapIndexes) > 0 {
		if err = sm.CreateSitemapIndex(sitemapIndexes, "sitemap/sitemap.xml"); err != nil {
			m.logger.Error("execUpdateSitemap", zap.Error(err))
			return
		}
	}

	return
}

// SEO
func (m *DBModel) InitSEO() {
	// 扫描dist目录下的所有HTML文件，将文件名作为SEO的关键字
	cfg := m.GetConfigOfSystem()
	dist := "dist"
	pages := map[string]string{
		"200.html":                "",
		"404.html":                "404 - 页面未找到 - ",
		"findpassword/index.html": "找回密码 - ",
		"index.html":              "",
		"login/index.html":        "用户登录 - ",
		"register/index.html":     "用户注册 - ",
		"search/index.html":       "文档搜索 - ",
		"upload/index.html":       "文档上传 - ",
	}
	filepath.Walk(dist, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		path = filepath.ToSlash(path)
		if filepath.Ext(path) == ".html" {
			name := strings.TrimPrefix(path, dist+"/")
			defaultTitle, ok := pages[name]
			if !ok && strings.HasPrefix(path, dist+"/admin") {
				defaultTitle = "管理后台 - "
			}

			m.logger.Debug("initSEO", zap.String("file", path), zap.String("title", defaultTitle))
			bs, _ := os.ReadFile(path)
			if doc, errDoc := goquery.NewDocumentFromReader(bytes.NewReader(bs)); errDoc != nil {
				m.logger.Error("initSEO", zap.Error(errDoc), zap.String("file", path))
			} else {
				doc.Find("title").SetText(defaultTitle + cfg.Sitename)
				doc.Find("meta[name='keywords']").SetAttr("content", cfg.Keywords)
				doc.Find("meta[name='description']").SetAttr("content", cfg.Description)
				doc.Find("meta[content='moredoc']").Remove()
				doc.Find("meta[name='og:type']").Remove()
				if htmlStr, errHtml := doc.Html(); errHtml == nil {
					os.WriteFile(path, []byte(htmlStr), os.ModePerm)
				}
			}
		}
		return nil
	})
}

func (m *DBModel) cronUpdateSitemap() {
	layout := "2006-01-02"
	lastUpdated := time.Now().Format(layout)
	for {
		hour, _ := strconv.Atoi(os.Getenv("MOREDOC_UPDATE_SITEMAP_HOUR")) // 默认为每天凌晨0点更新站点地图
		hour = hour % 24
		m.logger.Info("cronUpdateSitemap", zap.Int("hour", hour), zap.String("lastUpdated", lastUpdated))
		now := time.Now()
		if now.Hour() == hour && now.Format(layout) != lastUpdated {
			m.logger.Info("cronUpdateSitemap，start...")
			err := m.UpdateSitemap()
			if err != nil {
				m.logger.Info("cronUpdateSitemap，end...", zap.Error(err))
			}
			m.logger.Info("cronUpdateSitemap，end...")
			lastUpdated = now.Format(layout)
		}
		time.Sleep(1 * time.Minute)
	}
}

// 清理无效附件
// 1. 找出已被标记删除的附件
// 2. 查询是否存在相同hash的未被标记删除的附件，对于此类附件，则只删除附件记录而不删除附件文件。
// 3. 删除已被标记删除的附件
// 4. 对于文档类附件，要注意衍生的附件，如缩略图、PDF等，也要一并删除。
func (m *DBModel) cronCleanInvalidAttachment() {
	sleepDuration := 1 * time.Minute
	for {
		time.Sleep(1 * time.Second)
		m.logger.Info("cronCleanInvalidAttachment，start...")
		var (
			deletedAttachemnts, attachemnts []Attachment
			hashes                          []string
			hashMap                         = make(map[string]struct{})
			ids                             []int64
			retentionMinute                 = m.GetConfigOfSecurity(ConfigSecurityAttachmentRetentionMinute).AttachmentRetentionMinute
		)

		if retentionMinute < 0 {
			retentionMinute = 0
		}

		// 1. 找出已被标记删除的附件
		m.db.Unscoped().Where("deleted_at IS NOT NULL").Where("deleted_at < ?", time.Now().Add(-time.Duration(retentionMinute)*time.Minute)).Limit(100).Find(&deletedAttachemnts)
		if len(deletedAttachemnts) == 0 {
			m.logger.Info("cronCleanInvalidAttachment，end...")
			time.Sleep(sleepDuration)
			continue
		}

		for _, attachemnt := range deletedAttachemnts {
			hashes = append(hashes, attachemnt.Hash)
			ids = append(ids, attachemnt.Id)
		}

		// 2. 查询是否存在相同hash的未被标记删除的附件
		m.db.Select("hash").Where("hash IN (?)", hashes).Group("hash").Limit(len(hashes)).Find(&attachemnts)
		for _, attachemnt := range attachemnts {
			hashMap[attachemnt.Hash] = struct{}{}
		}

		// 3. 删除已被标记删除的附件
		err := m.db.Unscoped().Where("id IN (?)", ids).Delete(&Attachment{}).Error
		if err != nil {
			m.logger.Error("cronCleanInvalidAttachment", zap.Error(err))
			m.logger.Info("cronCleanInvalidAttachment，end...")
			continue
		}
		m.logger.Info("cronCleanInvalidAttachment", zap.Any("ids", ids), zap.Any("Attachemnts", deletedAttachemnts))
		for _, attachemnt := range deletedAttachemnts {
			if _, ok := hashMap[attachemnt.Hash]; !ok { // 删除附件文件
				m.logger.Debug("cronCleanInvalidAttachment", zap.String("path", attachemnt.Path), zap.Any("attachemnt", attachemnt))
				file := strings.TrimLeft(attachemnt.Path, "./")
				m.logger.Debug("cronCleanInvalidAttachment", zap.String("file", file))
				if err := os.Remove(file); err != nil {
					m.logger.Error("cronCleanInvalidAttachment", zap.Error(err), zap.String("file", file))
				}
				if attachemnt.Type == AttachmentTypeDocument { // 删除文档的衍生文件
					folder := strings.TrimSuffix(file, filepath.Ext(file))
					m.logger.Debug("cronCleanInvalidAttachment", zap.String("folder", folder))
					if err := os.RemoveAll(folder); err != nil {
						m.logger.Error("cronCleanInvalidAttachment", zap.Error(err), zap.String("folder", folder))
					}
				}
			}
		}
		m.logger.Info("cronCleanInvalidAttachment，end...")
	}
}

func (m *DBModel) cronMarkAttachmentDeleted() {
	// 定时标记删除24小时前上传的但是未被使用的附件
	for {
		time.Sleep(1 * time.Hour)
		var (
			configs []Config
			banners []Banner
			hashes  []string
		)

		// 1. 查找图片类配置
		m.db.Select("value").Where("input_type = ?", "image").Find(&configs)
		if len(configs) > 0 {
			for _, config := range configs {
				// 文件hash
				hash := strings.TrimSpace(strings.TrimSuffix(filepath.Base(config.Value), filepath.Ext(config.Value)))
				if hash != "" {
					hashes = append(hashes, hash)
				}
			}

		}

		// 2. 查找横幅类配置
		m.db.Select("path").Find(&banners)
		if len(banners) > 0 {
			for _, banner := range banners {
				// 文件hash
				hash := strings.TrimSpace(strings.TrimSuffix(filepath.Base(banner.Path), filepath.Ext(banner.Path)))
				if hash != "" {
					hashes = append(hashes, hash)
				}
			}
		}

		if len(hashes) > 0 {
			err := m.db.Where("`hash` NOT IN (?) and `type` in (?)", hashes, []int{AttachmentTypeConfig, AttachmentTypeBanner}).Delete(&Attachment{}).Error
			if err != nil {
				m.logger.Error("cronMarkAttachmentDeleted", zap.Error(err))
			}
		}

		// 非配置类和横幅类附件，如果type_id为0，则表示未被使用，超过24小时则标记删除
		m.logger.Info("cronMarkAttachmentDeleted start...")
		err := m.db.Where("`type` not in (?)  and type_id = ?", []int{AttachmentTypeConfig, AttachmentTypeBanner}, 0).Where("created_at < ?", time.Now().Add(-time.Duration(24)*time.Hour)).Delete(&Attachment{}).Error
		if err != nil {
			m.logger.Error("cronMarkAttachmentDeleted", zap.Error(err))
		}
		m.logger.Info("cronMarkAttachmentDeleted end...")
	}
}

func (m *DBModel) loopCovertDocument() {
	if convertDocumentRunning {
		return
	}
	convertDocumentRunning = true
	sleep := 10 * time.Second
	m.db.Model(&Document{}).Where("status = ?", DocumentStatusConverting).Update("status", DocumentStatusPending)
	for {
		now := time.Now()
		m.logger.Info("loopCovertDocument，start...")
		err := m.ConvertDocument()
		if err != nil && err != gorm.ErrRecordNotFound {
			m.logger.Info("loopCovertDocument，end...", zap.Error(err), zap.String("cost", time.Since(now).String()))
		}
		if err == gorm.ErrRecordNotFound {
			time.Sleep(sleep)
		}
	}
}

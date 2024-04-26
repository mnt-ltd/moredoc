package model

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"moredoc/util"
	"moredoc/util/converter"
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

type reconvertDocument struct {
	Id int64 `json:"id"`
}

var (
	isCreatingSitemap bool
	cacheReconvert    = "cache/reconvert"
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
		if err = m.db.Model(modelDocument).Select("id", "updated_at", "uuid").Limit(limit).Offset((page - 1) * limit).Order("id asc").Find(&documents).Error; err != nil && err != gorm.ErrRecordNotFound {
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
				// Loc:        fmt.Sprintf("%s/document/%d", domain, doc.Id),
				Loc:        fmt.Sprintf("%s/document/%s", domain, doc.UUID),
				Lastmod:    doc.UpdatedAt.Format(time.RFC3339),
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
			Lastmod: time.Now().Format(time.RFC3339),
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
				Lastmod:    article.UpdatedAt.Format(time.RFC3339),
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
			Lastmod: time.Now().Format(time.RFC3339),
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

			bs, _ := os.ReadFile(path)
			if doc, errDoc := goquery.NewDocumentFromReader(bytes.NewReader(bs)); errDoc != nil {
				m.logger.Error("initSEO", zap.Error(errDoc), zap.String("file", path))
			} else {
				m.logger.Debug("initSEO", zap.String("file", path), zap.String("title", defaultTitle+cfg.Sitename))
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
		m.logger.Debug("cronUpdateSitemap", zap.Int("hour", hour), zap.String("lastUpdated", lastUpdated))
		now := time.Now()
		if now.Hour() == hour && now.Format(layout) != lastUpdated {
			m.logger.Debug("cronUpdateSitemap，start...")
			err := m.UpdateSitemap()
			if err != nil {
				m.logger.Debug("cronUpdateSitemap，end...", zap.Error(err))
			}
			m.logger.Debug("cronUpdateSitemap，end...")
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
		m.logger.Debug("cronCleanInvalidAttachment，start...")
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
			m.logger.Debug("cronCleanInvalidAttachment，end...")
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
			m.logger.Debug("cronCleanInvalidAttachment，end...")
			continue
		}
		m.logger.Debug("cronCleanInvalidAttachment", zap.Any("ids", ids), zap.Any("Attachemnts", deletedAttachemnts))
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
		m.logger.Debug("cronCleanInvalidAttachment，end...")
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

		// 2. 查找轮播图类配置
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

		// 非配置类和轮播图类附件，如果type_id为0，则表示未被使用，超过24小时则标记删除
		m.logger.Debug("cronMarkAttachmentDeleted start...")
		err := m.db.Where("`type` not in (?)  and type_id = ?", []int{AttachmentTypeConfig, AttachmentTypeBanner}, 0).Where("created_at < ?", time.Now().Add(-time.Duration(24)*time.Hour)).Delete(&Attachment{}).Error
		if err != nil {
			m.logger.Error("cronMarkAttachmentDeleted", zap.Error(err))
		}
		m.logger.Debug("cronMarkAttachmentDeleted end...")
	}
}

func (m *DBModel) loopCovertDocument() {
	if convertDocumentRunning {
		return
	}
	// 清空缓存目录
	os.RemoveAll("cache/convert")
	convertDocumentRunning = true
	sleep := 10 * time.Second
	m.db.Model(&Document{}).Where("status = ?", DocumentStatusConverting).Update("status", DocumentStatusPending)
	for {
		now := time.Now()
		m.logger.Debug("loopCovertDocument，start...")
		err := m.ConvertDocument()
		if err != nil && err != gorm.ErrRecordNotFound {
			m.logger.Error("loopCovertDocument", zap.Error(err))
		}
		m.logger.Debug("loopCovertDocument，end...", zap.String("cost", time.Since(now).String()))
		if err == gorm.ErrRecordNotFound {
			time.Sleep(sleep)
		}
	}
}

func (m *DBModel) ReconvertDocoument(documentId int64, ext string) {
	ext = "." + strings.TrimLeft(ext, ".")
	os.RemoveAll(cacheReconvert)
	os.MkdirAll(cacheReconvert, os.ModePerm)
	if documentId <= 0 {
		m.reconvertAllDocument(ext)
		return
	}

	doc, err := m.GetDocument(documentId)
	if err != nil {
		m.logger.Error("ReconvertDocoument", zap.Error(err))
		return
	}
	if doc.Status != DocumentStatusConverted {
		m.logger.Error("ReconvertDocoument", zap.Error(errors.New("文档不是已转换的文档，不能重转")))
		return
	}
	m.reconvertDocument(&doc, ext)
}

func (m *DBModel) reconvertDocument(doc *Document, ext string) {
	m.logger.Debug("reconvertDocument", zap.Any("doc", doc), zap.String("ext", ext))
	if doc.PreviewExt == ext {
		m.logger.Info("reconvertDocument", zap.String("msg", "文档预览文件格式与指定格式一致，无需重转"), zap.String("document", doc.Title+doc.Ext))
		return
	}

	// 1. 下载文档预览文件
	attachment := m.GetAttachmentByTypeAndTypeId(AttachmentTypeDocument, doc.Id, "id", "hash")
	if attachment.Id == 0 {
		m.logger.Error("reconvertDocument", zap.String("msg", "文档预览文件不存在"), zap.String("document", doc.Title+doc.Ext))
		return
	}
	cacheDir := filepath.Join(cacheReconvert, strconv.FormatInt(doc.Id, 10))
	os.MkdirAll(cacheDir, os.ModePerm)
	defer os.RemoveAll(cacheDir)

	totalPreview := doc.Preview
	if totalPreview == 0 {
		totalPreview = doc.Pages
	}

	var (
		convertedTargets []string
		oldSrcFiles      []string
	)

	for i := 1; i <= totalPreview; i++ {
		// 已存在的预览文件
		isGZIP := false
		oldExt := doc.PreviewExt
		if doc.EnableGZIP && strings.HasSuffix(oldExt, ".svg") {
			oldExt = ".gzip.svg"
			isGZIP = true
		}

		// 目标文件
		dstFile := filepath.Join(cacheDir, fmt.Sprintf("%d%s", i, oldExt))
		// 源文件
		srcFile := fmt.Sprintf("documents/%s/%s/%d%s", strings.Join(strings.Split(attachment.Hash, "")[:5], "/"), attachment.Hash, i, oldExt)
		oldSrcFiles = append(oldSrcFiles, srcFile)
		err := util.CopyFile(srcFile, dstFile)
		if err != nil {
			m.logger.Error("reconvertDocument", zap.String("msg", "下载文档预览文件失败"), zap.String("document", doc.Title+doc.Ext), zap.Error(err))
			return
		}
		m.logger.Debug("reconvertDocument", zap.Bool("isGZIP", isGZIP), zap.String("msg", "下载文档预览文件成功"), zap.String("document", doc.Title+doc.Ext), zap.String("srcFile", srcFile), zap.String("dstFile", dstFile))
		if isGZIP { // 解压缩
			m.ungzipSVG(dstFile)
		}

		// 2. 转换文档预览文件
		convertedTargetFile := filepath.Join(cacheDir, fmt.Sprintf("%d%s", i, ext))
		if strings.HasSuffix(oldExt, ".svg") {
			// 如果是svg文件，则需要使用inkscape预先转为png
			tmpFile := filepath.Join(cacheDir, fmt.Sprintf("tmp-%d.png", i))
			err = converter.ConvertByInkscape(dstFile, tmpFile)
			if err == nil {
				if strings.HasSuffix(convertedTargetFile, ".png") {
					// 如果目标文件是png，则直接使用inkscape转换后的文件
					convertedTargetFile = tmpFile
				} else {
					// 如果目标文件不是png，则需要使用ImageMagick转换
					err = converter.ConvertByImageMagick(tmpFile, convertedTargetFile)
					os.RemoveAll(tmpFile)
				}
			}
		} else {
			err = converter.ConvertByImageMagick(dstFile, convertedTargetFile)
		}
		if err != nil {
			m.logger.Error("reconvertDocument", zap.String("msg", "转换文档预览文件失败"), zap.String("document", doc.Title+doc.Ext), zap.Error(err))
			return
		}
		convertedTargets = append(convertedTargets, convertedTargetFile)
	}

	// 3. 上传文档预览文件
	for i, srcFile := range convertedTargets {
		dstFile := fmt.Sprintf("documents/%s/%s/%d%s", strings.Join(strings.Split(attachment.Hash, "")[:5], "/"), attachment.Hash, i+1, ext)
		err := util.CopyFile(srcFile, dstFile)
		if err != nil {
			m.logger.Error("reconvertDocument", zap.String("msg", "上传文档预览文件失败"), zap.String("document", doc.Title+doc.Ext), zap.Error(err))
			return
		}
	}

	// 4. 更新数据库表的预览后缀
	// 查询同一hash的文档
	var (
		attachemnts []Attachment
		err         error
		data        = map[string]interface{}{
			"preview_ext": ext,
			"enable_gzip": false,
		}
	)

	m.db.Select("id", "type_id").Where("hash = ? and `type` = ?", attachment.Hash, AttachmentTypeDocument).Find(&attachemnts)
	if len(attachemnts) > 0 {
		var ids []int64
		for _, attachemnt := range attachemnts {
			ids = append(ids, attachemnt.TypeId)
		}
		err = m.db.Model(&Document{}).Where("id IN (?)", ids).Updates(data).Error
	} else {
		err = m.db.Model(doc).Updates(data).Error
	}
	if err != nil {
		m.logger.Error("reconvertDocument", zap.String("msg", "更新文档预览文件后缀失败"), zap.String("document", doc.Title+doc.Ext), zap.Error(err))
		return
	}

	// 5. 删除缓存文件，删除原预览文件
	for _, file := range oldSrcFiles {
		os.Remove(file)
	}
}

func (m *DBModel) reconvertAllDocument(ext string) {
	var cfg reconvertDocument
	bytes, _ := os.ReadFile("cache/reconvert.json")
	json.Unmarshal(bytes, &cfg)
	for {
		var doc Document
		m.db.Where("id > ?", cfg.Id).Where("status = ?", DocumentStatusConverted).Order("id asc").Find(&doc)
		if doc.Id == 0 {
			break
		}
		m.reconvertDocument(&doc, ext)
		cfg.Id = doc.Id
		bytes, _ = json.Marshal(cfg)
		os.WriteFile("cache/reconvert.json", bytes, os.ModePerm)
	}
}

func (m *DBModel) ungzipSVG(svg string) {
	m.logger.Info("ungzipSVG", zap.String("svg", svg))
	bs, err := os.ReadFile(svg)
	if err != nil {
		m.logger.Error("ungzipSVG", zap.Error(err))
		return
	}
	gz, err := gzip.NewReader(bytes.NewReader(bs))
	if err != nil {
		m.logger.Error("ungzipSVG", zap.Error(err))
		return
	}
	defer gz.Close()
	fp, err := os.Create(svg)
	if err != nil {
		m.logger.Error("ungzipSVG", zap.Error(err))
		return
	}
	defer fp.Close()
	io.Copy(fp, gz)
}

func (m *DBModel) cronCheckLatestVersion() {
	for {
		// 每小时检测一次
		m.RefreshLatestRelease()
		time.Sleep(1 * time.Hour)
	}
}

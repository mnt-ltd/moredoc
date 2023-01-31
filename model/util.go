package model

import (
	"fmt"
	"moredoc/util/sitemap"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
			beforeHour, _                   = strconv.Atoi(os.Getenv("MOREDOC_CLEAN_ATTACHMENT")) // 默认为每天凌晨0点更新站点地图
		)

		if beforeHour <= 0 {
			beforeHour = 24
		}

		// 1. 找出已被标记删除的附件
		m.db.Unscoped().Where("deleted_at IS NOT NULL").Where("deleted_at < ?", time.Now().Add(-time.Duration(beforeHour)*time.Hour)).Limit(100).Find(&deletedAttachemnts)
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
		// 1. 查找图片类配置
		var (
			configs []Config
			hashes  []string
		)
		m.db.Select("value").Where("input_type = ?", "image").Find(&configs)
		if len(configs) > 0 {
			for _, config := range configs {
				// 文件hash
				hash := strings.TrimSpace(strings.TrimSuffix(filepath.Base(config.Value), filepath.Ext(config.Value)))
				if hash != "" {
					hashes = append(hashes, hash)
				}
			}
			err := m.db.Where("`hash` NOT IN (?) and `type` = ?", hashes, AttachmentTypeConfig).Delete(&Attachment{}).Error
			if err != nil {
				m.logger.Error("cronMarkAttachmentDeleted", zap.Error(err))
			}
		}

		// 非配置类附件，如果type_id为0，则表示未被使用，超过24小时则标记删除
		m.logger.Info("cronMarkAttachmentDeleted start...")
		err := m.db.Where("type != ?  and type_id = ?", AttachmentTypeConfig, 0).Where("created_at < ?", time.Now().Add(-time.Duration(24)*time.Hour)).Delete(&Attachment{}).Error
		if err != nil {
			m.logger.Error("cronMarkAttachmentDeleted", zap.Error(err))
		}
		m.logger.Info("cronMarkAttachmentDeleted end...")
	}
}

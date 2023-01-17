package model

import (
	"fmt"
	"moredoc/util/sitemap"
	"os"
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
		hour, _ := strconv.Atoi(os.Getenv("CRON_UPDATE_SITEMAP_HOUR")) // 默认为每天凌晨0点更新站点地图
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

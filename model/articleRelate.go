package model

import (
	"moredoc/util"
	"moredoc/util/segword/jieba"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ArticleRelate struct {
	Id               int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	ArticleId        int64      `form:"article_id" json:"article_id,omitempty" gorm:"column:article_id;type:bigint(20);size:20;default:0;comment:;index:idx_article_id,unique"`
	RelatedArticleId string     `form:"related_article_id" json:"related_article_id,omitempty" gorm:"column:related_article_id;type:text;comment:;"`
	CreatedAt        *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt        *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (ArticleRelate) TableName() string {
	return tablePrefix + "article_relate"
}

// CreateArticleRelate 创建ArticleRelate
func (m *DBModel) CreateArticleRelate(articleRelate *ArticleRelate) (err error) {
	err = m.db.Create(articleRelate).Error
	if err != nil {
		m.logger.Error("CreateArticleRelate", zap.Error(err))
		return
	}
	return
}

// UpdateArticleRelate 更新ArticleRelate，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateArticleRelate(articleRelate *ArticleRelate, updateFields ...string) (err error) {
	db := m.db.Model(articleRelate)
	tableName := ArticleRelate{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", articleRelate.Id).Updates(articleRelate).Error
	if err != nil {
		m.logger.Error("UpdateArticleRelate", zap.Error(err))
	}
	return
}

// GetArticleRelate 根据id获取ArticleRelate
func (m *DBModel) GetArticleRelate(id int64, fields ...string) (articleRelate ArticleRelate, err error) {
	db := m.db

	fields = m.FilterValidFields(ArticleRelate{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&articleRelate).Error
	return
}

type OptionGetArticleRelateList struct {
	Page         int
	Size         int
	WithCount    bool                      // 是否返回总数
	Ids          []interface{}             // id列表
	SelectFields []string                  // 查询字段
	QueryRange   map[string][2]interface{} // map[field][]{min,max}
	QueryIn      map[string][]interface{}  // map[field][]{value1,value2,...}
	QueryLike    map[string][]interface{}  // map[field][]{value1,value2,...}
	Sort         []string
}

// GetArticleRelateList 获取ArticleRelate列表
func (m *DBModel) GetArticleRelateList(opt *OptionGetArticleRelateList) (articleRelateList []ArticleRelate, total int64, err error) {
	tableName := ArticleRelate{}.TableName()
	db := m.db.Model(&ArticleRelate{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetArticleRelateList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&articleRelateList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetArticleRelateList", zap.Error(err))
	}
	return
}

// DeleteArticleRelate 删除数据
func (m *DBModel) DeleteArticleRelate(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&ArticleRelate{}).Error
	if err != nil {
		m.logger.Error("DeleteArticleRelate", zap.Error(err))
	}
	return
}

func (m *DBModel) GetRelatedArticles(identifier string, fields ...string) (articles []Article, err error) {
	var (
		relate   ArticleRelate
		ids      []int64
		cfg      = m.GetConfigOfSecurity(ConfigSecurityDocumentRelatedDuration)
		keywords []interface{}
		opt      = &OptionGetArticleList{
			WithCount:    false,
			Page:         1,
			Size:         11,
			QueryIn:      make(map[string][]interface{}),
			QueryLike:    make(map[string][]interface{}),
			SelectFields: []string{"id", "title", "keywords", "identifier"},
		}
		isExpired bool
		article   Article
	)
	if len(fields) > 0 {
		opt.SelectFields = fields
	}

	if cfg.DocumentRelatedDuration <= 0 {
		return
	}

	// 文章不存在
	m.db.Select("id", "title", "keywords").Where("identifier = ?", identifier).First(&article)
	if article.Id == 0 {
		return
	}

	err = m.db.Where("article_id = ?", article.Id).First(&relate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetRelatedArticles", zap.Error(err))
		return
	}

	// 未过期
	if relate.Id > 0 && relate.UpdatedAt.Add(time.Duration(cfg.DocumentRelatedDuration)*time.Hour*24).After(time.Now()) {
		json.Unmarshal([]byte(relate.RelatedArticleId), &ids)
	} else {
		isExpired = true
	}

	if len(ids) == 0 {
		for _, kw := range strings.Split(article.Keywords, ",") {
			kw = strings.TrimSpace(kw)
			if kw == "" {
				continue
			}
			keywords = append(keywords, strings.TrimSpace(kw))
		}
		if len(keywords) == 0 {
			// 从标题中提取关键词
			for _, kv := range jieba.SegWords(article.Title) {
				keywords = append(keywords, kv)
			}
		}
		opt.QueryLike["title"] = keywords
		opt.QueryLike["keywords"] = keywords
		opt.QueryLike["description"] = keywords
	} else {
		opt.QueryIn["id"] = util.Slice2Interface(ids)
	}
	articles, _, _ = m.GetArticleList(opt)
	if isExpired && len(articles) > 0 {
		for _, art := range articles {
			if art.Id == article.Id {
				continue
			}
			ids = append(ids, art.Id)
			if len(ids) >= 10 {
				break
			}
		}
		bs, _ := json.Marshal(ids)
		relate.ArticleId = article.Id
		relate.RelatedArticleId = string(bs)
		if relate.Id > 0 {
			m.UpdateArticleRelate(&relate)
		} else {
			m.CreateArticleRelate(&relate)
		}
	}
	return
}

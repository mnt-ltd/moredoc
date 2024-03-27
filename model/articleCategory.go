package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ArticleCategory struct {
	Id         int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	ArticleId  int64      `form:"article_id" json:"article_id,omitempty" gorm:"column:article_id;type:int(11);size:11;default:0;index:idx_article_id;index:idx_article_category,unique;comment:文章ID;"`
	CategoryId int64      `form:"category_id" json:"category_id,omitempty" gorm:"column:category_id;type:int(11);size:11;default:0;index:idx_category_id;index:idx_article_category,unique;comment:分类ID;"`
	CreatedAt  *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt  *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (ArticleCategory) TableName() string {
	return tablePrefix + "article_category"
}

func (m *DBModel) GetArticleCategories(articleId ...int64) (categories []ArticleCategory, err error) {
	if len(articleId) == 0 {
		return
	}

	err = m.db.Where("article_id in ?", articleId).Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetArticleCategories", zap.Error(err))
	}
	return
}

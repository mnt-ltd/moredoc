package model

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DocumentScore struct {
	Id         int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	DocumentId int64      `form:"document_id" json:"document_id,omitempty" gorm:"column:document_id;type:bigint(20);size:20;default:0;comment:文档ID;index:idx_document_user,unique"`
	UserId     int64      `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;comment:用户ID;index:idx_document_user,unique"`
	Score      int        `form:"score" json:"score,omitempty" gorm:"column:score;type:int(11);size:11;default:0;comment:文档评分值，3位数，如500表示5分;"`
	CreatedAt  *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt  *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (DocumentScore) TableName() string {
	return tablePrefix + "document_score"
}

// CreateDocumentScore 创建DocumentScore
func (m *DBModel) CreateDocumentScore(documentScore *DocumentScore) (err error) {
	doc, _ := m.GetDocument(documentScore.DocumentId, "id", "score_count", "score")
	if doc.Id == 0 {
		err = fmt.Errorf("文档不存在")
		return
	}

	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 创建评分记录
	err = tx.Create(documentScore).Error
	if err != nil {
		m.logger.Error("CreateDocumentScore", zap.Error(err))
		return
	}

	score := (documentScore.Score + doc.Score*doc.ScoreCount) / (doc.ScoreCount + 1)
	// 对应文档，评分总数+1，评分总值+score
	err = tx.Model(&Document{}).Where("id = ?", documentScore.DocumentId).Updates(
		map[string]interface{}{
			"score_count": gorm.Expr("score_count + ?", 1),
			"score":       score,
		},
	).Error
	if err != nil {
		m.logger.Error("CreateDocumentScore", zap.Error(err))
		return
	}
	return
}

// GetDocumentScore 获取DocumentScore
func (m *DBModel) GetDocumentScore(userId, documentId int64) (documentScore DocumentScore, err error) {
	err = m.db.Where("user_id = ? and document_id = ?", userId, documentId).Find(&documentScore).Error
	return
}

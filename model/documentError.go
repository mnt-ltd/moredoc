package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DocumentError struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Message   string     `form:"message" json:"message,omitempty" gorm:"column:message;type:text;comment:文档转换失败原因;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (DocumentError) TableName() string {
	return tablePrefix + "document_error"
}

// SetDocumentConvertError 设置文档转换失败的错误信息
// 如果 err 为 nil，则删除转换失败的记录
// 如果 err 不为 nil，则创建或更新转换失败的记录
func (m *DBModel) SetDocumentConvertError(documentId int64, err error) error {
	de := &DocumentError{
		Id: documentId,
	}
	if err == nil {
		// 如果没有错误，则删除转换失败的记录
		err = m.db.Delete(de).Error
		if err != nil {
			m.logger.Error("SetConvertError", zap.Error(err))
			return err
		}
		return nil
	}
	de.Message = err.Error()
	var exist DocumentError
	if err = m.db.First(&exist, documentId).Error; err != nil {
		// 如果不存在，则创建
		err = m.db.Create(de).Error
		if err != nil {
			m.logger.Error("SetConvertError", zap.Error(err))
		}
		return err
	}
	// 如果存在，则更新
	return m.db.Model(&exist).Updates(de).Error
}

func (m *DBModel) GetConvertError(documentIds ...int64) (errors map[int64]string) {
	errors = make(map[int64]string)
	if len(documentIds) == 0 {
		return
	}

	var des []DocumentError
	err := m.db.Model(&DocumentError{}).Where("id IN (?)", documentIds).Find(&des).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConvertError", zap.Error(err))
		return
	}

	for _, de := range des {
		errors[de.Id] = de.Message
	}
	return
}

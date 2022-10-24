package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DocumentCategory struct {
	Id         int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	DocumentId int64      `form:"document_id" json:"document_id,omitempty" gorm:"column:document_id;type:bigint(20);size:20;default:0;comment:文档ID;"`
	CategoryId int64      `form:"category_id" json:"category_id,omitempty" gorm:"column:category_id;type:bigint(20);size:20;default:0;index:category_id;comment:分类ID;"`
	CreatedAt  *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt  *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

// 这里是proto文件中的结构体，可以根据需要删除或者调整
//message DocumentCategory {
// int64 id = 1;
// int64 document_id = 2;
// int64 category_id = 3;
// google.protobuf.Timestamp created_at = 4 [ (gogoproto.stdtime) = true ];
// google.protobuf.Timestamp updated_at = 5 [ (gogoproto.stdtime) = true ];
//}

func (DocumentCategory) TableName() string {
	return tablePrefix + "document_category"
}

// CreateDocumentCategory 创建DocumentCategory
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreateDocumentCategory(documentCategory *DocumentCategory) (err error) {
	err = m.db.Create(documentCategory).Error
	if err != nil {
		m.logger.Error("CreateDocumentCategory", zap.Error(err))
		return
	}
	return
}

// UpdateDocumentCategory 更新DocumentCategory，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateDocumentCategory(documentCategory *DocumentCategory, updateFields ...string) (err error) {
	db := m.db.Model(documentCategory)

	updateFields = m.FilterValidFields(DocumentCategory{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	}

	err = db.Where("id = ?", documentCategory.Id).Updates(documentCategory).Error
	if err != nil {
		m.logger.Error("UpdateDocumentCategory", zap.Error(err))
	}
	return
}

// GetDocumentCategory 根据id获取DocumentCategory
func (m *DBModel) GetDocumentCategory(id interface{}, fields ...string) (documentCategory DocumentCategory, err error) {
	db := m.db

	fields = m.FilterValidFields(DocumentCategory{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&documentCategory).Error
	return
}

type OptionGetDocumentCategoryList struct {
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

// GetDocumentCategoryList 获取DocumentCategory列表
func (m *DBModel) GetDocumentCategoryList(opt *OptionGetDocumentCategoryList) (documentCategoryList []DocumentCategory, total int64, err error) {
	model := &DocumentCategory{}
	db := m.db.Model(model)
	db = m.generateQueryIn(db, model.TableName(), opt.QueryIn)
	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetDocumentCategoryList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(model.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&documentCategoryList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetDocumentCategoryList", zap.Error(err))
	}
	return
}

// DeleteDocumentCategory 删除数据
// TODO: 删除数据之后，存在 document_category_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeleteDocumentCategory(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&DocumentCategory{}).Error
	if err != nil {
		m.logger.Error("DeleteDocumentCategory", zap.Error(err))
	}
	return
}

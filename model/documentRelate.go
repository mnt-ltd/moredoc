package model

import (
	"moredoc/util"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DocumentRelate struct {
	Id                int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	DocumentId        int64      `form:"document_id" json:"document_id,omitempty" gorm:"column:document_id;type:bigint(20);size:20;default:0;comment:;index:idx_document_id,unique"`
	RelatedDocumentId string     `form:"related_document_id" json:"related_document_id,omitempty" gorm:"column:related_document_id;type:text;comment:;"`
	CreatedAt         *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt         *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (DocumentRelate) TableName() string {
	return tablePrefix + "document_relate"
}

// CreateDocumentRelate 创建DocumentRelate
func (m *DBModel) CreateDocumentRelate(documentRelate *DocumentRelate) (err error) {
	err = m.db.Create(documentRelate).Error
	if err != nil {
		m.logger.Error("CreateDocumentRelate", zap.Error(err))
		return
	}
	return
}

// UpdateDocumentRelate 更新DocumentRelate，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateDocumentRelate(documentRelate *DocumentRelate, updateFields ...string) (err error) {
	db := m.db.Model(documentRelate)
	tableName := DocumentRelate{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", documentRelate.Id).Updates(documentRelate).Error
	if err != nil {
		m.logger.Error("UpdateDocumentRelate", zap.Error(err))
	}
	return
}

// GetDocumentRelate 根据id获取DocumentRelate
func (m *DBModel) GetDocumentRelate(id int64, fields ...string) (documentRelate DocumentRelate, err error) {
	db := m.db

	fields = m.FilterValidFields(DocumentRelate{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&documentRelate).Error
	return
}

type OptionGetDocumentRelateList struct {
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

// GetDocumentRelateList 获取DocumentRelate列表
func (m *DBModel) GetDocumentRelateList(opt *OptionGetDocumentRelateList) (documentRelateList []DocumentRelate, total int64, err error) {
	tableName := DocumentRelate{}.TableName()
	db := m.db.Model(&DocumentRelate{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetDocumentRelateList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&documentRelateList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetDocumentRelateList", zap.Error(err))
	}
	return
}

// DeleteDocumentRelate 删除数据
func (m *DBModel) DeleteDocumentRelate(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&DocumentRelate{}).Error
	if err != nil {
		m.logger.Error("DeleteDocumentRelate", zap.Error(err))
	}
	return
}

func (m *DBModel) GetRelatedDocuments(documentId int64) (docs []Document, err error) {
	var (
		docRelate DocumentRelate
		docIds    []int64
		cfg       = m.GetConfigOfSecurity(ConfigSecurityDocumentRelatedDuration)
		keywords  []interface{}
		opt       = &OptionGetDocumentList{
			WithCount: false,
			Page:      1,
			Size:      10,
			QueryIn:   make(map[string][]interface{}),
			QueryLike: make(map[string][]interface{}),
			SelectFields: []string{
				"id", "title", "ext",
			},
		}
		isExpired bool
	)

	if cfg.DocumentRelatedDuration <= 0 {
		return
	}

	err = m.db.Where("document_id = ?", documentId).First(&docRelate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetRelatedDocuments", zap.Error(err))
		return
	}

	// 未过期
	if docRelate.Id > 0 && docRelate.UpdatedAt.Add(time.Duration(cfg.DocumentRelatedDuration)*time.Hour*24).After(time.Now()) {
		json.Unmarshal([]byte(docRelate.RelatedDocumentId), &docIds)
	} else {
		isExpired = true
	}

	if len(docIds) == 0 {
		doc, _ := m.GetDocument(documentId, "id", "title", "keywords")
		if doc.Id > 0 {
			for _, kw := range strings.Split(doc.Keywords, ",") {
				keywords = append(keywords, strings.TrimSpace(kw))
			}
			opt.QueryLike["title"] = keywords
			opt.QueryLike["keywords"] = keywords
			opt.QueryLike["description"] = keywords
		}
	} else {
		opt.QueryIn["id"] = util.Slice2Interface(docIds)
	}
	docs, _, _ = m.GetDocumentList(opt)
	if isExpired && len(docs) > 0 {
		for _, doc := range docs {
			docIds = append(docIds, doc.Id)
		}
		bs, _ := json.Marshal(docIds)
		docRelate.DocumentId = documentId
		docRelate.RelatedDocumentId = string(bs)
		if docRelate.Id > 0 {
			m.UpdateDocumentRelate(&docRelate)
		} else {
			m.CreateDocumentRelate(&docRelate)
		}
	}
	return
}

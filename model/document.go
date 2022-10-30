package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	DocumentStatusPending    = iota // 待转换
	DocumentStatusConverting        // 转换中
	DocumentStatusConverted         // 已转换
	DocumentStatusFailed            // 转换失败
	DocumentStatusDisabled          // 已禁用
)

var DocumentStatusMap = map[int]struct{}{
	DocumentStatusPending:    {},
	DocumentStatusConverting: {},
	DocumentStatusConverted:  {},
	DocumentStatusFailed:     {},
	DocumentStatusDisabled:   {},
}

type Document struct {
	Id            int64           `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Title         string          `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(255);size:255;comment:文档名称;"`
	Keywords      string          `form:"keywords" json:"keywords,omitempty" gorm:"column:keywords;type:varchar(255);size:255;comment:文档关键字;"`
	Description   string          `form:"description" json:"description,omitempty" gorm:"column:description;type:varchar(512);size:512;comment:文档描述;"`
	UserId        int64           `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;index:user_id;comment:文档所属用户ID;"`
	Cover         string          `form:"cover" json:"cover,omitempty" gorm:"column:cover;type:varchar(255);size:255;comment:文档封面;"`
	Width         int             `form:"width" json:"width,omitempty" gorm:"column:width;type:int(11);size:11;default:0;comment:宽;"`
	Height        int             `form:"height" json:"height,omitempty" gorm:"column:height;type:int(11);size:11;default:0;comment:高;"`
	Preview       int             `form:"preview" json:"preview,omitempty" gorm:"column:preview;type:int(11);size:11;default:0;comment:允许预览页数;"`
	Pages         int             `form:"pages" json:"pages,omitempty" gorm:"column:pages;type:int(11);size:11;default:0;comment:文档页数;"`
	UUID          string          `form:"uuid" json:"uuid,omitempty" gorm:"column:uuid;type:varchar(36);size:36;index:idx_uuid,unique;comment:文档UUID，用于隐藏文档真实路径;"`
	DownloadCount int             `form:"download_count" json:"download_count,omitempty" gorm:"column:download_count;type:int(11);size:11;default:0;comment:下载人次;"`
	ViewCount     int             `form:"view_count" json:"view_count,omitempty" gorm:"column:view_count;type:int(11);size:11;default:0;comment:浏览人次;"`
	FavoriteCount int             `form:"favorite_count" json:"favorite_count,omitempty" gorm:"column:favorite_count;type:int(11);size:11;default:0;comment:收藏人次;"`
	CommentCount  int             `form:"comment_count" json:"comment_count,omitempty" gorm:"column:comment_count;type:int(11);size:11;default:0;comment:评论人次;"`
	Score         int             `form:"score" json:"score,omitempty" gorm:"column:score;type:int(11);size:11;default:0;comment:评分，3位整数表示，500表示5分;"`
	ScoreCount    int             `form:"score_count" json:"score_count,omitempty" gorm:"column:score_count;type:int(11);size:11;default:0;comment:评分数量;"`
	Price         int             `form:"price" json:"price,omitempty" gorm:"column:price;type:int(11);size:11;default:0;comment:价格，0表示免费;"`
	Size          int64           `form:"size" json:"size,omitempty" gorm:"column:size;type:bigint(20);size:20;default:0;comment:文件大小;"`
	Status        int             `form:"status" json:"status,omitempty" gorm:"column:status;type:smallint(6);size:6;default:0;index:status;comment:文档状态：0 待转换，1 转换中，2 转换完成，3 转换失败，4 禁用;"`
	CreatedAt     *time.Time      `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt     *time.Time      `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
	DeletedAt     *gorm.DeletedAt `form:"deleted_at" json:"deleted_at,omitempty" gorm:"column:deleted_at;type:datetime;index:idx_deleted_at;comment:删除时间;"`
	DeletedUserId int64           `form:"deleted_user_id" json:"deleted_user_id,omitempty" gorm:"column:deleted_user_id;type:bigint(20);size:20;default:0;comment:删除用户ID;"`
}

func (Document) TableName() string {
	return tablePrefix + "document"
}

// CreateDocument 创建Document
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreateDocument(document *Document) (err error) {
	err = m.db.Create(document).Error
	if err != nil {
		m.logger.Error("CreateDocument", zap.Error(err))
		return
	}
	return
}

// UpdateDocument 更新Document，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateDocument(document *Document, categoryId []int64, updateFields ...string) (err error) {
	sess := m.db.Begin()
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()

	var (
		oldDocCategories      []DocumentCategory
		oldDocCategoryIds     []int64
		newDocCategories      []DocumentCategory
		modelDocumentCategory = &DocumentCategory{}
		modelCategory         = &Category{}
	)

	sess.Table(modelDocumentCategory.TableName()).Select("category_id").Where("document_id = ?", document.Id).Find(&oldDocCategories)
	for _, cate := range oldDocCategories {
		oldDocCategoryIds = append(oldDocCategoryIds, cate.CategoryId)
	}

	if len(oldDocCategoryIds) > 0 {
		err = sess.Where("document_id = ?", document.Id).Delete(modelDocumentCategory).Error
		if err != nil {
			m.logger.Error("Delete DocumentCategory", zap.Error(err))
			return
		}

		// 更新分类统计
		err = sess.Model(modelCategory).Where("id in (?)", oldDocCategoryIds).Update("doc_count", gorm.Expr("doc_count - ?", 1)).Error
		if err != nil {
			m.logger.Error("Update doc_count--", zap.Error(err))
			return
		}
	}

	for _, cateId := range categoryId {
		newDocCategories = append(newDocCategories, DocumentCategory{
			DocumentId: document.Id,
			CategoryId: cateId,
		})
	}
	if len(newDocCategories) > 0 {
		m.logger.Debug("newDocCategories", zap.Any("newDocCategories", newDocCategories))
		err = sess.Create(&newDocCategories).Error
		if err != nil {
			m.logger.Error("Create New Category", zap.Error(err))
			return
		}

		err = sess.Model(modelCategory).Where("id in (?)", categoryId).Update("doc_count", gorm.Expr("doc_count + ?", 1)).Error
		if err != nil {
			m.logger.Error("Update doc_count++", zap.Error(err))
			return
		}
	}

	updateFields = m.FilterValidFields(Document{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		sess = sess.Select(updateFields)
	} else {
		sess = sess.Select(m.GetTableFields(document.TableName())).Omit("deleted_at", "deleted_user_id")
	}

	err = sess.Where("id = ?", document.Id).Updates(document).Error
	if err != nil {
		m.logger.Error("UpdateDocument", zap.Error(err))
		return
	}

	return
}

// UpdateDocumentStatus 更新文档状态。如文档转换失败，重新更改为待转换等
func (m *DBModel) UpdateDocumentStatus(status int, documentId []int64) (err error) {
	// 1. 查询文档所属用户与所属分类
	// 2. 如果是禁用文档或者由禁用状态变更为其他状态，用户文档数加减1，分类文档数加减1
	// 3. 更新文档状态
	return
}

// GetDocument 根据id获取Document
func (m *DBModel) GetDocument(id interface{}, fields ...string) (document Document, err error) {
	db := m.db

	fields = m.FilterValidFields(Document{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&document).Error
	return
}

type OptionGetDocumentList struct {
	Page         int
	Size         int
	WithCount    bool                      // 是否返回总数
	Ids          []interface{}             // id列表
	SelectFields []string                  // 查询字段
	QueryRange   map[string][2]interface{} // map[field][]{min,max}
	QueryIn      map[string][]interface{}  // map[field][]{value1,value2,...}
	QueryLike    map[string][]interface{}  // map[field][]{value1,value2,...}
	Sort         []string
	IsRecycle    bool // 是否是回收站模式查询
}

// GetDocumentList 获取Document列表
func (m *DBModel) GetDocumentList(opt *OptionGetDocumentList) (documentList []Document, total int64, err error) {
	tableName := Document{}.TableName()
	db := m.db.Model(&Document{})
	if opt.IsRecycle {
		// 回收站模式，只根据删除的倒序排序
		opt.Sort = []string{"deleted_at desc"}
		db = db.Unscoped().Where("deleted_at IS NOT NULL")
	}

	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)
	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if categoryIds, ok := opt.QueryIn["category_id"]; ok && len(categoryIds) > 0 {
		tableCategory := DocumentCategory{}.TableName()
		db = db.Joins("left join "+tableCategory+" dc on dc.document_id = "+tableName+".id").Where("dc.category_id in (?)", categoryIds)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetDocumentList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) > 0 {
		db = m.generateQuerySort(db, tableName, opt.Sort)
	} else {
		db = db.Order("id desc")
	}

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)
	err = db.Find(&documentList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetDocumentList", zap.Error(err))
	}
	return
}

// DeleteDocument 删除数据
// TODO: 删除数据之后，存在 document_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
// TODO: 如果是删除文档，则文档所属分类以及用户的文档数量都需要减1（需要判断文档是否是禁用再做处理）
func (m *DBModel) DeleteDocument(ids []int64, deletedUserId int64, deepDelete ...bool) (err error) {
	var (
		docs                  []Document
		docCates              []DocumentCategory
		docFields             = []string{"id", "status", "user_id", "deleted_at", "deleted_user_id"}
		docCateFields         = []string{"id", "document_id", "category_id"}
		modelUser             = &User{}
		modelDocument         = &Document{}
		modelDocumentCategory = &DocumentCategory{}
		modelCategory         = &Category{}
		docCateMap            = make(map[int64][]int64)
	)

	// 1. 查询文档信息
	m.db.Model(modelDocument).Select(docFields).Unscoped().Where("id in (?)", ids).Find(&docs)
	m.db.Model(modelDocumentCategory).Select(docCateFields).Where("document_id in (?)", ids).Find(&docCates)

	for _, docCate := range docCates {
		docCateMap[docCate.DocumentId] = append(docCateMap[docCate.DocumentId], docCate.CategoryId)
	}

	sess := m.db.Begin()
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()

	for _, doc := range docs {
		// 文档不是禁用状态，且未被逻辑删除，则需要减少用户、分类下的文档统计数量
		// if doc.Status != DocumentStatusDisabled && doc.DeletedAt == nil {
		if doc.DeletedAt == nil {
			err = sess.Model(modelUser).Where("id = ?", doc.UserId).Update("doc_count", gorm.Expr("doc_count - ?", 1)).Error
			if err != nil {
				m.logger.Error("DeleteDocument", zap.Error(err))
				return
			}

			if cateIds, ok := docCateMap[doc.Id]; ok && len(cateIds) > 0 {
				err = sess.Model(modelCategory).Where("id in (?)", cateIds).Update("doc_count", gorm.Expr("doc_count - ?", 1)).Error
				if err != nil {
					m.logger.Error("DeleteDocument", zap.Error(err))
					return
				}
			}
		}

		if len(deepDelete) > 0 && deepDelete[0] { // 彻底删除
			err = sess.Unscoped().Delete(&doc).Error
			if err != nil {
				m.logger.Error("DeleteDocument", zap.Error(err))
				return
			}

			// 关联的分类也需要删除
			err = sess.Unscoped().Where("document_id = ?", doc.Id).Delete(modelDocumentCategory).Error
			if err != nil {
				m.logger.Error("DeleteDocument", zap.Error(err))
				return
			}
		} else { // 逻辑删除
			err = sess.Model(&doc).Where("id = ?", doc.Id).Updates(map[string]interface{}{
				"deleted_at":      time.Now(),
				"deleted_user_id": deletedUserId,
			}).Error
		}

		if err != nil {
			m.logger.Error("DeleteDocument", zap.Error(err))
			return
		}
	}

	if len(deepDelete) > 0 && deepDelete[0] { // 彻底删除
		err = sess.Model(&Attachment{}).Where("type = ? and type_id in (?)", AttachmentTypeDocument, ids).Update("type_id", 0).Error
		if err != nil {
			m.logger.Error("DeleteDocument", zap.Error(err))
			return
		}
	}

	return
}

// RecoverRecycleDocument 恢复回收站中的文档
func (m *DBModel) RecoverRecycleDocument(documentId []int64) (err error) {
	var (
		modelDocument      = &Document{}
		modelCategory      = &Category{}
		modelUser          = &User{}
		documentCategories []DocumentCategory
		docs               []Document
	)

	m.db.Select([]string{"category_id"}).Where("document_id in (?)", documentId).Find(&documentCategories)
	m.db.Select("user_id").Unscoped().Where("id in (?)", documentId).Find(&docs)

	sess := m.db.Begin()
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()

	err = sess.Model(modelDocument).Unscoped().Where("id in ?", documentId).Updates(map[string]interface{}{
		"deleted_at":      nil,
		"deleted_user_id": 0,
	}).Error

	if err != nil {
		m.logger.Error("RecoverRecycleDocument", zap.Error(err))
		return
	}

	for _, docCate := range documentCategories {
		err = sess.Model(modelCategory).Where("id = ?", docCate.CategoryId).Update("doc_count", gorm.Expr("doc_count + ?", 1)).Error
		if err != nil {
			m.logger.Error("RecoverRecycleDocument", zap.Error(err))
			return
		}
	}

	for _, doc := range docs {
		err = sess.Model(modelUser).Where("id = ?", doc.UserId).Update("doc_count", gorm.Expr("doc_count + ?", 1)).Error
		if err != nil {
			m.logger.Error("RecoverRecycleDocument", zap.Error(err))
			return
		}
	}

	return
}

func (m *DBModel) ClearRecycleDocument() (err error) {
	var docs []Document
	err = m.db.Unscoped().Select("id").Where("deleted_at is not null").Find(&docs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("ClearRecycleDocument", zap.Error(err))
		return
	}

	if len(docs) == 0 {
		err = nil
		return
	}

	var ids []int64
	for _, doc := range docs {
		ids = append(ids, doc.Id)
	}

	err = m.DeleteDocument(ids, 0, true)
	if err != nil {
		m.logger.Error("DeleteDocument", zap.Error(err))
	}

	return
}

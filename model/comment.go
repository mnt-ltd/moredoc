package model

import (
	// "fmt"
	// "strings"
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	CommentStatusPending  = iota // 待审核
	CommentStatusApproved        // 已审核
	CommentStatusRejected        // 已拒绝
)

type Comment struct {
	Id           int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	UserId       int64      `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;index:idx_user_id;comment:发布评论的用户;"`
	ParentId     int64      `form:"parent_id" json:"parent_id,omitempty" gorm:"column:parent_id;type:bigint(20);size:20;default:0;comment:上级ID;index:idx_parent_id;"`
	Content      string     `form:"content" json:"content,omitempty" gorm:"column:content;type:text;comment:评论内容;"`
	DocumentId   int64      `form:"document_id" json:"document_id,omitempty" gorm:"column:document_id;type:bigint(20);size:20;default:0;comment:文档ID;index:idx_document_id;"`
	Status       int8       `form:"status" json:"status,omitempty" gorm:"column:status;type:tinyint(4);size:4;default:0;comment:0 待审，1过审，2拒绝;"`
	CommentCount int        `form:"comment_count" json:"comment_count,omitempty" gorm:"column:comment_count;type:int(11);size:11;default:0;comment:评论数量;"`
	IP           string     `form:"ip" json:"ip,omitempty" gorm:"column:ip;type:varchar(64);size:64;default:'';comment:IP地址;"`
	CreatedAt    *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:评论时间;"`
	UpdatedAt    *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:评论更新时间;"`
}

func (Comment) TableName() string {
	return tablePrefix + "comment"
}

// CreateComment 创建Comment
func (m *DBModel) CreateComment(comment *Comment) (err error) {
	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Create(comment).Error
	if err != nil {
		m.logger.Error("CreateComment", zap.Error(err))
		return
	}

	// 文档评论数+1
	err = tx.Model(&Document{}).Where("id = ?", comment.DocumentId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		m.logger.Error("CreateComment", zap.Error(err))
		return
	}

	// 用户评论数+1
	err = tx.Model(&User{}).Where("id = ?", comment.UserId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		m.logger.Error("CreateComment", zap.Error(err))
		return
	}

	// 更新上级评论的评论数
	if comment.ParentId > 0 {
		err = tx.Model(&Comment{}).Where("id = ?", comment.ParentId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			m.logger.Error("CreateComment", zap.Error(err))
			return
		}
	}

	return
}

// UpdateComment 更新Comment，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateComment(comment *Comment, updateFields ...string) (err error) {
	db := m.db.Model(comment)
	tableName := Comment{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", comment.Id).Updates(comment).Error
	if err != nil {
		m.logger.Error("UpdateComment", zap.Error(err))
	}
	return
}

// GetComment 根据id获取Comment
func (m *DBModel) GetComment(id interface{}, fields ...string) (comment Comment, err error) {
	db := m.db

	fields = m.FilterValidFields(Comment{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&comment).Error
	return
}

type OptionGetCommentList struct {
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

// GetCommentList 获取Comment列表
func (m *DBModel) GetCommentList(opt *OptionGetCommentList) (commentList []Comment, total int64, err error) {
	tableName := Comment{}.TableName()
	db := m.db.Model(&Comment{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetCommentList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&commentList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetCommentList", zap.Error(err))
	}
	return
}

// DeleteComment 删除数据
// 删除评论之后，对应文档的评论数量也要减少，对应的父级文档评论数量也要减少，用户评论数量也要减少
func (m *DBModel) DeleteComment(ids []int64, limitUserId ...int64) (err error) {
	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var (
		comments []Comment
		user     = &User{}
		document = &Document{}
	)
	cond := []string{"id in (?)"}
	args := []interface{}{ids}
	if len(limitUserId) > 0 {
		cond = append(cond, "user_id in (?)")
		args = append(args, limitUserId)
	}
	condStr := strings.Join(cond, " and ")
	tx.Where(condStr, args...).Select("id", "parent_id", "document_id", "user_id").Find(&comments)
	if len(comments) == 0 {
		err = errors.New("评论不存在或没有权限删除")
		return err
	}

	err = tx.Where(condStr, args...).Delete(&Comment{}).Error
	if err != nil {
		m.logger.Error("DeleteComment", zap.Error(err))
		return
	}

	for _, comment := range comments {
		// 更新文档评论数
		err = tx.Model(document).Where("id = ?", comment.DocumentId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			m.logger.Error("DeleteComment", zap.Error(err))
			return
		}

		// 更新父级评论数
		if comment.ParentId > 0 {
			err = tx.Model(&comment).Where("id = ?", comment.ParentId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
			if err != nil {
				m.logger.Error("DeleteComment", zap.Error(err))
				return
			}
		}

		// 更新用户评论数
		err = tx.Model(user).Where("id = ?", comment.UserId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			m.logger.Error("DeleteComment", zap.Error(err))
			return
		}
	}

	return
}

func (m *DBModel) UpdateCommentStatus(ids []int64, status int32) (err error) {
	err = m.db.Model(&Comment{}).Where("id in (?) and status != ?", ids, status).Update("status", status).Error
	if err != nil {
		m.logger.Error("UpdateCommentStatus", zap.Error(err))
	}
	return
}

func (m *DBModel) CountComment() (count int64, err error) {
	err = m.db.Model(&Comment{}).Count(&count).Error
	if err != nil {
		m.logger.Error("CountComment", zap.Error(err))
	}
	return
}

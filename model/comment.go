package model

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Comment struct {
	Id         int64     `form:"id" json:"id,omitempty" gorm:"column:id;type:bigint(20) unsigned;default:0;primarykey;autoIncrement;comment:回复 id;"`
	UserId     int64     `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20) unsigned;default:0;comment:发表用户 id;"`
	DocumentId int64     `form:"document_id" json:"document_id,omitempty" gorm:"column:document_id;type:bigint(20) unsigned;default:0;index:document_id;comment:文档ID;"`
	ParentId   int64     `form:"parent_id" json:"parent_id,omitempty" gorm:"column:parent_id;type:bigint(20) unsigned;default:0;comment:父级ID，上一个评论ID;"`
	Content    string    `form:"content" json:"content,omitempty" gorm:"column:content;type:text;default:;comment:内容;"`
	Ip         string    `form:"ip" json:"ip,omitempty" gorm:"column:ip;type:varchar(16);size:16;default:;comment:ip 地址;"`
	ReplyCount int       `form:"reply_count" json:"reply_count,omitempty" gorm:"column:reply_count;type:int(10) unsigned;default:0;comment:关联回复数;"`
	LikeCount  int       `form:"like_count" json:"like_count,omitempty" gorm:"column:like_count;type:int(10) unsigned;default:0;comment:喜欢数;"`
	CreatedAt  time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;default:;comment:创建时间;"`
	UpdatedAt  time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;default:;comment:更新时间;"`
	DeletedAt  gorm.DeletedAt
	IsFirst    int8 `form:"is_first" json:"is_first,omitempty" gorm:"column:is_first;type:tinyint(3) unsigned;default:0;comment:是否首个回复;"`
	IsComment  int8 `form:"is_comment" json:"is_comment,omitempty" gorm:"column:is_comment;type:tinyint(3) unsigned;default:0;comment:是否是回复回帖的内容;"`
	IsApproved int8 `form:"is_approved" json:"is_approved,omitempty" gorm:"column:is_approved;type:tinyint(3) unsigned;default:1;comment:是否合法;"`
}

func (Comment) TableName() string {
	return tablePrefix + "comment"
}

// CreateComment 创建Comment
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreateComment(comment *Comment) (err error) {
	err = m.db.Create(comment).Error
	if err != nil {
		m.logger.Error("CreateComment", zap.Error(err))
		return
	}
	return
}

// UpdateComment 更新Comment，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateComment(comment *Comment, updateFields ...string) (err error) {
	db := m.db.Model(comment)

	updateFields = m.FilterValidFields(Comment{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
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
func (m *DBModel) GetCommentList(opt OptionGetCommentList) (commentList []Comment, total int64, err error) {
	db := m.db.Model(&Comment{})

	for field, rangeValue := range opt.QueryRange {
		fields := m.FilterValidFields(Comment{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		if rangeValue[0] != nil {
			db = db.Where(fmt.Sprintf("%s >= ?", field), rangeValue[0])
		}
		if rangeValue[1] != nil {
			db = db.Where(fmt.Sprintf("%s <= ?", field), rangeValue[1])
		}
	}

	for field, values := range opt.QueryIn {
		fields := m.FilterValidFields(Comment{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(fmt.Sprintf("%s in (?)", field), values)
	}

	for field, values := range opt.QueryLike {
		fields := m.FilterValidFields(Comment{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(strings.TrimSuffix(fmt.Sprintf(strings.Join(make([]string, len(values)+1), "%s like ? or"), field), "or"), values...)
	}

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

	opt.SelectFields = m.FilterValidFields(Comment{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) > 0 {
		var sorts []string
		for _, sort := range opt.Sort {
			slice := strings.Split(sort, " ")
			if len(m.FilterValidFields(Comment{}.TableName(), slice[0])) == 0 {
				continue
			}

			if len(slice) == 2 {
				sorts = append(sorts, fmt.Sprintf("%s %s", slice[0], slice[1]))
			} else {
				sorts = append(sorts, fmt.Sprintf("%s desc", slice[0]))
			}
		}
		if len(sorts) > 0 {
			db = db.Order(strings.Join(sorts, ","))
		}
	}

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&commentList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetCommentList", zap.Error(err))
	}
	return
}

// DeleteComment 删除数据
// TODO: 删除数据之后，存在 comment_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeleteComment(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Comment{}).Error
	if err != nil {
		m.logger.Error("DeleteComment", zap.Error(err))
	}
	return
}

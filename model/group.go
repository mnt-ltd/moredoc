package model

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Group struct {
	Id                    int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:用户组 id;"`
	Title                 string     `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(64);size:64;index:title,unique;comment:用户组名称;"`
	Color                 string     `form:"color" json:"color,omitempty" gorm:"column:color;type:varchar(20);size:20;comment:颜色;"`
	IsDefault             bool       `form:"is_default" json:"is_default,omitempty" gorm:"column:is_default;type:tinyint(3);default:0;index:is_default;comment:是否默认;"`
	IsDisplay             bool       `form:"is_display" json:"is_display,omitempty" gorm:"column:is_display;type:tinyint(3);default:0;comment:是否显示在用户名后;"`
	Description           string     `form:"description" json:"description,omitempty" gorm:"column:description;type:varchar(255);size:255;comment:用户组描述;"`
	UserCount             int        `form:"user_count" json:"user_count,omitempty" gorm:"column:user_count;type:int(11);size:11;default:0;comment:用户数量;"`
	Sort                  int        `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;comment:排序，值越大越靠前;"`
	EnableUpload          bool       `form:"enable_upload" json:"enable_upload,omitempty" gorm:"column:enable_upload;type:tinyint(3);default:0;comment:是否允许上传文档;"`
	EnableDocumentReview  bool       `form:"enable_document_review" json:"enable_document_review,omitempty" gorm:"column:enable_document_review;type:tinyint(3);default:0;comment:文档是否需要审核;"`
	EnableComment         bool       `form:"enable_comment" json:"enable_comment,omitempty" gorm:"column:enable_comment;type:tinyint(3);default:1;comment:是否允许评论;"`
	EnableCommentApproval bool       `form:"enable_comment_approval" json:"enable_comment_approval,omitempty" gorm:"column:enable_comment_approval;type:tinyint(3);default:0;comment:评论是否需要审核;"`
	EnableArticle         bool       `form:"enable_article" json:"enable_article,omitempty" gorm:"column:enable_article;type:tinyint(3);default:1;comment:是否允许发布文章;"`
	EnableArticleApproval bool       `form:"enable_article_approval" json:"enable_article_approval,omitempty" gorm:"column:enable_article_approval;type:tinyint(3);default:0;comment:文章是否需要审核;"`
	CreatedAt             *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt             *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Group) TableName() string {
	return tablePrefix + "group"
}

// CreateGroup 创建Group
func (m *DBModel) CreateGroup(group *Group) (err error) {
	sess := m.db.Begin()
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()
	if group.IsDefault {
		err = sess.Model(&Group{}).Where("is_default > ?", 0).Updates(map[string]interface{}{"is_default": false}).Error
		if err != nil {
			m.logger.Error("CreateGroup", zap.Error(err))
			return
		}
	}

	err = sess.Create(group).Error
	if err != nil {
		m.logger.Error("CreateGroup", zap.Error(err))
		return
	}
	return
}

// UpdateGroup 更新Group，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateGroup(group *Group, updateFields ...string) (err error) {
	sess := m.db.Begin()
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()

	if group.IsDefault {
		err = sess.Model(&Group{}).Where("is_default > ? and id != ?", 0, group.Id).Updates(map[string]interface{}{"is_default": false}).Error
		if err != nil {
			m.logger.Error("UpdateGroup", zap.Error(err))
			return
		}
	} else {
		var count int64
		sess.Model(&Group{}).Where("is_default > ? and id != ?", 0, group.Id).Count(&count)
		if count == 0 {
			err = errors.New("至少要有一个默认用户组")
			return
		}
	}

	updateFields = m.FilterValidFields(Group{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		sess = sess.Select(updateFields)
	} else { // 不更新用户统计数据
		sess = sess.Select(m.GetTableFields(Group{}.TableName())).Omit("id", "user_count")
	}

	err = sess.Where("id = ?", group.Id).Updates(group).Error
	if err != nil {
		m.logger.Error("UpdateGroup", zap.Error(err))
	}
	return
}

// GetGroup 根据id获取Group
func (m *DBModel) GetGroup(id int64, fields ...string) (group Group, err error) {
	db := m.db

	fields = m.FilterValidFields(Group{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&group).Error
	return
}

func (m *DBModel) GetGroupByTitle(title string) (group Group, err error) {
	err = m.db.Where("title = ?", title).First(&group).Error
	return
}

type OptionGetGroupList struct {
	Page         int
	Size         int
	WithCount    bool                     // 是否返回总数
	Ids          []interface{}            // id列表
	SelectFields []string                 // 查询字段
	QueryIn      map[string][]interface{} // map[field][]{value1,value2,...}
	QueryLike    map[string][]interface{} // map[field][]{value1,value2,...}
}

// GetGroupList 获取Group列表
func (m *DBModel) GetGroupList(opt *OptionGetGroupList) (groupList []Group, total int64, err error) {
	db := m.db.Model(&Group{})
	tableName := Group{}.TableName()

	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetGroupList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(Group{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = db.Order("sort desc, id asc").Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&groupList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetGroupList", zap.Error(err))
	}
	return
}

// GetGroupList 获取Group列表
func (m *DBModel) GetUserGroups(userId int64) (groupList []Group, err error) {
	tableName := Group{}.TableName() + " g"
	tableUserGroup := UserGroup{}.TableName() + " ug"
	db := m.DB().Table(tableName).Joins(
		"left join "+tableUserGroup+" on g.id = ug.group_id and ug.user_id = ?", userId,
	)
	db = db.Select(m.GetTableFields(tableName))
	err = db.Find(&groupList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetGroupList", zap.Error(err))
	}
	return
}

// DeleteGroup 删除数据
// 组下存在用户的，不能删除
// 默认组不能删除
func (m *DBModel) DeleteGroup(ids []int64) (err error) {
	var total int64
	m.db.Model(&Group{}).Where("id in (?) and (user_count > ? or is_default = ?)", ids, 0, true).Count(&total)
	if total > 0 {
		err = errors.New("默认分组以及分组下存在用户的组不能删除")
		return
	}

	err = m.db.Where("id in (?) and user_count = ?", ids, 0).Delete(&Group{}).Error
	if err != nil {
		m.logger.Error("DeleteGroup", zap.Error(err))
	}
	return
}

// GetDefaultUserGroup 获取默认的用户组
func (m *DBModel) GetDefaultUserGroup() (group Group, err error) {
	err = m.db.Where("is_default = ?", true).First(&group).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetDefaultUserGroup", zap.Error(err))
	}
	return
}

package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserGroup struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	UserId    int64      `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;index:user_group,unique;index:user_id;comment:用户ID;"`
	GroupId   int64      `form:"group_id" json:"group_id,omitempty" gorm:"column:group_id;type:bigint(20);size:20;default:0;index:user_group,unique;comment:组ID;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

// 这里是proto文件中的结构体，可以根据需要删除或者调整
//message UserGroup {
// int64 id = 1;
// int64 user_id = 2;
// int64 group_id = 3;
//   = 0;
//   = 0;
//}

func (UserGroup) TableName() string {
	return tablePrefix + "user_group"
}

// CreateUserGroup 创建UserGroup
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreateUserGroup(userGroup *UserGroup) (err error) {
	err = m.db.Create(userGroup).Error
	if err != nil {
		m.logger.Error("CreateUserGroup", zap.Error(err))
		return
	}
	return
}

// UpdateUserGroup 更新UserGroup，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateUserGroup(userGroup *UserGroup, updateFields ...string) (err error) {
	db := m.db.Model(userGroup)

	updateFields = m.FilterValidFields(UserGroup{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	}

	err = db.Where("id = ?", userGroup.Id).Updates(userGroup).Error
	if err != nil {
		m.logger.Error("UpdateUserGroup", zap.Error(err))
	}
	return
}

// GetUserGroup 根据id获取UserGroup
func (m *DBModel) GetUserGroup(id interface{}, fields ...string) (userGroup UserGroup, err error) {
	db := m.db

	fields = m.FilterValidFields(UserGroup{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&userGroup).Error
	return
}

type OptionGetUserGroupList struct {
	Page         int
	Size         int
	WithCount    bool                     // 是否返回总数
	Ids          []interface{}            // id列表
	SelectFields []string                 // 查询字段
	QueryIn      map[string][]interface{} // map[field][]{value1,value2,...}
	Sort         []string
}

// GetUserGroupList 获取UserGroup列表
func (m *DBModel) GetUserGroupList(opt *OptionGetUserGroupList) (userGroupList []UserGroup, total int64, err error) {
	db := m.db.Model(&UserGroup{})
	tableName := UserGroup{}.TableName()
	db = m.generateQueryIn(db, tableName, opt.QueryIn)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetUserGroupList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(UserGroup{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	err = db.Find(&userGroupList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetUserGroupList", zap.Error(err))
	}
	return
}

// DeleteUserGroup 删除数据
// TODO: 删除数据之后，存在 user_group_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeleteUserGroup(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&UserGroup{}).Error
	if err != nil {
		m.logger.Error("DeleteUserGroup", zap.Error(err))
	}
	return
}

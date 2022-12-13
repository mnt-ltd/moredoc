package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Permission struct {
	Id          int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Method      string     `form:"method" json:"method,omitempty" gorm:"column:method;type:varchar(16);size:16;index:method_path,unique;comment:请求方法，grpc为空;"`
	Path        string     `form:"path" json:"path,omitempty" gorm:"column:path;type:varchar(128);size:128;index:method_path,unique;comment:API路径;"`
	Title       string     `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(255);size:255;comment:中文名称;"`
	Description string     `form:"description" json:"description,omitempty" gorm:"column:description;type:varchar(255);size:255;comment:权限描述;"`
	CreatedAt   *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt   *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Permission) TableName() string {
	return tablePrefix + "permission"
}

// CreatePermission 创建Permission
func (m *DBModel) CreatePermission(permission *Permission) (err error) {
	err = m.db.Create(permission).Error
	if err != nil {
		m.logger.Error("CreatePermission", zap.Error(err))
		return
	}
	return
}

// UpdatePermission 更新Permission，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdatePermission(permission *Permission, updateFields ...string) (err error) {
	db := m.db.Model(permission)

	updateFields = m.FilterValidFields(Permission{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	}

	err = db.Where("id = ?", permission.Id).Updates(permission).Error
	if err != nil {
		m.logger.Error("UpdatePermission", zap.Error(err))
	}
	return
}

// GetPermission 根据id获取Permission
func (m *DBModel) GetPermission(id interface{}, fields ...string) (permission Permission, err error) {
	db := m.db

	fields = m.FilterValidFields(Permission{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&permission).Error
	return
}

// GetPermissionByMethodPath
func (m *DBModel) GetPermissionByMethodPath(method, path string, createIfNotExist bool, fields ...string) (permission Permission, err error) {
	db := m.db

	fields = m.FilterValidFields(Permission{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	db = db.Where("path = ?", path)
	if method != "" {
		db = db.Where("method = ?", method)
	} else {
		db = db.Where("method IS NULL or method = '' or method = 'GRPC'")
	}

	err = db.First(&permission).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetPermissionByIdentifier", zap.Error(err))
		return
	}

	if permission.Id > 0 {
		return
	}

	if createIfNotExist {
		if method == "" {
			method = "GRPC"
		}
		permission.Method = method
		permission.Path = path
		err = m.CreatePermission(&permission)
		if err != nil {
			m.logger.Error("GetPermissionByIdentifier", zap.Error(err))
			return
		}
	}
	return
}

// DeletePermission 删除数据
func (m *DBModel) DeletePermission(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Permission{}).Error
	if err != nil {
		m.logger.Error("DeletePermission", zap.Error(err))
	}
	return
}

// CheckPermissionByUserId 根据用户ID，检查用户是否有权限
func (m *DBModel) CheckPermissionByUserId(userId int64, path string, httpMethod ...string) (permission Permission, yes bool) {
	var (
		userGroups []UserGroup
		groupId    []int64
		method     string
	)

	if len(httpMethod) > 0 {
		method = httpMethod[0]
	}

	if userId > 0 {
		m.db.Where("user_id = ?", userId).Find(&userGroups)
		for _, ug := range userGroups {
			groupId = append(groupId, ug.GroupId)
		}
	}

	permission, yes = m.CheckPermissionByGroupId(groupId, method, path)
	// NOTE: ID为1的用户，拥有所有权限，可以理解为类似linux的root用户
	if !yes && userId == 1 {
		yes = true
		return
	}

	return permission, yes
}

// CheckPermissionByGroupId 根据用户所属用户组ID，检查用户是否有权限
func (m *DBModel) CheckPermissionByGroupId(groupId []int64, method, path string) (permission Permission, yes bool) {
	var err error
	fields := []string{"id", "method", "path", "title"}
	permission, err = m.GetPermissionByMethodPath(method, path, true, fields...)
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("CheckPermissionByGroupId", zap.Error(err))
	}

	if permission.Id == 0 { // 权限控制表里面不存在的记录，默认允许访问
		yes = true
		return
	}

	// 校验当前登录了的用户所属用户组，是否有权限
	var groupPermission GroupPermission
	err = m.db.Where("group_id in (?) and permission_id = ?", groupId, permission.Id).First(&groupPermission).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("CheckPermissionByGroupId", zap.Error(err))
	}

	// 如果有权限，返回true
	return permission, groupPermission.Id > 0
}

type OptionGetPermissionList struct {
	Page         int
	Size         int
	WithCount    bool                     // 是否返回总数
	SelectFields []string                 // 查询字段
	QueryIn      map[string][]interface{} // map[field][]{value1,value2,...}
	QueryLike    map[string][]interface{} // map[field][]{value1,value2,...}
}

// GetPermissionList 获取Permission列表
func (m *DBModel) GetPermissionList(opt *OptionGetPermissionList) (permissionList []Permission, total int64, err error) {
	db := m.db.Model(&Permission{})
	tableName := Permission{}.TableName()

	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetPermissionList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = db.Order("path asc").Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&permissionList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetPermissionList", zap.Error(err))
	}
	return
}

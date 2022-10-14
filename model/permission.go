package model

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Permission struct {
	Id          int64     `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Method      string    `form:"method" json:"method,omitempty" gorm:"column:method;type:varchar(16);size:16;index:method_path,unique;comment:请求方法，grpc为空;"`
	Path        string    `form:"path" json:"path,omitempty" gorm:"column:path;type:varchar(128);size:128;index:method_path,unique;comment:API路径;"`
	Title       string    `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(255);size:255;comment:中文名称;"`
	Description string    `form:"description" json:"description,omitempty" gorm:"column:description;type:varchar(255);size:255;comment:权限描述;"`
	CreatedAt   time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Permission) TableName() string {
	return tablePrefix + "permission"
}

// 这里是proto文件中的结构体，可以根据需要删除或者调整
//message Permission {
// int64 id = 1;
// string category = 2;
// string title = 3;
// string description = 4;
// string identifier = 5;
//   = 0;
//   = 0;
//}

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
		db = db.Where("method IS NULL or method = ''")
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

//  CheckPermissionByUserId 根据用户ID，检查用户是否有权限
func (m *DBModel) CheckPermissionByUserId(userId int64, method, path string) (yes bool) {
	var (
		userGroups []UserGroup
		groupId    []int64
	)

	// NOTE: ID为1的用户，拥有所有权限，可以理解为类似linux的root用户
	// TODO: 后续应该直接打开这里的注释
	// if userId == 1 {
	// 	return true
	// }

	if userId > 0 {
		m.db.Where("user_id = ?", userId).Find(&userGroups)
		for _, ug := range userGroups {
			groupId = append(groupId, ug.GroupId)
		}
	}

	yes = m.CheckPermissionByGroupId(groupId, method, path)
	return yes || userId == 1
}

//  CheckPermissionByGroupId 根据用户所属用户组ID，检查用户是否有权限
func (m *DBModel) CheckPermissionByGroupId(groupId []int64, method, path string) (yes bool) {
	fields := []string{"id", "method", "path"}
	permission, err := m.GetPermissionByMethodPath(method, path, true, fields...)
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("CheckPermissionByGroupId", zap.Error(err))
	}

	if permission.Id == 0 { // 权限控制表里面不存在的记录，默认允许访问
		return true
	}

	// 校验当前登录了的用户所属用户组，是否有权限
	var groupPermission GroupPermission
	err = m.db.Where("group_id in (?) and permission_id = ?", groupId, permission.Id).First(&groupPermission).Error
	if err != nil {
		m.logger.Error("CheckPermissionByGroupId", zap.Error(err))
	}

	// 如果有权限，返回true
	return groupPermission.Id > 0
}

type OptionGetPermissionList struct {
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

// GetPermissionList 获取Permission列表
func (m *DBModel) GetPermissionList(opt OptionGetPermissionList) (permissionList []Permission, total int64, err error) {
	db := m.db.Model(&Permission{})

	for field, rangeValue := range opt.QueryRange {
		fields := m.FilterValidFields(Permission{}.TableName(), field)
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
		fields := m.FilterValidFields(Permission{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(fmt.Sprintf("%s in (?)", field), values)
	}

	for field, values := range opt.QueryLike {
		fields := m.FilterValidFields(Permission{}.TableName(), field)
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
			m.logger.Error("GetPermissionList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(Permission{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) > 0 {
		var sorts []string
		for _, sort := range opt.Sort {
			slice := strings.Split(sort, " ")
			if len(m.FilterValidFields(Permission{}.TableName(), slice[0])) == 0 {
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

	err = db.Find(&permissionList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetPermissionList", zap.Error(err))
	}
	return
}

package model

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GroupPermission struct {
	Id           int64     `form:"id" json:"id,omitempty" gorm:"column:id;type:bigint(20);size:20;default:0;comment:;"`
	GroupId      int64     `form:"group_id" json:"group_id,omitempty" gorm:"primaryKey;autoIncrement;index:group_permission,unique;index:group_id;column:group_id;comment:组ID;"`
	PermissionId int64     `form:"permission_id" json:"permission_id,omitempty" gorm:"primaryKey;autoIncrement;index:group_permission,unique;column:permission_id;comment:权限ID;"`
	CreatedAt    time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt    time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

// 这里是proto文件中的结构体，可以根据需要删除或者调整
//message GroupPermission {
// int64 id = 1;
// int64 group_id = 2;
// int64 permission_id = 3;
//   = 0;
//   = 0;
//}

func (GroupPermission) TableName() string {
	return tablePrefix + "group_permission"
}

// CreateGroupPermission 创建GroupPermission
func (m *DBModel) CreateGroupPermission(groupPermission *GroupPermission) (err error) {
	err = m.db.Create(groupPermission).Error
	if err != nil {
		m.logger.Error("CreateGroupPermission", zap.Error(err))
		return
	}
	return
}

// UpdateGroupPermission 更新GroupPermission，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateGroupPermission(groupPermission *GroupPermission, updateFields ...string) (err error) {
	db := m.db.Model(groupPermission)

	updateFields = m.FilterValidFields(GroupPermission{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	}

	err = db.Where("id = ?", groupPermission.Id).Updates(groupPermission).Error
	if err != nil {
		m.logger.Error("UpdateGroupPermission", zap.Error(err))
	}
	return
}

// GetGroupPermission 根据id获取GroupPermission
func (m *DBModel) GetGroupPermission(id interface{}, fields ...string) (groupPermission GroupPermission, err error) {
	db := m.db

	fields = m.FilterValidFields(GroupPermission{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&groupPermission).Error
	return
}

// GetGroupPermissionByGroupIdPermissionId(groupId int64, permissionId int64, fields ...string) 根据唯一索引获取GroupPermission
func (m *DBModel) GetGroupPermissionByGroupIdPermissionId(groupId int64, permissionId int64, fields ...string) (groupPermission GroupPermission, err error) {
	db := m.db

	fields = m.FilterValidFields(GroupPermission{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	db = db.Where("group_id = ?", groupId)

	db = db.Where("permission_id = ?", permissionId)

	err = db.First(&groupPermission).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetGroupPermissionByGroupIdPermissionId", zap.Error(err))
		return
	}
	return
}

type OptionGetGroupPermissionList struct {
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

// GetGroupPermissionList 获取GroupPermission列表
func (m *DBModel) GetGroupPermissionList(opt OptionGetGroupPermissionList) (groupPermissionList []GroupPermission, total int64, err error) {
	db := m.db.Model(&GroupPermission{})

	for field, rangeValue := range opt.QueryRange {
		fields := m.FilterValidFields(GroupPermission{}.TableName(), field)
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
		fields := m.FilterValidFields(GroupPermission{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(fmt.Sprintf("%s in (?)", field), values)
	}

	for field, values := range opt.QueryLike {
		fields := m.FilterValidFields(GroupPermission{}.TableName(), field)
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
			m.logger.Error("GetGroupPermissionList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(GroupPermission{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) > 0 {
		var sorts []string
		for _, sort := range opt.Sort {
			slice := strings.Split(sort, " ")
			if len(m.FilterValidFields(GroupPermission{}.TableName(), slice[0])) == 0 {
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

	err = db.Find(&groupPermissionList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetGroupPermissionList", zap.Error(err))
	}
	return
}

// DeleteGroupPermission 删除数据
func (m *DBModel) DeleteGroupPermission(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&GroupPermission{}).Error
	if err != nil {
		m.logger.Error("DeleteGroupPermission", zap.Error(err))
	}
	return
}

package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Navigation struct {
	Id          int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Title       string     `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(255);size:255;comment:链接名称;"`
	Href        string     `form:"href" json:"href,omitempty" gorm:"column:href;type:varchar(255);size:255;comment:跳转链接;"`
	Target      string     `form:"target" json:"target,omitempty" gorm:"column:target;type:varchar(16);size:16;comment:打开方式;"`
	Color       string     `form:"color" json:"color,omitempty" gorm:"column:color;type:varchar(16);size:16;comment:链接颜色;"`
	Sort        int        `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;comment:排序，值越大越靠前;"`
	Enable      *bool      `form:"enable" json:"enable,omitempty" gorm:"column:enable;type:int(11);size:11;default:0;comment:是否启用;"`
	ParentId    int64      `form:"parent_id" json:"parent_id,omitempty" gorm:"column:parent_id;type:int(11);size:11;default:0;comment:上级id;"`
	Description string     `form:"description" json:"description,omitempty" gorm:"column:description;type:varchar(1024);size:1024;comment:描述;"`
	CreatedAt   *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt   *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Navigation) TableName() string {
	return tablePrefix + "navigation"
}

// CreateNavigation 创建Navigation
func (m *DBModel) CreateNavigation(navigation *Navigation) (err error) {
	err = m.db.Create(navigation).Error
	if err != nil {
		m.logger.Error("CreateNavigation", zap.Error(err))
		return
	}
	return
}

// UpdateNavigation 更新Navigation，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateNavigation(navigation *Navigation, updateFields ...string) (err error) {
	db := m.db.Model(navigation)
	tableName := Navigation{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", navigation.Id).Updates(navigation).Error
	if err != nil {
		m.logger.Error("UpdateNavigation", zap.Error(err))
	}
	return
}

// GetNavigation 根据id获取Navigation
func (m *DBModel) GetNavigation(id interface{}, fields ...string) (navigation Navigation, err error) {
	db := m.db

	fields = m.FilterValidFields(Navigation{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&navigation).Error
	return
}

type OptionGetNavigationList struct {
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

// GetNavigationList 获取Navigation列表
func (m *DBModel) GetNavigationList(opt *OptionGetNavigationList) (navigationList []Navigation, total int64, err error) {
	tableName := Navigation{}.TableName()
	db := m.db.Model(&Navigation{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetNavigationList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&navigationList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetNavigationList", zap.Error(err))
	}
	return
}

// DeleteNavigation 删除数据
func (m *DBModel) DeleteNavigation(ids []int64) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Navigation{}).Error
	if err != nil {
		m.logger.Error("DeleteNavigation", zap.Error(err))
	}
	return
}

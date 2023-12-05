package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Advertisement struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	UserId    int64      `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;comment:用户ID;"`
	Position  string     `form:"position" json:"position,omitempty" gorm:"column:position;type:varchar(64);size:64;index:idx_position;comment:广告位;"`
	StartTime *time.Time `form:"start_time" json:"start_time,omitempty" gorm:"column:start_time;type:datetime;comment:开始时间;"`
	EndTime   *time.Time `form:"end_time" json:"end_time,omitempty" gorm:"column:end_time;type:datetime;comment:截止时间;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
	Content   string     `form:"content" json:"content,omitempty" gorm:"column:content;type:longtext;comment:广告内容;"`
	Enable    bool       `form:"enable" json:"enable,omitempty" gorm:"column:enable;type:tinyint(1);size:1;default:1;comment:是否启用;"`
	Remark    string     `form:"remark" json:"remark,omitempty" gorm:"column:remark;type:text;comment:备注;"`
}

func (Advertisement) TableName() string {
	return tablePrefix + "advertisement"
}

// CreateAdvertisement 创建Advertisement
func (m *DBModel) CreateAdvertisement(advertisement *Advertisement) (err error) {
	err = m.db.Create(advertisement).Error
	if err != nil {
		m.logger.Error("CreateAdvertisement", zap.Error(err))
		return
	}
	return
}

// UpdateAdvertisement 更新Advertisement，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateAdvertisement(advertisement *Advertisement, updateFields ...string) (err error) {
	db := m.db.Model(advertisement)
	tableName := Advertisement{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", advertisement.Id).Updates(advertisement).Error
	if err != nil {
		m.logger.Error("UpdateAdvertisement", zap.Error(err))
	}
	return
}

// GetAdvertisement 根据id获取Advertisement
func (m *DBModel) GetAdvertisement(id int64, fields ...string) (advertisement Advertisement, err error) {
	db := m.db

	fields = m.FilterValidFields(Advertisement{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&advertisement).Error
	return
}

type OptionGetAdvertisementList struct {
	Page         int
	Size         int
	WithCount    bool                      // 是否返回总数
	Ids          []int64                   // id列表
	SelectFields []string                  // 查询字段
	QueryRange   map[string][2]interface{} // map[field][]{min,max}
	QueryIn      map[string][]interface{}  // map[field][]{value1,value2,...}
	QueryLike    map[string][]interface{}  // map[field][]{value1,value2,...}
	Sort         []string
}

// GetAdvertisementList 获取Advertisement列表
func (m *DBModel) GetAdvertisementList(opt *OptionGetAdvertisementList) (advertisementList []Advertisement, total int64, err error) {
	tableName := Advertisement{}.TableName()
	db := m.db.Model(&Advertisement{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetAdvertisementList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	// TODO: 没有排序参数的话，可以自行指定排序字段
	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&advertisementList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetAdvertisementList", zap.Error(err))
	}
	return
}

// DeleteAdvertisement 删除数据
// TODO: 删除数据之后，存在 advertisement_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeleteAdvertisement(ids []int64) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Advertisement{}).Error
	if err != nil {
		m.logger.Error("DeleteAdvertisement", zap.Error(err))
	}
	return
}

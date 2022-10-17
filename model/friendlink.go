package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Friendlink struct {
	Id          int        `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Title       string     `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(64);size:64;comment:链接名称;"`
	Link        string     `form:"link" json:"link,omitempty" gorm:"column:link;type:varchar(255);size:255;comment:链接地址;"`
	Description string     `form:"description" json:"description,omitempty" gorm:"column:description;type:text;comment:描述，备注;"`
	Sort        int        `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;comment:排序，值越大越靠前;"`
	Enable      bool       `form:"enable" json:"enable,omitempty" gorm:"column:enable;type:tinyint(4);size:4;default:0;"`
	CreatedAt   *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt   *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Friendlink) TableName() string {
	return tablePrefix + "friendlink"
}

// GetFriendlinkPublicFields 获取Friendlink的公开字段
func (m *DBModel) GetFriendlinkPublicFields() []string {
	return []string{"id", "title", "link"}
}

// CreateFriendlink 创建Friendlink
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreateFriendlink(friendlink *Friendlink) (err error) {
	err = m.db.Create(friendlink).Error
	if err != nil {
		m.logger.Error("CreateFriendlink", zap.Error(err))
		return
	}
	return
}

// UpdateFriendlink 更新Friendlink，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateFriendlink(friendlink *Friendlink, updateFields ...string) (err error) {
	db := m.db.Model(friendlink)

	updateFields = m.FilterValidFields(Friendlink{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(friendlink.TableName()))
	}

	err = db.Where("id = ?", friendlink.Id).Updates(friendlink).Error
	if err != nil {
		m.logger.Error("UpdateFriendlink", zap.Error(err))
	}
	return
}

// GetFriendlink 根据id获取Friendlink
func (m *DBModel) GetFriendlink(id interface{}, fields ...string) (friendlink Friendlink, err error) {
	db := m.db

	fields = m.FilterValidFields(Friendlink{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&friendlink).Error
	return
}

type OptionGetFriendlinkList struct {
	Page         int
	Size         int
	WithCount    bool                     // 是否返回总数
	SelectFields []string                 // 查询字段
	QueryIn      map[string][]interface{} // map[field][]{value1,value2,...}
	QueryLike    map[string][]interface{} // map[field][]{value1,value2,...}
}

// GetFriendlinkList 获取Friendlink列表
func (m *DBModel) GetFriendlinkList(opt *OptionGetFriendlinkList) (friendlinkList []Friendlink, total int64, err error) {
	db := m.db.Model(&Friendlink{})
	tableName := Friendlink{}.TableName()

	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetFriendlinkList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(Friendlink{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Order("enable desc,sort desc").Find(&friendlinkList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetFriendlinkList", zap.Error(err))
	}
	return
}

// DeleteFriendlink 删除数据
func (m *DBModel) DeleteFriendlink(ids []int64) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Friendlink{}).Error
	if err != nil {
		m.logger.Error("DeleteFriendlink", zap.Error(err))
	}
	return
}

package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Dynamic struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	UserId    int64      `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;index:idx_user_id;comment:;"`
	Content   string     `form:"content" json:"content,omitempty" gorm:"column:content;type:text;comment:内容;"`
	Type      int        `form:"type" json:"type,omitempty" gorm:"column:type;type:smallint(6);size:6;default:0;comment:类型;index:idx_type;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

const (
	DynamicTypeComment        = 1  // 发表评论
	DynamicTypeFavorite       = 2  // 收藏文档
	DynamicTypeUpload         = 3  // 上传文档
	DynamicTypeDownload       = 4  // 下载文档
	DynamicTypeLogin          = 5  // 登录
	DynamicTypeRegister       = 6  // 注册
	DynamicTypeAvatar         = 7  // 更新了头像
	DynamicTypePassword       = 8  // 修改密码
	DynamicTypeInfo           = 9  // 修改个人信息
	DynamicTypeVerify         = 10 // 实名认证
	DynamicTypeSign           = 11 // 签到
	DynamicTypeShare          = 12 // 分享文档
	DynamicTypeFollow         = 13 // 关注用户
	DynamicTypeDeleteDocument = 14 // 删除文档
)

func (Dynamic) TableName() string {
	return tablePrefix + "dynamic"
}

// CreateDynamic 创建Dynamic
func (m *DBModel) CreateDynamic(dynamic *Dynamic) (err error) {
	err = m.db.Create(dynamic).Error
	if err != nil {
		m.logger.Error("CreateDynamic", zap.Error(err))
		return
	}
	return
}

// UpdateDynamic 更新Dynamic，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateDynamic(dynamic *Dynamic, updateFields ...string) (err error) {
	db := m.db.Model(dynamic)
	tableName := Dynamic{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", dynamic.Id).Updates(dynamic).Error
	if err != nil {
		m.logger.Error("UpdateDynamic", zap.Error(err))
	}
	return
}

// GetDynamic 根据id获取Dynamic
func (m *DBModel) GetDynamic(id interface{}, fields ...string) (dynamic Dynamic, err error) {
	db := m.db

	fields = m.FilterValidFields(Dynamic{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&dynamic).Error
	return
}

type OptionGetDynamicList struct {
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

// GetDynamicList 获取Dynamic列表
func (m *DBModel) GetDynamicList(opt *OptionGetDynamicList) (dynamicList []Dynamic, total int64, err error) {
	tableName := Dynamic{}.TableName()
	db := m.db.Model(&Dynamic{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetDynamicList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	} else {
		db = db.Select(m.GetTableFields(tableName))
	}

	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&dynamicList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetDynamicList", zap.Error(err))
	}
	return
}

// DeleteDynamic 删除数据
// TODO: 删除数据之后，存在 dynamic_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeleteDynamic(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Dynamic{}).Error
	if err != nil {
		m.logger.Error("DeleteDynamic", zap.Error(err))
	}
	return
}

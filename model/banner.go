package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Banner struct {
	Id          int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Title       string     `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(255);size:255;comment:横幅名称;"`
	Description string     `form:"description" json:"description,omitempty" gorm:"column:description;type:varchar(255);size:255;comment:横幅描述、备注;"`
	Path        string     `form:"path" json:"path,omitempty" gorm:"column:path;type:varchar(255);size:255;comment:横幅地址;"`
	Sort        int        `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;comment:排序，值越大越靠前;"`
	Enable      bool       `form:"enable" json:"enable,omitempty" gorm:"column:enable;type:tinyint(4);size:4;default:1;"`
	Type        int8       `form:"type" json:"type,omitempty" gorm:"column:type;type:tinyint(4);size:4;default:0;comment:0 网站横幅，1 小程序横幅，2 APP横幅;"`
	CreatedAt   *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt   *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
	Url         string     `form:"url" json:"url,omitempty" gorm:"column:url;type:varchar(255);size:255;comment:横幅跳转地址;"`
}

func (Banner) TableName() string {
	return tablePrefix + "banner"
}

// CreateBanner 创建Banner
func (m *DBModel) CreateBanner(banner *Banner) (err error) {
	err = m.db.Create(banner).Error
	if err != nil {
		m.logger.Error("CreateBanner", zap.Error(err))
		return
	}
	return
}

// UpdateBanner 更新Banner，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateBanner(banner *Banner, updateFields ...string) (err error) {
	db := m.db.Model(banner)
	tableName := Banner{}.TableName()
	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) == 0 {
		updateFields = m.GetTableFields(tableName)
	}
	db = db.Select(updateFields)

	err = db.Where("id = ?", banner.Id).Updates(banner).Error
	if err != nil {
		m.logger.Error("UpdateBanner", zap.Error(err))
	}
	return
}

// GetBanner 根据id获取Banner
func (m *DBModel) GetBanner(id interface{}, fields ...string) (banner Banner, err error) {
	db := m.db

	fields = m.FilterValidFields(Banner{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&banner).Error
	return
}

type OptionGetBannerList struct {
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

// GetBannerList 获取Banner列表
func (m *DBModel) GetBannerList(opt *OptionGetBannerList) (bannerList []Banner, total int64, err error) {
	db := m.db.Model(&Banner{})
	tableName := Banner{}.TableName()

	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetBannerList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Order("enable desc, sort desc").Find(&bannerList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetBannerList", zap.Error(err))
	}
	return
}

// DeleteBanner 删除数据
func (m *DBModel) DeleteBanner(ids []int64) (err error) {
	sess := m.db.Begin()
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()

	err = sess.Where("id in (?)", ids).Delete(&Banner{}).Error
	if err != nil {
		m.logger.Error("DeleteBanner", zap.Error(err))
		return
	}

	// 附件，标记删除
	err = sess.Model(&Attachment{}).Where("type = ? and type_id in (?)", AttachmentTypeBanner, ids).Update("type_id", 0).Error
	if err != nil {
		m.logger.Error("DeleteBanner", zap.Error(err))
		return
	}

	return
}

func (m *DBModel) CountBanner() (count int64, err error) {
	err = m.db.Model(&Banner{}).Count(&count).Error
	return
}

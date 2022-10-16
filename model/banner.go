package model

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Banner struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Title     string     `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(255);size:255;comment:横幅名称;"`
	Path      string     `form:"path" json:"path,omitempty" gorm:"column:path;type:varchar(255);size:255;comment:横幅地址;"`
	Sort      int        `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;comment:排序，值越大越靠前;"`
	Status    int8       `form:"status" json:"status,omitempty" gorm:"column:status;type:tinyint(4);size:4;default:0;comment:0 正常，1禁用;"`
	Category  int8       `form:"category" json:"category,omitempty" gorm:"column:category;type:tinyint(4);size:4;default:0;comment:0 PC横幅，1 H5横幅，2 小程序横幅，3 APP横幅;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
	Url       string     `form:"url" json:"url,omitempty" gorm:"column:url;type:varchar(255);size:255;comment:横幅跳转地址;"`
}

// 这里是proto文件中的结构体，可以根据需要删除或者调整
//message Banner {
// int64 id = 1;
// string title = 2;
// string path = 3;
// int32 sort = 4;
// int32 status = 5;
// int32 category = 6;
// google.protobuf.Timestamp created_at = 7 [ (gogoproto.stdtime) = true ];
// google.protobuf.Timestamp updated_at = 8 [ (gogoproto.stdtime) = true ];
// string url = 9;
//}

func (Banner) TableName() string {
	return tablePrefix + "banner"
}

// CreateBanner 创建Banner
// TODO: 创建成功之后，注意相关表统计字段数值的增减
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

	updateFields = m.FilterValidFields(Banner{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	}

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
func (m *DBModel) GetBannerList(opt OptionGetBannerList) (bannerList []Banner, total int64, err error) {
	db := m.db.Model(&Banner{})

	for field, rangeValue := range opt.QueryRange {
		fields := m.FilterValidFields(Banner{}.TableName(), field)
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
		fields := m.FilterValidFields(Banner{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(fmt.Sprintf("%s in (?)", field), values)
	}

	for field, values := range opt.QueryLike {
		fields := m.FilterValidFields(Banner{}.TableName(), field)
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
			m.logger.Error("GetBannerList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(Banner{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) > 0 {
		var sorts []string
		for _, sort := range opt.Sort {
			slice := strings.Split(sort, " ")
			if len(m.FilterValidFields(Banner{}.TableName(), slice[0])) == 0 {
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

	err = db.Find(&bannerList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetBannerList", zap.Error(err))
	}
	return
}

// DeleteBanner 删除数据
// TODO: 删除数据之后，存在 banner_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeleteBanner(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Banner{}).Error
	if err != nil {
		m.logger.Error("DeleteBanner", zap.Error(err))
	}
	return
}

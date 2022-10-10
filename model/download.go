package model

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Download struct {
	Id         int64     `form:"id" json:"id,omitempty" gorm:"column:id;type:bigint(20);size:20;default:0;primarykey;autoIncrement;comment:;"`
	UserId     int64     `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;index:user_id;comment:下载文档的用户ID;"`
	DocumentId int64     `form:"document_id" json:"document_id,omitempty" gorm:"column:document_id;type:bigint(20);size:20;default:0;comment:被下载的文档ID;"`
	Ip         string    `form:"ip" json:"ip,omitempty" gorm:"column:ip;type:varchar(16);size:16;default:;comment:下载文档的用户IP;"`
	CreatedAt  time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt  time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Download) TableName() string {
	return tablePrefix + "download"
}

// CreateDownload 创建Download
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreateDownload(download *Download) (err error) {
	err = m.db.Create(download).Error
	if err != nil {
		m.logger.Error("CreateDownload", zap.Error(err))
		return
	}
	return
}

// UpdateDownload 更新Download，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateDownload(download *Download, updateFields ...string) (err error) {
	db := m.db.Model(download)

	updateFields = m.FilterValidFields(Download{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	}

	err = db.Where("id = ?", download.Id).Updates(download).Error
	if err != nil {
		m.logger.Error("UpdateDownload", zap.Error(err))
	}
	return
}

// GetDownload 根据id获取Download
func (m *DBModel) GetDownload(id interface{}, fields ...string) (download Download, err error) {
	db := m.db

	fields = m.FilterValidFields(Download{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&download).Error
	return
}

type OptionGetDownloadList struct {
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

// GetDownloadList 获取Download列表
func (m *DBModel) GetDownloadList(opt OptionGetDownloadList) (downloadList []Download, total int64, err error) {
	db := m.db.Model(&Download{})

	for field, rangeValue := range opt.QueryRange {
		fields := m.FilterValidFields(Download{}.TableName(), field)
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
		fields := m.FilterValidFields(Download{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(fmt.Sprintf("%s in (?)", field), values)
	}

	for field, values := range opt.QueryLike {
		fields := m.FilterValidFields(Download{}.TableName(), field)
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
			m.logger.Error("GetDownloadList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(Download{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) > 0 {
		var sorts []string
		for _, sort := range opt.Sort {
			slice := strings.Split(sort, " ")
			if len(m.FilterValidFields(Download{}.TableName(), slice[0])) == 0 {
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

	err = db.Find(&downloadList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetDownloadList", zap.Error(err))
	}
	return
}

// DeleteDownload 删除数据
// TODO: 删除数据之后，存在 download_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeleteDownload(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Download{}).Error
	if err != nil {
		m.logger.Error("DeleteDownload", zap.Error(err))
	}
	return
}

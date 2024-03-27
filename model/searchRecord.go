package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SearchRecord struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	UserId    int64      `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;comment:搜索用户;"`
	Ip        string     `form:"ip" json:"ip,omitempty" gorm:"column:ip;type:varchar(64);size:64;comment:IP地址;"`
	Total     int        `form:"total" json:"total,omitempty" gorm:"column:total;type:int(11);size:11;default:0;comment:搜索结果;"`
	Page      int        `form:"page" json:"page,omitempty" gorm:"column:page;type:int(11);size:11;default:0;comment:搜索页码;"`
	UserAgent string     `form:"user_agent" json:"user_agent,omitempty" gorm:"column:user_agent;type:varchar(512);size:512;comment:请求客户端;"`
	Keywords  string     `form:"keywords" json:"keywords,omitempty" gorm:"column:keywords;type:varchar(64);size:64;comment:搜索关键字;"`
	SpendTime float64    `form:"spend_time" json:"spend_time,omitempty" gorm:"column:spend_time;type:decimal(10,2);size:10;default:0.00;comment:搜索耗时;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;index:idx_created_at;comment:创建时间，搜索时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
	Type      int        `form:"type" json:"type,omitempty" gorm:"column:type;type:int(11);size:11;default:0;comment:搜索类型,0文档，1文章;"`
}

var searchRecordQueue = make(chan *SearchRecord, 1024)

func (SearchRecord) TableName() string {
	return tablePrefix + "search_record"
}

// CreateSearchRecord 创建SearchRecord
func (m *DBModel) CreateSearchRecord(searchRecord *SearchRecord) (err error) {
	// 添加到队列
	now := time.Now()
	searchRecord.CreatedAt = &now
	searchRecord.UpdatedAt = &now
	searchRecordQueue <- searchRecord
	return
}

// UpdateSearchRecord 更新SearchRecord，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateSearchRecord(searchRecord *SearchRecord, updateFields ...string) (err error) {
	db := m.DB().Model(searchRecord)
	tableName := SearchRecord{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", searchRecord.Id).Updates(searchRecord).Error
	if err != nil {
		m.logger.Error("UpdateSearchRecord", zap.Error(err))
	}
	return
}

// GetSearchRecord 根据id获取SearchRecord
func (m *DBModel) GetSearchRecord(id int64, fields ...string) (searchRecord SearchRecord, err error) {
	db := m.db

	fields = m.FilterValidFields(SearchRecord{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&searchRecord).Error
	return
}

type OptionGetSearchRecordList struct {
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

// GetSearchRecordList 获取SearchRecord列表
func (m *DBModel) GetSearchRecordList(opt *OptionGetSearchRecordList) (searchRecordList []SearchRecord, total int64, err error) {
	tableName := SearchRecord{}.TableName()
	db := m.DB().Model(&SearchRecord{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetSearchRecordList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&searchRecordList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetSearchRecordList", zap.Error(err))
	}
	return
}

// DeleteSearchRecord 删除数据
func (m *DBModel) DeleteSearchRecord(ids []int64) (err error) {
	err = m.DB().Where("id in (?)", ids).Delete(&SearchRecord{}).Error
	if err != nil {
		m.logger.Error("DeleteSearchRecord", zap.Error(err))
	}
	return
}

// 通过队列创建搜索记录
func (m *DBModel) createSearchRecordFromQueue() {
	// 读取队列，批量添加到数据库中
	var (
		searchRecordList []*SearchRecord
		ticker           = time.NewTicker(time.Second * 10)
	)
	for {
		select {
		case searchRecord := <-searchRecordQueue:
			searchRecordList = append(searchRecordList, searchRecord)
		case <-ticker.C:
			if len(searchRecordList) > 0 {
				err := m.DB().CreateInBatches(&searchRecordList, 10).Error
				if err != nil {
					m.logger.Error("createSearchRecordByQueue", zap.Error(err))
				}
				searchRecordList = nil

				// 清理过期数据
				retentionDays := m.GetConfigOfSecurity(ConfigSecuritySearchRecordRetentionDays).SearchRecordRetentionDays
				err = m.DB().Where("created_at < ?", time.Now().AddDate(0, 0, -int(retentionDays)).Format("2006-01-02 00:00:00")).Delete(&SearchRecord{}).Error
				if err != nil {
					m.logger.Error("createSearchRecordByQueue", zap.Error(err))
				}
			}
		}
	}
}

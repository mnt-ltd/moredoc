package model

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Download struct {
	Id         int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	UserId     int64      `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;index:idx_user_id;comment:下载文档的用户ID;"`
	DocumentId int64      `form:"document_id" json:"document_id,omitempty" gorm:"column:document_id;type:bigint(20);size:20;index:idx_document_id;default:0;comment:被下载的文档ID;"`
	Ip         string     `form:"ip" json:"ip,omitempty" gorm:"column:ip;type:varchar(64);size:64;index:idx_ip;comment:下载文档的用户IP;"`
	IsPay      bool       `form:"is_pay" json:"is_pay,omitempty" gorm:"column:is_pay;type:tinyint(1);size:1;default:0;comment:是否付费下载;"`
	CreatedAt  *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;index:idx_created_at"`
	UpdatedAt  *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Download) TableName() string {
	return tablePrefix + "download"
}

// CreateDownload 创建Download
func (m *DBModel) CreateDownload(download *Download) (err error) {
	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Create(download).Error
	if err != nil {
		m.logger.Error("CreateDownload", zap.Error(err))
		return
	}

	err = tx.Model(&Document{}).Where("id = ?", download.DocumentId).Update("download_count", gorm.Expr("download_count + ?", 1)).Error
	if err != nil {
		m.logger.Error("CreateDownload", zap.Error(err))
		return
	}

	doc, _ := m.GetDocument(download.DocumentId, "id", "user_id", "price")
	if download.IsPay && doc.Price > 0 && download.UserId > 0 {
		// 下载该文档的用户扣除积分
		err = tx.Model(&User{}).Where("id = ?", download.UserId).Update("credit_count", gorm.Expr("credit_count - ?", doc.Price)).Error
		if err != nil {
			m.logger.Error("CreateDownload", zap.Error(err))
			return
		}

		// 文档的作者增加积分
		err = tx.Model(&User{}).Where("id = ?", doc.UserId).Update("credit_count", gorm.Expr("credit_count + ?", doc.Price)).Error
		if err != nil {
			m.logger.Error("CreateDownload", zap.Error(err))
			return
		}
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

// CanIFreeDownload 判断用户是否可以免费下载
// 最后一次付费下载时间 + 免费下载时长 > 当前时间
func (m *DBModel) CanIFreeDownload(userId, documentId int64) bool {
	var download Download
	m.db.Where("user_id = ? and document_id = ? and is_pay = ?", userId, documentId, 1).Last(&download)
	m.logger.Debug("CanIFreeDownload", zap.Any("Last Download", download))
	if download.Id == 0 {
		return false
	}

	cfg := m.GetConfigOfDownload(ConfigDownloadFreeDownloadDuration)
	m.logger.Debug("CanIFreeDownload", zap.Int32("FreeDownloadDuration", cfg.FreeDownloadDuration))
	if cfg.FreeDownloadDuration <= 0 {
		return true
	}
	m.logger.Debug("CanIFreeDownload", zap.Any("CreatedAt", download.CreatedAt), zap.Time("Now", time.Now()), zap.Time("After", download.CreatedAt.Add(time.Duration(cfg.FreeDownloadDuration)*time.Hour*24)))
	return download.CreatedAt.Add(time.Duration(cfg.FreeDownloadDuration) * time.Hour * 24).After(time.Now())
}

// CountDownloadToday 统计用户今日下载次数
func (m *DBModel) CountDownloadTodayForUser(userId int64) (total int64) {
	if userId == 0 {
		return
	}
	err := m.db.Model(&Download{}).Where("user_id = ?", userId).Where("created_at >= ?", time.Now().Format("2006-01-02")).Count(&total).Error
	if err != nil {
		m.logger.Error("CountDownloadToday", zap.Error(err))
	}
	return
}

// CountDownloadToday 统计IP今日下载次数
func (m *DBModel) CountDownloadTodayForIP(ip string) (total int64) {
	if ip == "" {
		return
	}
	err := m.db.Model(&Download{}).Where("ip = ?", ip).Where("created_at >= ?", time.Now().Format("2006-01-02")).Count(&total).Error
	if err != nil {
		m.logger.Error("CountDownloadTodayForIP", zap.Error(err))
	}
	return
}

package model

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Language struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Language  string     `form:"language" json:"language,omitempty" gorm:"column:language;type:varchar(64);size:64;comment:语言;"`
	Enable    bool       `form:"enable" json:"enable,omitempty" gorm:"column:enable;type:tinyint(1);size:1;default:0;index:idx_enable;comment:是否启用;"`
	Code      string     `form:"code" json:"code,omitempty" gorm:"column:code;type:varchar(16);size:16;comment:语言代码;index:idx_code,unique;"`
	Total     int        `form:"total" json:"total,omitempty" gorm:"column:total;type:int(11);size:11;default:0;comment:文档数;"`
	Sort      int        `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;comment:排序，值越大越靠前;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Language) TableName() string {
	return tablePrefix + "language"
}

func (m *DBModel) UpdateLanguageStatus(languageId []int64, enable bool) (err error) {
	err = m.DB().Model(&Language{}).Where("id in (?)", languageId).Update("enable", enable).Error
	if err != nil {
		m.logger.Error("UpdateLanguage", zap.Error(err))
	}
	return
}

// GetLanguage 根据id获取Language
func (m *DBModel) GetLanguage(id int64, fields ...string) (language Language, err error) {
	db := m.db

	fields = m.FilterValidFields(Language{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&language).Error
	return
}

// UpdateLanguage 更新Language
func (m *DBModel) UpdateLanguage(language *Language, field ...string) (err error) {
	db := m.DB().Model(language)
	if len(field) > 0 {
		db = db.Select(field)
	}
	err = db.Updates(language).Error
	if err != nil {
		m.logger.Error("UpdateLanguage", zap.Error(err))
	}
	return
}

type OptionGetLanguageList struct {
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

// GetLanguageList 获取Language列表
func (m *DBModel) GetLanguageList(opt *OptionGetLanguageList) (languageList []Language, total int64, err error) {
	tableName := Language{}.TableName()
	db := m.DB().Model(&Language{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetLanguageList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) == 0 {
		opt.Sort = []string{"enable desc", "sort desc", "id asc"}
	}
	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&languageList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetLanguageList", zap.Error(err))
	}
	return
}

func (m *DBModel) CreateLanguage(language *Language) (err error) {
	exist, _ := m.GetLanguageByCode(language.Code)
	if exist.Id > 0 {
		err = errors.New("语言代码已存在")
		return
	}

	err = m.DB().Create(language).Error
	if err != nil {
		m.logger.Error("CreateLanguage", zap.Error(err))
	}
	return
}

func (m *DBModel) DeleteLanguage(id []int64) (err error) {
	err = m.DB().Where("id in (?)", id).Delete(&Language{}).Error
	if err != nil {
		m.logger.Error("DeleteLanguage", zap.Error(err))
	}
	return
}

func (m *DBModel) GetLanguageByCode(code string, field ...string) (language Language, err error) {
	db := m.DB()
	if len(field) > 0 {
		db = db.Select(field)
	}

	err = db.Where("code = ?", code).Find(&language).Error
	if err != nil {
		m.logger.Error("GetLanguageByCode", zap.Error(err))
	}
	return
}

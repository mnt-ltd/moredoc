package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Language struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Language  string     `form:"language" json:"language,omitempty" gorm:"column:language;type:varchar(64);size:64;comment:语言;"`
	Enable    bool       `form:"enable" json:"enable,omitempty" gorm:"column:enable;type:tinyint(1);size:1;default:0;index:idx_enable;comment:是否启用;"`
	Code      string     `form:"code" json:"code,omitempty" gorm:"column:code;type:varchar(16);size:16;comment:语言代码;index:idx_code;"`
	Total     int        `form:"total" json:"total,omitempty" gorm:"column:total;type:int(11);size:11;default:0;comment:文档数;"`
	Sort      int        `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;comment:排序，值越大越靠前;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

// 这里是proto文件中的结构体，可以根据需要删除或者调整
//message Language {
// int64 id = 1;
// string language = 2;
// bool enable = 3;
// string code = 4;
// int32 total = 5;
// int32 sort = 6;
//}

func (Language) TableName() string {
	return tablePrefix + "language"
}

// CreateLanguage 创建Language
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreateLanguage(language *Language) (err error) {
	err = m.DB().Create(language).Error
	if err != nil {
		m.logger.Error("CreateLanguage", zap.Error(err))
		return
	}
	return
}

// UpdateLanguage 更新Language，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateLanguage(language *Language, updateFields ...string) (err error) {
	db := m.DB().Model(language)
	tableName := Language{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", language.Id).Updates(language).Error
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

	// TODO: 没有排序参数的话，可以自行指定排序字段
	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&languageList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetLanguageList", zap.Error(err))
	}
	return
}

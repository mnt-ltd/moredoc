package model

import (
	"fmt"
	"moredoc/util/captcha"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	// ConfigCategorySystem 系统配置：系统名称、logo、版权信息、是否闭站等
	ConfigCategorySystem = "system"
	// ConfigCategoryUser 用户配置：是否启用注册、是否需要审核等
	ConfigCategoryUser = "user"
	// ConfigCategoryEmail 邮箱配置：smtp服务器、端口、用户名、密码、发件人
	ConfigCategoryEmail = "email"
	// ConfigCategoryCaptcha 验证码配置：是否启用验证码、验证码有效期、验证码长度、验证码类型等
	ConfigCategoryCaptcha = "captcha"
	// ConfigCategoryJWT JWT配置：JWT有效期、JWT加密密钥等
	ConfigCategoryJWT = "jwt"
	// ConfigCategorySecurity 安全配置项
	ConfigCategorySecurity = "security"
)

type Config struct {
	Id          int64     `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Label       string    `form:"label" json:"label,omitempty" gorm:"column:label;type:varchar(64);size:64;comment:标签名称;"`
	Name        string    `form:"name" json:"name,omitempty" gorm:"column:name;type:varchar(64);size:64;index:name_category,unique;comment:表单字段名称;"`
	Value       string    `form:"value" json:"value,omitempty" gorm:"column:value;type:text;comment:值;"`
	Placeholder string    `form:"placeholder" json:"placeholder,omitempty" gorm:"column:placeholder;type:varchar(255);size:255;comment:提示信息;"`
	InputType   string    `form:"input_type" json:"input_type,omitempty" gorm:"column:input_type;type:varchar(32);size:32;default:text;comment:输入类型;"`
	Category    string    `form:"category" json:"category,omitempty" gorm:"column:category;type:varchar(32);size:32;index:name_category,unique;index:category;comment:所属类别;"`
	Sort        int       `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;comment:同一category下的排序;"`
	Options     string    `form:"options" json:"options,omitempty" gorm:"column:options;type:text;comment:针对checkbox等的枚举值;"`
	CreatedAt   time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

// 这里是proto文件中的结构体，可以根据需要删除或者调整
//message Config {
// int64 id = 1;
// string label = 2;
// string name = 3;
// string value = 4;
// int32 placeholder = 5;
// int32 input_type = 6;
// string category = 7;
// int32 sort = 8;
// string options = 9;
// google.protobuf.Timestamp created_at = 10 [ (gogoproto.stdtime) = true ];
// google.protobuf.Timestamp updated_at = 11 [ (gogoproto.stdtime) = true ];
//}

func (Config) TableName() string {
	return tablePrefix + "config"
}

// CreateConfig 创建Config
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreateConfig(config *Config) (err error) {
	err = m.db.Create(config).Error
	if err != nil {
		m.logger.Error("CreateConfig", zap.Error(err))
		return
	}
	return
}

// UpdateConfig 更新Config，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateConfig(config *Config, updateFields ...string) (err error) {
	db := m.db.Model(config)

	updateFields = m.FilterValidFields(Config{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	}

	err = db.Where("id = ?", config.Id).Updates(config).Error
	if err != nil {
		m.logger.Error("UpdateConfig", zap.Error(err))
	}
	return
}

// GetConfig 根据id获取Config
func (m *DBModel) GetConfig(id interface{}, fields ...string) (config Config, err error) {
	db := m.db

	fields = m.FilterValidFields(Config{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&config).Error
	return
}

// GetConfigByNameCategory(name string, category string, fields ...string) 根据唯一索引获取Config
func (m *DBModel) GetConfigByNameCategory(name string, category string, fields ...string) (config Config, err error) {
	db := m.db

	fields = m.FilterValidFields(Config{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	db = db.Where("name = ?", name)

	db = db.Where("category = ?", category)

	err = db.First(&config).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigByNameCategory", zap.Error(err))
		return
	}
	return
}

type OptionGetConfigList struct {
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

// GetConfigList 获取Config列表
func (m *DBModel) GetConfigList(opt OptionGetConfigList) (configList []Config, total int64, err error) {
	db := m.db.Model(&Config{})

	for field, rangeValue := range opt.QueryRange {
		fields := m.FilterValidFields(Config{}.TableName(), field)
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
		fields := m.FilterValidFields(Config{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(fmt.Sprintf("%s in (?)", field), values)
	}

	for field, values := range opt.QueryLike {
		fields := m.FilterValidFields(Config{}.TableName(), field)
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
			m.logger.Error("GetConfigList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(Config{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) > 0 {
		var sorts []string
		for _, sort := range opt.Sort {
			slice := strings.Split(sort, " ")
			if len(m.FilterValidFields(Config{}.TableName(), slice[0])) == 0 {
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

	err = db.Find(&configList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigList", zap.Error(err))
	}
	return
}

// DeleteConfig 删除数据
// TODO: 删除数据之后，存在 config_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeleteConfig(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Config{}).Error
	if err != nil {
		m.logger.Error("DeleteConfig", zap.Error(err))
	}
	return
}

const (
	ConfigCaptchaLength = "length"
	ConfigCaptchaWidth  = "width"
	ConfigCaptchaHeight = "height"
	ConfigCaptchaType   = "type"
)

type ConfigCaptcha struct {
	Length int    `json:"length"` // 验证码长度
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Type   string `json:"type"` // 验证码类型
}

// GetConfigOfCaptcha 获取验证码配置
func (m *DBModel) GetConfigOfCaptcha() (config ConfigCaptcha) {
	var configs []Config
	err := m.db.Where("category = ?", ConfigCategoryCaptcha).Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfCaptcha", zap.Error(err))
	}

	for _, cfg := range configs {
		switch cfg.Name {
		case "length":
			config.Length, _ = strconv.Atoi(cfg.Value)
			if config.Length <= 0 {
				config.Length = 6
			}
		case "width":
			config.Width, _ = strconv.Atoi(cfg.Value)
			if config.Width <= 0 {
				config.Width = 240
			}
		case "height":
			config.Height, _ = strconv.Atoi(cfg.Value)
			if config.Height <= 0 {
				config.Height = 60
			}
		case "type":
			// 验证码类型
			config.Type = cfg.Value
		}
	}
	return
}

const (
	ConfigSystemDomain      = "domain"
	ConfigSystemTitle       = "title"
	ConfigSystemDescription = "description"
	ConfigSystemKeywords    = "keywords"
	ConfigSystemLogo        = "logo"
	ConfigSystemFavicon     = "favicon"
	ConfigSystemIcp         = "icp"
	ConfigSystemAnalytics   = "analytics"
	ConfigSystemTheme       = "theme"
	ConfigSystemCopyright   = "copyright"
)

type ConfigSystem struct {
	Domain      string `json:"domain"`      // 站点域名，不带 HTTPS:// 和 HTTP://
	Title       string `json:"title"`       // 系统名称
	Keywords    string `json:"keywords"`    // 系统关键字
	Description string `json:"description"` // 系统描述
	Logo        string `json:"logo"`        // logo
	Favicon     string `json:"favicon"`     // logo
	Theme       string `json:"theme"`       // 网站主题
	Copyright   string `json:"copyright"`   // 版权信息
	ICP         string `json:"icp"`         // 网站备案
	Analytics   string `json:"analytics"`   // 统计代码
}

// GetConfigOfSystem 获取系统配置
func (m *DBModel) GetConfigOfSystem() (config ConfigSystem) {
	var confgis []Config
	err := m.db.Where("category = ?", ConfigCategorySystem).Find(&confgis).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfSystem", zap.Error(err))
	}

	var data = make(map[string]interface{})

	for _, cfg := range confgis {
		switch cfg.Name {
		// 字符串类型的配置项
		case "title", "description", "keywords", "logo", "favicon", "icp", "domain", "analytics", "theme", "copyright":
			data[cfg.Name] = cfg.Value
		}
	}

	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

const (
	ConfigSecurityIsClose                   = "is_close"                     // 是否关闭注册
	ConfigSecurityEnableRegister            = "enable_register"              // 是否允许注册
	ConfigSecurityEnableCaptchaLogin        = "enable_captcha_login"         // 是否开启登录验证码
	ConfigSecurityEnableCaptchaRegister     = "enable_captcha_register"      // 是否开启注册验证码
	ConfigSecurityEnableCaptchaComment      = "enable_captcha_comment"       // 是否开启注册验证码
	ConfigSecurityEnableCaptchaFindPassword = "enable_captcha_find_password" // 是否开启注册验证码
	ConfigSecurityEnableCaptchaUpload       = "enable_captcha_upload"        // 是否开启注册验证码
)

type ConfigSecurity struct {
	IsClose                   bool `json:"is_close"`                     // 是否闭站
	EnableRegister            bool `json:"enable_register"`              // 是否启用注册
	EnableCaptchaLogin        bool `json:"enable_captcha_login"`         // 是否启用登录验证码
	EnableCaptchaRegister     bool `json:"enable_captcha_register"`      // 是否启用注册验证码
	EnableCaptchaComment      bool `json:"enable_captcha_comment"`       // 是否启用评论验证码
	EnableCaptchaFindPassword bool `json:"enable_captcha_find_password"` // 找回密码是否需要验证码
	EnableCaptchaUpload       bool `json:"enable_captcha_upload"`        // 上传文档是否需要验证码
}

// GetConfigOfSecurity 获取安全配置
func (m *DBModel) GetConfigOfSecurity(name ...string) (config ConfigSecurity) {
	var configs []Config
	db := m.db.Where("category = ?", ConfigCategorySecurity)
	if len(name) > 0 {
		db = db.Where("name in (?)", name)
	}
	err := db.Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfSecurity", zap.Error(err))
	}

	var data = make(map[string]interface{})

	for _, cfg := range configs {
		switch cfg.Name {
		case "is_close", "enable_register", "enable_captcha_login", "enable_captcha_register", "enable_captcha_comment", "enable_captcha_find_password", "enable_captcha_upload":
			value, _ := strconv.ParseBool(cfg.Value)
			data[cfg.Name] = value
		}
	}

	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

func (m *DBModel) initConfig() (err error) {
	// 初始化配置项
	cfgs := []Config{
		// 系统配置项
		{Category: ConfigCategorySystem, Name: ConfigSystemTitle, Label: "网站名称", Value: "MOREDOC · 魔刀文库", Placeholder: "请输入您网站的名称", InputType: "text", Sort: 1, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemDescription, Label: "网站描述", Value: "MOREDOC · 魔刀文库", Placeholder: "请输入您网站的描述", InputType: "text", Sort: 2, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemKeywords, Label: "网站关键字", Value: "MOREDOC · 魔刀文库", Placeholder: "请输入您网站的关键字", InputType: "text", Sort: 3, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemLogo, Label: "网站Logo", Value: "", Placeholder: "请输入您网站的Logo", InputType: "text", Sort: 4, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemFavicon, Label: "网站Favicon", Value: "", Placeholder: "请输入您网站的Favicon", InputType: "text", Sort: 5, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemIcp, Label: "网站备案号", Value: "", Placeholder: "请输入您网站的备案号", InputType: "text", Sort: 6, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemDomain, Label: "网站域名", Value: "", Placeholder: "请输入您网站的域名", InputType: "textarea", Sort: 7, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemAnalytics, Label: "网站统计代码", Value: "", Placeholder: "请输入您网站的统计代码", InputType: "text", Sort: 8, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemTheme, Label: "网站主题", Value: "default", Placeholder: "请输入您网站的主题", InputType: "text", Sort: 9, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemCopyright, Label: "网站版权信息", Value: "", Placeholder: "请输入您网站的版权信息", InputType: "text", Sort: 10, Options: ""},

		// 验证码配置项
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaHeight, Label: "验证码高度", Value: "60", Placeholder: "请输入验证码高度，默认为60", InputType: "number", Sort: 13, Options: ""},
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaWidth, Label: "验证码宽度", Value: "240", Placeholder: "请输入验证码宽度，默认为240", InputType: "number", Sort: 14, Options: ""},
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaLength, Label: "验证码长度", Value: "5", Placeholder: "请输入验证码长度，默认为6", InputType: "number", Sort: 15, Options: ""},
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaType, Label: "验证码类型", Value: "digit", Placeholder: "请选择验证码类型，默认为数字", InputType: "select", Sort: 16, Options: captcha.CaptchaTypeOptions},

		// 安全配置项
		{Category: ConfigCategorySecurity, Name: ConfigSecurityIsClose, Label: "是否关闭网站", Value: "false", Placeholder: "请选择是否关闭网站", InputType: "swith", Sort: 17, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableRegister, Label: "是否允许注册", Value: "true", Placeholder: "请选择是否允许用户注册", InputType: "swith", Sort: 18, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaLogin, Label: "是否开启登录验证码", Value: "true", Placeholder: "请选择是否开启登录验证码", InputType: "swith", Sort: 19, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaRegister, Label: "是否开启注册验证码", Value: "true", Placeholder: "请选择是否开启注册验证码", InputType: "swith", Sort: 20, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaComment, Label: "是否开启评论验证码", Value: "true", Placeholder: "请选择是否开启评论验证码", InputType: "swith", Sort: 21, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaFindPassword, Label: "是否开启找回密码验证码", Value: "true", Placeholder: "请选择是否开启找回密码验证码", InputType: "swith", Sort: 22, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaUpload, Label: "是否开启文档上传验证码", Value: "true", Placeholder: "请选择是否开启文档上传验证码", InputType: "swith", Sort: 23, Options: ""},
	}

	for _, cfg := range cfgs {
		existConfig, _ := m.GetConfigByNameCategory(cfg.Name, cfg.Category, "id")
		if existConfig.Id > 0 {
			// 更新除了值之外的所有字段
			cfg.Id = existConfig.Id
			err = m.db.Omit("value").Updates(&cfg).Error
			if err != nil {
				m.logger.Error("initConfig", zap.Error(err))
				return
			}
			continue
		}
		err = m.CreateConfig(&cfg)
		if err != nil {
			m.logger.Error("initConfig", zap.Error(err))
			return
		}
	}
	return
}

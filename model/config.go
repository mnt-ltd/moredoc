package model

import (
	"moredoc/util/captcha"
	"strconv"
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
	// ConfigCategorySecurity 安全配置项
	ConfigCategorySecurity = "security"
	// ConfigCategoryFooter 底部链接
	ConfigCategoryFooter = "footer"
)

type Config struct {
	Id          int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Label       string     `form:"label" json:"label,omitempty" gorm:"column:label;type:varchar(64);size:64;comment:标签名称;"`
	Name        string     `form:"name" json:"name,omitempty" gorm:"column:name;type:varchar(64);size:64;index:name_category,unique;comment:表单字段名称;"`
	Value       string     `form:"value" json:"value,omitempty" gorm:"column:value;type:text;comment:值;"`
	Placeholder string     `form:"placeholder" json:"placeholder,omitempty" gorm:"column:placeholder;type:varchar(255);size:255;comment:提示信息;"`
	InputType   string     `form:"input_type" json:"input_type,omitempty" gorm:"column:input_type;type:varchar(32);size:32;default:text;comment:输入类型;"`
	Category    string     `form:"category" json:"category,omitempty" gorm:"column:category;type:varchar(32);size:32;index:name_category,unique;index:category;comment:所属类别;"`
	Sort        int        `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;comment:同一category下的排序，这里按顺序排序，值越小越靠前;"`
	Options     string     `form:"options" json:"options,omitempty" gorm:"column:options;type:text;comment:针对checkbox等的枚举值;"`
	CreatedAt   *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt   *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

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

// UpdateConfigs 配置项批量更新
// TODO: value值为6个*的，需要特殊处理
func (m *DBModel) UpdateConfigs(configs []*Config, updateFields ...string) (err error) {
	sess := m.db.Begin()
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()

	tableName := Config{}.TableName()
	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) == 0 {
		updateFields = m.GetTableFields(tableName)
	}

	for _, config := range configs {
		if err = sess.Select(updateFields).Save(config).Error; err != nil {
			m.logger.Error("UpdateConfigs", zap.Error(err))
			return
		}
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
	SelectFields []string                 // 查询字段
	QueryIn      map[string][]interface{} // map[field][]{value1,value2,...}
}

// GetConfigList 获取Config列表
func (m *DBModel) GetConfigList(opt *OptionGetConfigList) (configList []Config, err error) {
	db := m.db.Model(&Config{})
	db = m.generateQueryIn(db, Config{}.TableName(), opt.QueryIn)
	err = db.Order("sort asc").Find(&configList).Error
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
	ConfigSystemSitename           = "sitename"
	ConfigSystemDomain             = "domain"
	ConfigSystemTitle              = "title"
	ConfigSystemDescription        = "description"
	ConfigSystemKeywords           = "keywords"
	ConfigSystemLogo               = "logo"
	ConfigSystemFavicon            = "favicon"
	ConfigSystemIcp                = "icp"
	ConfigSystemAnalytics          = "analytics"
	ConfigSystemCopyrightStartYear = "copyright_start_year"
)

type ConfigSystem struct {
	Sitename           string `json:"sitename"`             // 网站名称
	Domain             string `json:"domain"`               // 站点域名，不带 HTTPS:// 和 HTTP://
	Title              string `json:"title"`                // 网站首页标题
	Keywords           string `json:"keywords"`             // 系统关键字
	Description        string `json:"description"`          // 系统描述
	Logo               string `json:"logo"`                 // logo
	Favicon            string `json:"favicon"`              // logo
	ICP                string `json:"icp"`                  // 网站备案
	Analytics          string `json:"analytics"`            // 统计代码
	CopyrightStartYear string `json:"copyright_start_year"` // 版权年
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
		data[cfg.Name] = cfg.Value
	}

	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

const (
	ConfigSecurityMaxDocumentSize           = "max_document_size"            // 是否关闭注册
	ConfigSecurityIsClose                   = "is_close"                     // 是否关闭注册
	ConfigSecurityCloseStatement            = "close_statement"              // 闭站说明
	ConfigSecurityEnableRegister            = "enable_register"              // 是否允许注册
	ConfigSecurityEnableCaptchaLogin        = "enable_captcha_login"         // 是否开启登录验证码
	ConfigSecurityEnableCaptchaRegister     = "enable_captcha_register"      // 是否开启注册验证码
	ConfigSecurityEnableCaptchaComment      = "enable_captcha_comment"       // 是否开启注册验证码
	ConfigSecurityEnableCaptchaFindPassword = "enable_captcha_find_password" // 是否开启注册验证码
)

type ConfigSecurity struct {
	MaxDocumentSize           int32  `json:"max_document_size"`            // 允许上传的最大文档大小
	IsClose                   bool   `json:"is_close"`                     // 是否闭站
	CloseStatement            string `json:"close_statement"`              // 闭站说明
	EnableRegister            bool   `json:"enable_register"`              // 是否启用注册
	EnableCaptchaLogin        bool   `json:"enable_captcha_login"`         // 是否启用登录验证码
	EnableCaptchaRegister     bool   `json:"enable_captcha_register"`      // 是否启用注册验证码
	EnableCaptchaComment      bool   `json:"enable_captcha_comment"`       // 是否启用评论验证码
	EnableCaptchaFindPassword bool   `json:"enable_captcha_find_password"` // 找回密码是否需要验证码
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
		case "max_document_size":
			data[cfg.Name], _ = strconv.Atoi(cfg.Value)
		default:
			data[cfg.Name] = cfg.Value
		}
	}

	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

const (
	ConfigFooterAbout     = "about"     // 关于我们
	ConfigFooterContact   = "contact"   // 联系我们
	ConfigFooterAgreement = "agreement" // 用户协议
	ConfigFooterCopyright = "copyright" // 版权信息
	ConfigFooterFeedback  = "feedback"  // 反馈信息
)

type ConfigFooter struct {
	About     string `json:"about"`     // 关于我们
	Contact   string `json:"contact"`   // 联系我们
	Agreement string `json:"agreement"` // 用户协议、文库协议
	Copyright string `json:"copyright"` // 版权信息、免责声明
	Feedback  string `json:"feedback"`  // 反馈
}

func (m *DBModel) GetConfigOfFooter() (config ConfigFooter) {
	var configs []Config
	err := m.db.Where("category = ?", ConfigCategoryFooter).Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfFooter", zap.Error(err))
	}

	var data = make(map[string]interface{})

	for _, cfg := range configs {
		data[cfg.Name] = cfg.Value
	}

	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

func (m *DBModel) initConfig() (err error) {
	// 初始化配置项
	cfgs := []Config{
		// 系统配置项
		{Category: ConfigCategorySystem, Name: ConfigSystemSitename, Label: "网站名称", Value: "魔刀文库", Placeholder: "请输入您网站的名称，如：魔刀文库", InputType: "text", Sort: 1, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemTitle, Label: "首页标题", Value: "MOREDOC · 魔刀文库", Placeholder: "请输入您网站的首页标题，如：魔刀文库，强大、专业的文库系统", InputType: "text", Sort: 2, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemKeywords, Label: "网站关键字", Value: "MOREDOC · 魔刀文库", Placeholder: "请输入您网站的关键字", InputType: "text", Sort: 3, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemDescription, Label: "网站描述", Value: "MOREDOC · 魔刀文库", Placeholder: "请输入您网站的描述", InputType: "textarea", Sort: 4, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemLogo, Label: "网站Logo", Value: "", Placeholder: "请输入您网站的Logo路径", InputType: "image", Sort: 6, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemFavicon, Label: "网站Favicon", Value: "", Placeholder: "请输入您网站的Favicon路径", InputType: "image", Sort: 6, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemIcp, Label: "网站备案号", Value: "", Placeholder: "请输入您网站的备案号", InputType: "text", Sort: 6, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemDomain, Label: "网站域名", Value: "", Placeholder: "请输入您网站的域名访问地址，如 https://moredoc.mnt.ltd，用以生成网站地图sitemap", InputType: "text", Sort: 7, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemCopyrightStartYear, Label: "版权起始年", Value: "2019", Placeholder: "请输入您网站版权起始年，如：2019，则前台会显示如 ©2019 - 2022 的字样", InputType: "text", Sort: 8, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemAnalytics, Label: "网站统计代码", Value: "", Placeholder: "请输入您网站的统计代码", InputType: "textarea", Sort: 9, Options: ""},

		// 验证码配置项
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaHeight, Label: "验证码高度", Value: "60", Placeholder: "请输入验证码高度，默认为60", InputType: "number", Sort: 13, Options: ""},
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaWidth, Label: "验证码宽度", Value: "240", Placeholder: "请输入验证码宽度，默认为240", InputType: "number", Sort: 14, Options: ""},
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaLength, Label: "验证码长度", Value: "5", Placeholder: "请输入验证码长度，默认为6", InputType: "number", Sort: 15, Options: ""},
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaType, Label: "验证码类型", Value: "digit", Placeholder: "请选择验证码类型，默认为数字", InputType: "select", Sort: 16, Options: captcha.CaptchaTypeOptions},

		// 安全配置项
		{Category: ConfigCategorySecurity, Name: ConfigSecurityMaxDocumentSize, Label: "最大文档大小(MB)", Value: "50", Placeholder: "允许用户上传的最大文档大小，默认为50，即50MB", InputType: "number", Sort: 15, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityIsClose, Label: "是否关闭网站", Value: "false", Placeholder: "请选择是否关闭网站", InputType: "switch", Sort: 16, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityCloseStatement, Label: "闭站说明", Value: "false", Placeholder: "关闭网站后，页面提示的内容", InputType: "textarea", Sort: 17, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableRegister, Label: "是否允许注册", Value: "true", Placeholder: "请选择是否允许用户注册", InputType: "switch", Sort: 18, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaLogin, Label: "是否开启登录验证码", Value: "true", Placeholder: "请选择是否开启登录验证码", InputType: "switch", Sort: 19, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaRegister, Label: "是否开启注册验证码", Value: "true", Placeholder: "请选择是否开启注册验证码", InputType: "switch", Sort: 20, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaComment, Label: "是否开启评论验证码", Value: "true", Placeholder: "请选择是否开启评论验证码", InputType: "switch", Sort: 21, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaFindPassword, Label: "是否开启找回密码验证码", Value: "true", Placeholder: "请选择是否开启找回密码验证码", InputType: "switch", Sort: 22, Options: ""},

		// 底部链接
		{Category: ConfigCategoryFooter, Name: ConfigFooterAbout, Label: "关于我们", Value: "", Placeholder: "请输入关于我们的链接地址，留空表示不显示", InputType: "text", Sort: 24, Options: ""},
		{Category: ConfigCategoryFooter, Name: ConfigFooterContact, Label: "联系我们", Value: "", Placeholder: "请输入联系我们的链接地址，留空表示不显示", InputType: "text", Sort: 25, Options: ""},
		{Category: ConfigCategoryFooter, Name: ConfigFooterAgreement, Label: "文库协议", Value: "", Placeholder: "请输入文库协议的链接地址，留空表示不显示", InputType: "text", Sort: 26, Options: ""},
		{Category: ConfigCategoryFooter, Name: ConfigFooterCopyright, Label: "免责声明", Value: "", Placeholder: "请输入免责声明的链接地址，留空表示不显示", InputType: "text", Sort: 27, Options: ""},
		{Category: ConfigCategoryFooter, Name: ConfigFooterFeedback, Label: "意见反馈", Value: "", Placeholder: "请输入意见反馈的链接地址，留空表示不显示", InputType: "text", Sort: 28, Options: ""},
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

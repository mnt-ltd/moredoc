package model

import (
	"crypto/tls"
	"errors"
	"moredoc/util/captcha"
	"moredoc/util/filetil"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
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
	// ConfigCategoryConverter 转换配置项
	ConfigCategoryConverter = "converter"
	// 下载配置
	ConfigCategoryDownload = "download"
	// 积分规则
	ConfigCategoryScore = "score"
	// 显示配置
	ConfigCategoryDisplay = "display"
)

const (
	InputTypeNumber      = "number"       // 数字
	InputTypeSwitch      = "switch"       // switch
	InputTypeSelect      = "select"       // 单选
	InputTypeSelectMulti = "select-multi" // 多选
	InputTypeText        = "text"         // 文本框
	InputTypeTextarea    = "textarea"     // 多行文本框
	InputTypePassword    = "password"     // 密码
	InputTypeImage       = "image"        // 图片
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
		m.logger.Debug("UpdateConfigs", zap.Any("config", config), zap.Any("updateFields", updateFields))
		if err = sess.Select(updateFields).Updates(config).Error; err != nil {
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

const (
	ConfigSystemSitename            = "sitename"
	ConfigSystemDomain              = "domain"
	ConfigSystemTitle               = "title"
	ConfigSystemDescription         = "description"
	ConfigSystemKeywords            = "keywords"
	ConfigSystemRecommendWords      = "recommend_words"
	ConfigSystemLogo                = "logo"
	ConfigSystemFavicon             = "favicon"
	ConfigSystemLoginBackground     = "login_background"
	ConfigSystemRegistrerBackground = "register_background"
	ConfigSystemIcp                 = "icp"
	ConfigSystemSecIcp              = "sec_icp"
	ConfigSystemAnalytics           = "analytics"
	ConfigSystemCopyrightStartYear  = "copyright_start_year"
)

type ConfigSystem struct {
	Sitename                        string   `json:"sitename"`             // 网站名称
	Domain                          string   `json:"domain"`               // 站点域名，带 HTTPS:// 或 HTTP://
	Title                           string   `json:"title"`                // 网站首页标题
	Keywords                        string   `json:"keywords"`             // 系统关键字
	Description                     string   `json:"description"`          // 系统描述
	Logo                            string   `json:"logo"`                 // logo
	Favicon                         string   `json:"favicon"`              // favicon
	ConfigSystemRegistrerBackground string   `json:"register_background"`  // 注册页面背景图
	ConfigSystemLoginBackground     string   `json:"login_background"`     // 登录页面背景图
	SecICP                          string   `json:"sec_icp"`              // 网站备案
	ICP                             string   `json:"icp"`                  // 网站备案
	Analytics                       string   `json:"analytics"`            // 统计代码
	CopyrightStartYear              string   `json:"copyright_start_year"` // 版权年
	RecommendWords                  []string `json:"recommend_words"`      // 推荐词，首页收缩推荐词
}

const (
	ConfigSecurityMaxDocumentSize           = "max_document_size"            // 是否关闭注册
	ConfigSecurityCommentInterval           = "comment_interval"             // 评论时间间隔
	ConfigSecurityIsClose                   = "is_close"                     // 是否关闭注册
	ConfigSecurityCloseStatement            = "close_statement"              // 闭站说明
	ConfigSecurityEnableRegister            = "enable_register"              // 是否允许注册
	ConfigSecurityEnableCaptchaLogin        = "enable_captcha_login"         // 是否开启登录验证码
	ConfigSecurityEnableCaptchaRegister     = "enable_captcha_register"      // 是否开启注册验证码
	ConfigSecurityEnableCaptchaComment      = "enable_captcha_comment"       // 是否开启注册验证码
	ConfigSecurityEnableCaptchaFindPassword = "enable_captcha_find_password" // 是否开启注册验证码
	ConfigSecurityDocumentRelatedDuration   = "document_related_duration"    // 相关文档有效期，默认为7天，最小值为1
	ConfigSecurityDocumentAllowedExt        = "document_allowed_ext"         // 允许上传的文档类型
	ConfigSecurityLoginRequired             = "login_required"               // 是否需要登录才能查看文档
	ConfigSecurityAttachmentRetentionMinute = "attachment_retention_minute"  // 被删除的附件，保留时长，单位分钟
)

type ConfigSecurity struct {
	MaxDocumentSize           int32    `json:"max_document_size"`            // 允许上传的最大文档大小
	CommentInterval           int32    `json:"comment_interval"`             // 评论时间间隔, 单位秒
	DocumentRelatedDuration   int32    `json:"document_related_duration"`    // 相关文档有效期，默认为7天，最小值为1
	AttachmentRetentionMinute int32    `json:"attachment_retention_minute"`  // 被删除的附件，保留时长，单位分钟
	IsClose                   bool     `json:"is_close"`                     // 是否闭站
	CloseStatement            string   `json:"close_statement"`              // 闭站说明
	EnableRegister            bool     `json:"enable_register"`              // 是否启用注册
	EnableCaptchaLogin        bool     `json:"enable_captcha_login"`         // 是否启用登录验证码
	EnableCaptchaRegister     bool     `json:"enable_captcha_register"`      // 是否启用注册验证码
	EnableCaptchaComment      bool     `json:"enable_captcha_comment"`       // 是否启用评论验证码
	EnableCaptchaFindPassword bool     `json:"enable_captcha_find_password"` // 找回密码是否需要验证码
	DocumentAllowedExt        []string `json:"document_allowed_ext"`         // 允许上传的文档类型
	LoginRequired             bool     `json:"login_required"`               // 是否需要登录才能查看文档
}

const (
	ConfigConverterMaxPreview                    = "max_preview"                      // 最大预览页数
	ConfigConverterTimeout                       = "timeout"                          // 转换超时时间
	ConfigConverterEnableSVGO                    = "enable_svgo"                      // 是否启用 SVGO
	ConfigConverterEnableGZIP                    = "enable_gzip"                      // 是否启用 GZIP
	ConfigConverterEnableConvertRepeatedDocument = "enable_convert_repeated_document" // 是否转换已转换的重复文档
)

const (
	// 是否允许游客下载
	ConfigDownloadEnableGuestDownload = "enable_guest_download"
	// 每天允许下载的次数
	ConfigDownloadTimesEveryDay = "times_every_day"
	// 下载链接地址签名密钥
	ConfigDownloadSecretKey = "secret_key"
	// 生成的下载链接有效期，单位为秒
	ConfigDownloadUrlDuration = "url_duration"
	// 购买文档后多少天内允许免费重复下载
	ConfigDownloadFreeDownloadDuration = "free_download_duration"
	//  允许每个IP每天下载的次数
	ConfigDownloadTimesEveryIP = "times_every_ip"
)

type ConfigDownload struct {
	EnableGuestDownload  bool   `json:"enable_guest_download"`  // 是否允许游客下载(不需要登录，针对免费文档)
	TimesEveryDay        int32  `json:"times_every_day"`        // 允许用户每天下载的次数（针对登录用户）
	TimesEveryIP         int32  `json:"times_every_ip"`         // 允许每个IP每天下载的次数（针对所有用户）
	SecretKey            string `json:"secret_key"`             // 下载链接地址签名密钥
	UrlDuration          int32  `json:"url_duration"`           // 生成的下载链接有效期，单位为秒
	FreeDownloadDuration int32  `json:"free_download_duration"` // 购买文档后多少天内允许免费重复下载
}

// ConfigConverter 转换配置
type ConfigConverter struct {
	MaxPreview                    int  `json:"max_preview"`                      // 文档所允许的最大预览页数，0 表示不限制，全部转换
	Timeout                       int  `json:"timeout"`                          // 转换超时时间，单位为分钟，默认30分钟
	EnableSVGO                    bool `json:"enable_svgo"`                      // 是否对svg启用SVGO压缩。转换效率会有所下降。相对直接的svg文件，可以节省1/2的存储空间
	EnableGZIP                    bool `json:"enable_gzip"`                      // 是否对svg启用GZIP压缩。转换效率会有所下降。相对直接的svg文件，可以节省3/4的存储空间
	EnableConvertRepeatedDocument bool `json:"enable_convert_repeated_document"` // 是否转换已转换的重复文档。如果开启，会导致转换效率下降，但是可以节省大量的存储空间
	// GZIP和svgo都开启，转换效率会有所下降，可以综合节省约85%的存储空间
}

const (
	ConfigEmailEnable    = "enable" // 是否启用邮件服务
	ConfigEmailHost      = "host"   // SMTP 服务器地址
	ConfigEmailPort      = "port"   // SMTP 服务器端口
	ConfigEmailIsTLS     = "is_tls" // 是否启用TLS
	ConfigEmailFromName  = "from_name"
	ConfigEmailUsername  = "username" // SMTP 用户名
	ConfigEmailPassword  = "password" // SMTP 密码
	ConfigEmailDuration  = "duration" // 验证码有效期，单位为分钟
	ConfigEmailTestEmail = "test_email"
	ConfigEmailReplyTo   = "reply_to"
	ConfigEmailSecret    = "secret" // 找回密码邮件的签名密钥
)

type ConfigEmail struct {
	Enable    bool   `json:"enable"` // 是否启用邮件服务
	Host      string `json:"host"`   // SMTP 服务器地址
	Port      int    `json:"port"`   // SMTP 服务器端口
	IsTLS     bool   `json:"is_tls"` // 是否启用TLS
	FromName  string `json:"from_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Duration  int    `json:"duration"` // 验证码有效期，单位为分钟
	TestEmail string `json:"test_email"`
	Secret    string `json:"secret"`
	// ReplyTo   string `json:"reply_to"`
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

const (
	// 积分名称
	ConfigScoreCreditName = "credit_name"
	// 注册积分
	ConfigScoreRegister = "register"
	// 签到积分
	ConfigScoreSignIn = "sign_in"
	// 上传文档积分
	ConfigScoreUploadDocument = "upload_document"
	// 每日上传文档积分次数限制
	ConfigScoreUploadDocumentLimit = "upload_document_limit"
	// 删除上传文档积分
	ConfigScoreDeleteDocument = "delete_document"
	// 文档被收藏获得积分
	ConfigScoreDocumentCollected = "document_collected"
	// 文档被收藏获得积分次数限制
	ConfigScoreDocumentCollectedLimit = "document_collected_limit"
	// 文档被评论获得积分
	ConfigScoreDocumentCommented = "document_commented"
	// 文档被评论获得积分次数限制
	ConfigScoreDocumentCommentedLimit = "document_commented_limit"
)

type ConfigScore struct {
	CreditName             string `json:"credit_name"`              // 积分名称
	Register               int32  `json:"register"`                 // 注册积分
	SignIn                 int32  `json:"sign_in"`                  // 签到积分
	UploadDocument         int32  `json:"upload_document"`          // 上传文档积分
	UploadDocumentLimit    int32  `json:"upload_document_limit"`    // 每日上传文档积分次数限制
	DeleteDocument         int32  `json:"delete_document"`          // 删除上传文档积分
	DocumentCollected      int32  `json:"document_collected"`       // 文档被收藏获得积分
	DocumentCollectedLimit int32  `json:"document_collected_limit"` // 文档被收藏获得积分次数限制
	DocumentCommented      int32  `json:"document_commented"`       // 文档被评论获得积分
	DocumentCommentedLimit int32  `json:"document_commented_limit"` // 文档被评论获得积分次数限制
}

const (
	ConfigDisplayShowRegisterUserCount = "show_register_user_count" // 是否显示注册用户数量
	ConfigDisplayVirtualRegisterCount  = "virtual_register_count"   // 虚拟注册用户数量
	ConfigDisplayShowIndexCategories   = "show_index_categories"    // 是否显示首页分类
	ConfigDisplayPagesPerRead          = "pages_per_read"           // 每次阅读的页数
	ConfigDisplayCopyrightStatement    = "copyright_statement"      // 在页面最底部的版权声明
	ConfigDisplayMaxSearchPages        = "max_search_pages"         // 搜索结果最大页数
)

type ConfigDisplay struct {
	ShowRegisterUserCount bool   `json:"show_register_user_count"` // 是否显示注册用户数量
	VirtualRegisterCount  int64  `json:"virtual_register_count"`   // 虚拟注册用户数量
	ShowIndexCategories   bool   `json:"show_index_categories"`    // 是否显示首页分类
	PagesPerRead          int32  `json:"pages_per_read"`           // 每次阅读的页数
	MaxSearchPages        int32  `json:"max_search_pages"`         // 搜索结果最大页数
	CopyrightStatement    string `json:"copyright_statement"`      // 在页面最底部的版权声明
}

func (m *DBModel) GetConfigOfDisplay(name ...string) (config ConfigDisplay) {
	var configs []Config

	db := m.db
	if len(name) > 0 {
		db = db.Where("name IN (?)", name)
	}
	err := db.Where("category = ?", ConfigCategoryDisplay).Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfDisplay", zap.Error(err))
	}

	data := m.convertConfig2Map(configs)

	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

func (m *DBModel) GetConfigOfFooter() (config ConfigFooter) {
	var configs []Config
	err := m.db.Where("category = ?", ConfigCategoryFooter).Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfFooter", zap.Error(err))
	}

	data := m.convertConfig2Map(configs)

	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

func (m *DBModel) GetConfigOfDownload(name ...string) (config ConfigDownload) {
	var configs []Config

	db := m.db
	if len(name) > 0 {
		db = db.Where("name IN (?)", name)
	}
	err := db.Where("category = ?", ConfigCategoryDownload).Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfDownload", zap.Error(err))
	}

	data := m.convertConfig2Map(configs)
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

// GetConfigOfCaptcha 获取验证码配置
func (m *DBModel) GetConfigOfCaptcha() (config ConfigCaptcha) {
	var configs []Config
	err := m.db.Where("category = ?", ConfigCategoryCaptcha).Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfCaptcha", zap.Error(err))
	}
	data := m.convertConfig2Map(configs)
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	if config.Length <= 0 {
		config.Length = 6
	}
	if config.Width <= 0 {
		config.Width = 240
	}
	if config.Height <= 0 {
		config.Height = 60
	}
	return
}

// GetConfigOfSystem 获取系统配置
func (m *DBModel) GetConfigOfSystem(name ...string) (config ConfigSystem) {
	var confgis []Config
	db := m.db.Where("category = ?", ConfigCategorySystem)
	if len(name) > 0 {
		db = db.Where("name IN (?)", name)
	}
	err := db.Find(&confgis).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfSystem", zap.Error(err))
	}

	data := m.convertConfig2Map(confgis)
	// 注意：推荐的关键字，要特殊处理下
	if recommendWords, ok := data[ConfigSystemRecommendWords]; ok {
		var words []string
		if rws, ok := recommendWords.(string); ok {
			for _, word := range strings.Split(rws, ",") {
				word = strings.TrimSpace(word)
				if word == "" {
					continue
				}
				words = append(words, word)
			}
		}
		data[ConfigSystemRecommendWords] = words
	}
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)
	return
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

	data := m.convertConfig2Map(configs)
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

func (m *DBModel) GetConfigOfConverter() (config ConfigConverter) {
	var configs []Config
	err := m.db.Where("category = ?", ConfigCategoryConverter).Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfConverter", zap.Error(err))
	}

	data := m.convertConfig2Map(configs)
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

func (m *DBModel) GetConfigOfEmail(name ...string) (config ConfigEmail) {
	var configs []Config
	db := m.db.Where("category = ?", ConfigCategoryEmail)
	if len(name) > 0 {
		db = db.Where("name in (?)", name)
	}
	err := db.Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfEmail", zap.Error(err))
	}

	data := m.convertConfig2Map(configs)
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)
	m.logger.Debug("GetConfigOfEmail", zap.Any("data", data), zap.Any("config", config))

	return
}

func (m *DBModel) GetConfigOfScore(name ...string) (config ConfigScore) {
	var configs []Config
	db := m.db.Where("category = ?", ConfigCategoryScore)
	if len(name) > 0 {
		db = db.Where("name in (?)", name)
	}
	err := db.Find(&configs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetConfigOfScore", zap.Error(err))
	}

	data := m.convertConfig2Map(configs)
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &config)

	return
}

func (m *DBModel) SendMail(subject, email string, body string) error {
	cfg := m.GetConfigOfEmail()
	m.logger.Debug("SendMail", zap.Any("cfg", cfg), zap.String("email", email), zap.String("subject", subject), zap.String("body", body))
	if !cfg.Enable {
		return errors.New("邮件服务未启用")
	}
	fromName := cfg.Username
	if fn := strings.TrimSpace(cfg.FromName); fn != "" {
		fromName = fn
	}
	message := gomail.NewMessage()
	message.SetHeader("From", message.FormatAddress(cfg.Username, fromName))
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)
	// if cfg.ReplyTo != "" {
	// 	message.SetHeader("Reply-To", cfg.ReplyTo)
	// }

	mail := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	if cfg.IsTLS {
		mail.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return mail.DialAndSend(message)
}
func (m *DBModel) GetCreditName() string {
	creditName := m.GetConfigOfScore(ConfigScoreCreditName).CreditName
	if creditName == "" {
		creditName = "魔豆"
	}
	return creditName
}

func (m *DBModel) convertConfig2Map(cfgs []Config) (data map[string]interface{}) {
	data = make(map[string]interface{})
	for _, cfg := range cfgs {
		switch strings.TrimSpace(cfg.InputType) {
		case InputTypeNumber:
			data[cfg.Name], _ = strconv.Atoi(cfg.Value)
		case InputTypeSwitch:
			data[cfg.Name], _ = strconv.ParseBool(cfg.Value)
		case InputTypeSelectMulti:
			arr := strings.Split(cfg.Value, ",")
			if len(arr) == 1 && arr[0] == "" {
				arr = []string{}
			}
			data[cfg.Name] = arr
		case InputTypeSelect, InputTypeText, InputTypeTextarea, InputTypePassword, InputTypeImage:
			data[cfg.Name] = cfg.Value
		default:
			// 这里只是做一个兼容处理。报错是为了提醒开发者，如果有新的输入类型，需要增加 InptType 的枚举常量，并在相应的表单进行处理
			m.logger.Error("convertConfig2Map", zap.Error(errors.New("未知的输入类型，请在枚举常量InputType中进行定义")), zap.String("inputType", cfg.InputType), zap.String("name", cfg.Name), zap.String("value", cfg.Value))
			data[cfg.Name] = cfg.Value
		}
	}
	return
}

func (m *DBModel) initConfig() (err error) {
	closeStatement := `<div>尊敬的用户，您好：</div>
	<div>为了给您带来更好的使用体验，<strong>魔豆文库</strong> 正在对服务进行升级维护，预计恢复时间为 <span style="color:red">2024-10-24 06:00:00</span>，请您稍后再进行访问。</div>
	<div>升级维护期间，普通用户将无法正常使用（系统管理人员出于维护的需要不受升级影响）。</div>
	<div>由此带来的不便，敬请谅解。</div>`
	// 初始化配置项
	cfgs := []Config{
		// 系统配置项
		{Category: ConfigCategorySystem, Name: ConfigSystemSitename, Label: "网站名称", Value: "魔豆文库", Placeholder: "请输入您网站的名称，如：魔豆文库", InputType: InputTypeText, Sort: 10, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemTitle, Label: "首页标题", Value: "MOREDOC · 魔豆文库", Placeholder: "请输入您网站的首页标题，如：魔豆文库，强大、专业的文库系统", InputType: InputTypeText, Sort: 20, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemKeywords, Label: "网站关键字", Value: "MOREDOC · 魔豆文库", Placeholder: "请输入您网站的关键字", InputType: InputTypeText, Sort: 30, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemDescription, Label: "网站描述", Value: "MOREDOC · 魔豆文库", Placeholder: "请输入您网站的描述", InputType: InputTypeTextarea, Sort: 40, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemRecommendWords, Label: "首页搜索推荐词", Value: "", Placeholder: "网站首页搜索推荐关键字，多个关键字用英文逗号分隔", InputType: InputTypeText, Sort: 50, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemLogo, Label: "网站Logo", Value: "", Placeholder: "请上传一张图片作为网站Logo", InputType: InputTypeImage, Sort: 60, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemFavicon, Label: "网站Favicon", Value: "", Placeholder: "请上传一张方方正正的小图片作为网站favicon，建议为 .ico 的图片", InputType: InputTypeImage, Sort: 61, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemRegistrerBackground, Label: "注册页背景图", Value: "", Placeholder: "请上传一张图片作为注册页背景图", InputType: InputTypeImage, Sort: 62, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemLoginBackground, Label: "登录页背景图", Value: "", Placeholder: "请上传一张图片作为登录页背景图", InputType: InputTypeImage, Sort: 63, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemSecIcp, Label: "京公网安备", Value: "", Placeholder: "请输入您网站的京公网安备备案号", InputType: InputTypeText, Sort: 68, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemIcp, Label: "网站备案号", Value: "", Placeholder: "请输入您网站的备案号", InputType: InputTypeText, Sort: 69, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemDomain, Label: "网站域名", Value: "https://moredoc.mnt.ltd", Placeholder: "请输入您网站的域名访问地址，带 https:// 或 http:// 如 https://moredoc.mnt.ltd，用以生成网站地图sitemap", InputType: InputTypeText, Sort: 70, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemCopyrightStartYear, Label: "版权起始年", Value: "2019", Placeholder: "请输入您网站版权起始年，如：2019，则前台会显示如 ©2019 - 20xx 的字样", InputType: InputTypeText, Sort: 80, Options: ""},
		{Category: ConfigCategorySystem, Name: ConfigSystemAnalytics, Label: "网站统计代码", Value: "", Placeholder: "请输入您网站的统计代码，当前只支持百度统计", InputType: InputTypeTextarea, Sort: 90, Options: ""},

		// 验证码配置项
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaHeight, Label: "验证码高度", Value: "60", Placeholder: "请输入验证码高度，默认为60", InputType: InputTypeNumber, Sort: 13, Options: ""},
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaWidth, Label: "验证码宽度", Value: "240", Placeholder: "请输入验证码宽度，默认为240", InputType: InputTypeNumber, Sort: 14, Options: ""},
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaLength, Label: "验证码长度", Value: "5", Placeholder: "请输入验证码长度，默认为6", InputType: InputTypeNumber, Sort: 15, Options: ""},
		{Category: ConfigCategoryCaptcha, Name: ConfigCaptchaType, Label: "验证码类型", Value: "digit", Placeholder: "请选择验证码类型，默认为数字", InputType: InputTypeSelect, Sort: 16, Options: captcha.CaptchaTypeOptions},

		// 安全配置项
		{Category: ConfigCategorySecurity, Name: ConfigSecurityMaxDocumentSize, Label: "最大文档大小(MB)", Value: "50", Placeholder: "允许用户上传的最大文档大小，默认为50，即50MB。配置时仍需配置反向代理服务所允许传输的最大大小，如nginx的client_max_body_size值！", InputType: InputTypeNumber, Sort: 1, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityCommentInterval, Label: "评论时间间隔", Value: "10", Placeholder: "用户评论时间间隔，单位为秒。0表示不限制。", InputType: InputTypeNumber, Sort: 2, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityDocumentRelatedDuration, Label: "文档的【相关文档】有效期", Value: "7", Placeholder: "文档的相关联文档的有效期，默认为7，即7天，0或小于0，表示不开启相关文档功能", InputType: InputTypeNumber, Sort: 15, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityAttachmentRetentionMinute, Label: "从回收站清除的文档文件保留时长", Value: "1440", Placeholder: "单位为分钟，默认为1440，即24小时。文档被从回收站清除之后，文档文件并未真正删除，用以规避操作失误的情况。设为 0 则表示即时删除文档文件，用以释放存储空间。", InputType: InputTypeNumber, Sort: 16, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityDocumentAllowedExt, Label: "允许上传的文档类型", Value: "", Placeholder: "留空表示允许程序所支持的全部文档类型", InputType: InputTypeSelectMulti, Sort: 30, Options: strings.Join(filetil.GetDocumentExts(), "\n")},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityIsClose, Label: "是否关闭网站", Value: "false", Placeholder: "请选择是否关闭网站", InputType: InputTypeSwitch, Sort: 160, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityCloseStatement, Label: "闭站说明", Value: closeStatement, Placeholder: "关闭网站后，页面提示内容（支持HTML）", InputType: InputTypeTextarea, Sort: 170, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableRegister, Label: "是否允许注册", Value: "true", Placeholder: "请选择是否允许用户注册", InputType: InputTypeSwitch, Sort: 18, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityLoginRequired, Label: "是否登录才能访问", Value: "false", Placeholder: "用户是否登录了才能访问", InputType: InputTypeSwitch, Sort: 18, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaLogin, Label: "是否开启登录验证码", Value: "true", Placeholder: "请选择是否开启登录验证码", InputType: InputTypeSwitch, Sort: 19, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaRegister, Label: "是否开启注册验证码", Value: "true", Placeholder: "请选择是否开启注册验证码", InputType: InputTypeSwitch, Sort: 20, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaComment, Label: "是否开启评论验证码", Value: "true", Placeholder: "请选择是否开启评论验证码", InputType: InputTypeSwitch, Sort: 21, Options: ""},
		{Category: ConfigCategorySecurity, Name: ConfigSecurityEnableCaptchaFindPassword, Label: "是否开启找回密码验证码", Value: "true", Placeholder: "请选择是否开启找回密码验证码", InputType: InputTypeSwitch, Sort: 22, Options: ""},

		// 底部链接
		{Category: ConfigCategoryFooter, Name: ConfigFooterAbout, Label: "关于我们", Value: "/article/about", Placeholder: "请输入关于我们的链接地址，留空表示不显示", InputType: InputTypeText, Sort: 24, Options: ""},
		{Category: ConfigCategoryFooter, Name: ConfigFooterContact, Label: "联系我们", Value: "/article/contact", Placeholder: "请输入联系我们的链接地址，留空表示不显示", InputType: InputTypeText, Sort: 25, Options: ""},
		{Category: ConfigCategoryFooter, Name: ConfigFooterAgreement, Label: "文库协议", Value: "/article/agreement", Placeholder: "请输入文库协议的链接地址，留空表示不显示", InputType: InputTypeText, Sort: 26, Options: ""},
		{Category: ConfigCategoryFooter, Name: ConfigFooterCopyright, Label: "免责声明", Value: "/article/copyright", Placeholder: "请输入免责声明的链接地址，留空表示不显示", InputType: InputTypeText, Sort: 27, Options: ""},
		{Category: ConfigCategoryFooter, Name: ConfigFooterFeedback, Label: "意见反馈", Value: "/article/feedback", Placeholder: "请输入意见反馈的链接地址，留空表示不显示", InputType: InputTypeText, Sort: 28, Options: ""},

		// 转换配置项
		{Category: ConfigCategoryConverter, Name: ConfigConverterMaxPreview, Label: "最大预览页数", Value: "0", Placeholder: "文档允许的最大预览页数，0表示不限制", InputType: InputTypeNumber, Sort: 15, Options: ""},
		{Category: ConfigCategoryConverter, Name: ConfigConverterTimeout, Label: "转换超时(分钟)", Value: "30", Placeholder: "文档转换超时时间，默认为30分钟", InputType: InputTypeNumber, Sort: 16, Options: ""},
		{Category: ConfigCategoryConverter, Name: ConfigConverterEnableGZIP, Label: "是否启用GZIP压缩", Value: "true", Placeholder: "是否对文档SVG预览文件启用GZIP压缩，启用后转换效率会【稍微】下降，但相对直接的SVG文件减少75%的存储空间", InputType: InputTypeSwitch, Sort: 17, Options: ""},
		{Category: ConfigCategoryConverter, Name: ConfigConverterEnableSVGO, Label: "是否启用SVGO", Value: "false", Placeholder: "是否对文档SVG预览文件启用SVGO压缩，启用后转换效率会【明显】下降，但相对直接的SVG文件减少50%左右的存储空间", InputType: InputTypeSwitch, Sort: 18, Options: ""},
		{Category: ConfigCategoryConverter, Name: ConfigConverterEnableConvertRepeatedDocument, Label: "是否转换重复文档", Value: "false", Placeholder: "对于已转换过的文档，再次被上传时是否再转换一次", InputType: InputTypeSwitch, Sort: 20, Options: ""},

		// 下载配置
		{Category: ConfigCategoryDownload, Name: ConfigDownloadEnableGuestDownload, Label: "是否允许游客下载", Value: "false", Placeholder: "启用之后，未登录用户可以下载免费文档", InputType: InputTypeSwitch, Sort: 10, Options: ""},
		{Category: ConfigCategoryDownload, Name: ConfigDownloadFreeDownloadDuration, Label: "购买文档后多少天内允许免费重复下载", Value: "0", Placeholder: "0表示再次下载仍需购买，大于0表示指定多少天内有效", InputType: InputTypeNumber, Sort: 20, Options: ""},
		{Category: ConfigCategoryDownload, Name: ConfigDownloadUrlDuration, Label: "下载链接有效时长(秒)", Value: "60", Placeholder: "生成文档下载链接后多少秒之后链接失效", InputType: InputTypeNumber, Sort: 30, Options: ""},
		{Category: ConfigCategoryDownload, Name: ConfigDownloadTimesEveryDay, Label: "允许登录用户每天下载次数", Value: "10", Placeholder: "允许登录用户每天下载次数，0表示不允许下载", InputType: InputTypeNumber, Sort: 40, Options: ""},
		{Category: ConfigCategoryDownload, Name: ConfigDownloadTimesEveryIP, Label: "允许每个IP每天下载次数", Value: "10", Placeholder: "允许每个IP每天下载的次数，0表示不允许下载（针对所有用户）", InputType: InputTypeNumber, Sort: 41, Options: ""},
		{Category: ConfigCategoryDownload, Name: ConfigDownloadSecretKey, Label: "链接签名密钥", Value: "moredoc", Placeholder: "链接签名密钥，用于加密下载链接", InputType: InputTypeText, Sort: 50, Options: ""},

		// 积分规则配置
		{Category: ConfigCategoryScore, Name: ConfigScoreCreditName, Label: "积分名称", Value: "魔豆", Placeholder: "请输入网站的积分名称，默认为魔豆", InputType: InputTypeText, Sort: 1, Options: ""},
		{Category: ConfigCategoryScore, Name: ConfigScoreRegister, Label: "注册", Value: "10", Placeholder: "注册时获得的积分", InputType: InputTypeNumber, Sort: 10, Options: ""},
		{Category: ConfigCategoryScore, Name: ConfigScoreSignIn, Label: "签到", Value: "1", Placeholder: "每日签到获得的积分", InputType: InputTypeNumber, Sort: 20, Options: ""},
		{Category: ConfigCategoryScore, Name: ConfigScoreDeleteDocument, Label: "删除文档", Value: "1", Placeholder: "删除上传文档扣除的积分，0表示不扣除", InputType: InputTypeNumber, Sort: 25, Options: ""},
		{Category: ConfigCategoryScore, Name: ConfigScoreUploadDocument, Label: "上传文档", Value: "5", Placeholder: "上传一篇文档可获得的积分", InputType: InputTypeNumber, Sort: 30, Options: ""},
		{Category: ConfigCategoryScore, Name: ConfigScoreUploadDocumentLimit, Label: "每日上传文档奖励次数", Value: "1", Placeholder: "每天最多可以获得多少次文档上传奖励，0表示无奖励。", InputType: InputTypeNumber, Sort: 40, Options: ""},
		{Category: ConfigCategoryScore, Name: ConfigScoreDocumentCollected, Label: "文档被收藏", Value: "1", Placeholder: "上传的文档被收藏后获得的积分", InputType: InputTypeNumber, Sort: 50, Options: ""},
		{Category: ConfigCategoryScore, Name: ConfigScoreDocumentCollectedLimit, Label: "每日文档被收藏奖励次数", Value: "1", Placeholder: "每天最多可以获得多少次文档被收藏奖励，0表示无奖励。", InputType: InputTypeNumber, Sort: 60, Options: ""},
		{Category: ConfigCategoryScore, Name: ConfigScoreDocumentCommented, Label: "文档被评论", Value: "1", Placeholder: "上传的文档被评论后获得的积分", InputType: InputTypeNumber, Sort: 70, Options: ""},
		{Category: ConfigCategoryScore, Name: ConfigScoreDocumentCommentedLimit, Label: "每日文档被评论奖励次数", Value: "1", Placeholder: "每天最多可以获得多少次文档被评论奖励，0表示无奖励。", InputType: InputTypeNumber, Sort: 80, Options: ""},

		// 邮件配置
		{Category: ConfigCategoryEmail, Name: ConfigEmailEnable, Label: "是否启用邮件服务", Value: "false", Placeholder: "邮件服务，用于找回账户密码", InputType: InputTypeSwitch, Sort: 10, Options: ""},
		{Category: ConfigCategoryEmail, Name: ConfigEmailHost, Label: "SMTP 服务器地址", Value: "", Placeholder: "如：smtp.exmail.com", InputType: InputTypeText, Sort: 20, Options: ""},
		{Category: ConfigCategoryEmail, Name: ConfigEmailPort, Label: "SMTP 服务器端口", Value: "465", Placeholder: "如：465", InputType: InputTypeNumber, Sort: 30, Options: ""},
		{Category: ConfigCategoryEmail, Name: ConfigEmailIsTLS, Label: "是否启用TLS", Value: "true", Placeholder: "如果是TLS端口，请启用", InputType: InputTypeSwitch, Sort: 40, Options: ""},
		{Category: ConfigCategoryEmail, Name: ConfigEmailFromName, Label: "发件人名称", Value: "", Placeholder: "请输入您要展示的发件人名称", InputType: InputTypeText, Sort: 50, Options: ""},
		{Category: ConfigCategoryEmail, Name: ConfigEmailUsername, Label: "SMTP 账号", Value: "", Placeholder: "请输入您的邮箱账户", InputType: InputTypeText, Sort: 60, Options: ""},
		{Category: ConfigCategoryEmail, Name: ConfigEmailPassword, Label: "SMTP 密码", Value: "", Placeholder: "请输入您的邮箱密码", InputType: InputTypePassword, Sort: 70, Options: ""},
		{Category: ConfigCategoryEmail, Name: ConfigEmailDuration, Label: "邮件有效期", Value: "30", Placeholder: "找回密码时链接有效期，默认为30，表示30分钟", InputType: InputTypeNumber, Sort: 80, Options: ""},
		{Category: ConfigCategoryEmail, Name: ConfigEmailSecret, Label: "签名密钥", Value: "moredoc", Placeholder: "找回密码链接签名密钥", InputType: InputTypeText, Sort: 80, Options: ""},
		{Category: ConfigCategoryEmail, Name: ConfigEmailTestEmail, Label: "测试邮箱", Value: "", Placeholder: "用于每次变更配置时保存发送测试邮件", InputType: InputTypeText, Sort: 90, Options: ""},

		// 展示配置
		{Category: ConfigCategoryDisplay, Name: ConfigDisplayShowRegisterUserCount, Label: "是否显示注册用户数", Value: "true", Placeholder: "网站首页，是否显示注册用户数", InputType: InputTypeSwitch, Sort: 10, Options: ""},
		{Category: ConfigCategoryDisplay, Name: ConfigDisplayVirtualRegisterCount, Label: "网站虚拟注册用户数", Value: "0", Placeholder: "网站首页显示的用户数=真实注册用户数+虚拟注册用户数，用以避免网站初期用户注册数过少的尴尬窘境", InputType: InputTypeNumber, Sort: 20, Options: ""},
		{Category: ConfigCategoryDisplay, Name: ConfigDisplayShowIndexCategories, Label: "是否显示横栏分类", Value: "true", Placeholder: "网站首页中间横栏位置，是否显示分类", InputType: InputTypeSwitch, Sort: 30, Options: ""},
		{Category: ConfigCategoryDisplay, Name: ConfigDisplayPagesPerRead, Label: "文档【继续阅读】的页数", Value: "5", Placeholder: "用户阅读文档，每次点击继续阅读按钮时阅读的页数，默认为5，表示5页", InputType: InputTypeNumber, Sort: 40, Options: ""},
		{Category: ConfigCategoryDisplay, Name: ConfigDisplayMaxSearchPages, Label: "文档搜索结果最大页数", Value: "100", Placeholder: "搜索结果，默认最大展示100页，0表示不限制", InputType: InputTypeNumber, Sort: 50, Options: ""},
		{Category: ConfigCategoryDisplay, Name: ConfigDisplayCopyrightStatement, Label: "版权声明", Value: "本站文档数据由用户上传，仅供学习交流，如侵犯您的权益，请联系我们进行删除。", Placeholder: "网站最底部版权声明，支持HTML", InputType: InputTypeTextarea, Sort: 60, Options: ""},
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

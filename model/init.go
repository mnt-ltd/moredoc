package model

import (
	"database/sql"
	"errors"
	"moredoc/conf"
	"moredoc/util/captcha"
	"strings"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type TableColumn struct {
	Field      string `gorm:"Field"`
	Type       string `gorm:"Type"`
	Collation  string `gorm:"Collation"`
	Null       string `gorm:"Null"`
	Key        string `gorm:"Key"`
	Default    string `gorm:"Default"`
	Extra      string `gorm:"Extra"`
	Privileges string `gorm:"Privileges"`
	Comment    string `gorm:"Comment"`
}

var tablePrefix string

type DBModel struct {
	db             *gorm.DB
	tablePrefix    string
	logger         *zap.Logger
	tableFields    map[string][]string
	tableFieldsMap map[string]map[string]struct{}
}

func NewDBModel(cfg *conf.Database, lg *zap.Logger) (m *DBModel, err error) {
	if lg == nil {
		err = errors.New("logger cant be nil")
		return
	}

	tablePrefix = cfg.Prefix

	m = &DBModel{
		logger:         lg.Named("model"),
		tablePrefix:    cfg.Prefix,
		tableFields:    make(map[string][]string),
		tableFieldsMap: make(map[string]map[string]struct{}),
	}

	var (
		db    *gorm.DB
		sqlDB *sql.DB
	)

	sqlLogLevel := logger.Info
	if !cfg.ShowSQL {
		sqlLogLevel = logger.Silent
	}

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       cfg.DSN, // DSN data source name
		DefaultStringSize:         255,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.Prefix, // 表名前缀，`User`表为`t_users`
			SingularTable: true,       // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		Logger: logger.Default.LogMode(sqlLogLevel),
	})
	if err != nil {
		m.logger.Error("NewDBModel", zap.Error(err), zap.Any("config", cfg))
		return
	}

	sqlDB, err = db.DB()
	if err != nil {
		m.logger.Error("db.DB()", zap.Error(err))
		return
	}

	if cfg.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdle)
	}

	if cfg.MaxOpen > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxOpen)
	}

	m.db = db

	// 获取所有数据库表，并把数据库表字段加入到全局map，以便根据指定字段查询数据
	tables, err := m.ShowTables()
	if err != nil {
		m.logger.Error("ShowTables", zap.Error(err))
		return nil, err
	}

	for _, table := range tables {
		columns, err := m.showTableColumn(table)
		if err != nil {
			m.logger.Error("showTableColumn", zap.Error(err))
			return nil, err
		}

		var fields []string
		for _, col := range columns {
			fields = append(fields, col.Field)
		}

		m.tableFields[table] = fields
		filedsMap := make(map[string]struct{})
		for _, field := range fields {
			filedsMap[field] = struct{}{}
		}
		m.tableFieldsMap[table] = filedsMap
	}
	return
}

func (m *DBModel) SyncDB() (err error) {
	tableModels := []interface{}{
		&Attachment{},
		&Banner{},
		&Category{},
		&Config{},
		&Document{},
		&DocumentCategory{},
		&DocumentScore{},
		&Download{},
		&Friendlink{},
		&User{},
		&Group{},
		&UserGroup{},
		&Permission{},
		&GroupPermission{},
	}
	if err = m.db.AutoMigrate(tableModels...); err != nil {
		m.logger.Fatal("SyncDB", zap.Error(err))
	}

	if err = m.initDatabase(); err != nil {
		m.logger.Fatal("SyncDB", zap.Error(err))
	}
	return
}

func (m *DBModel) GetDB() *gorm.DB {
	return m.db
}

func (m *DBModel) ShowTables() (tables []string, err error) {
	err = m.db.Raw("show tables").Scan(&tables).Error
	if err != nil {
		m.logger.Error("ShowTables", zap.Error(err))
	}
	return
}

// FilterValidFields 过滤掉不存在的字段
func (m *DBModel) FilterValidFields(tableName string, fields ...string) (validFields []string) {
	fieldsMap, ok := m.tableFieldsMap[tableName]
	if ok {
		for _, field := range fields {
			field = strings.ToLower(strings.TrimSpace(field))
			if _, ok := fieldsMap[field]; ok {
				validFields = append(validFields, field)
			}
		}
	}
	return
}

func (m *DBModel) showTableColumn(tableName string) (columns []TableColumn, err error) {
	err = m.db.Raw("SHOW FULL COLUMNS FROM " + tableName).Find(&columns).Error
	if err != nil {
		m.logger.Error("ShowTableColumn", zap.Error(err))
	}
	return
}

// initialDatabase 初始化数据库相关数据
func (m *DBModel) initDatabase() (err error) {
	// 1. 初始化用户
	if err = m.initUser(); err != nil {
		m.logger.Error("initialDatabase", zap.Error(err))
		return
	}

	// 2. 初始化配置
	if err = m.initConfig(); err != nil {
		m.logger.Error("initialDatabase", zap.Error(err))
		return
	}
	return
}

func (m *DBModel) initUser() (err error) {
	// 如果不存在任意用户，则初始化一个用户作为管理员
	var existUser User
	m.db.Select("id").First(&existUser)
	if existUser.Id > 0 {
		return
	}

	// 初始化一个用户
	user := &User{Username: "admin", Password: "123456"}
	groupId := 1 // ID==1的用户组为管理员组
	err = m.CreateUser(user, int64(groupId))
	if err != nil {
		m.logger.Error("initUser", zap.Error(err))
	}
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

		// JWT 配置项
		{Category: ConfigCategoryJWT, Name: ConfigJWTDuration, Label: "Token有效期", Value: "365", Placeholder: "用户Token签名有效期，单位为天，默认365天", InputType: "number", Sort: 11, Options: ""},
		{Category: ConfigCategoryJWT, Name: ConfigJWTSecret, Label: "Token密钥", Value: uuid.Must(uuid.NewV4()).String(), Placeholder: "用户Token签名密钥，修改之后，之前所有的token签名都将失效，请慎重修改", InputType: "text", Sort: 12, Options: ""},

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

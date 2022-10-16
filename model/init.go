package model

import (
	"database/sql"
	"errors"
	"fmt"
	"moredoc/conf"
	"strings"
	"sync"

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
	validToken     sync.Map // map[tokenUUID]struct{} 有效的token uuid
	invalidToken   sync.Map // map[tokenUUID]struct{} 存在，未过期但无效token，比如读者退出登录后的token
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
		&Logout{},
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
	// if err = m.initPermission(); err != nil {
	// 	m.logger.Error("initialDatabase", zap.Error(err))
	// 	return
	// }

	// 初始化用户组及其权限
	if err = m.initGroupAndPermission(); err != nil {
		m.logger.Error("initialDatabase", zap.Error(err))
		return
	}

	// 初始化用户
	if err = m.initUser(); err != nil {
		m.logger.Error("initialDatabase", zap.Error(err))
		return
	}

	// 初始化配置
	if err = m.initConfig(); err != nil {
		m.logger.Error("initialDatabase", zap.Error(err))
		return
	}
	return
}

// 初始化用户组
func (m *DBModel) initGroupAndPermission() (err error) {
	groups := []Group{
		{Id: 1, Title: "超级管理员", IsDisplay: 1, Description: "系统超级管理员", UserCount: 0, Sort: 0},
		{Id: 2, Title: "普通用户", IsDisplay: 1, Description: "普通用户", UserCount: 0, Sort: 0, IsDefault: 1},
		{Id: 3, Title: "游客", IsDisplay: 1, Description: "游客", UserCount: 0, Sort: 0},
	}

	// 如果没有任何用户组，则初始化
	var existGroup Group
	m.db.First(&existGroup)
	if existGroup.Id > 0 {
		return
	}

	err = m.db.Create(&groups).Error
	if err != nil {
		m.logger.Error("initGroup", zap.Error(err))
	}

	return
}

// generateQueryLike 生成like查询。Like 查询比较特殊，统一用or来拼接查询的字段
func (m *DBModel) generateQueryLike(db *gorm.DB, tableName string, queryLike map[string][]interface{}) *gorm.DB {
	if len(queryLike) > 0 {
		var likeQuery []string
		var likeValues []interface{}
		for field, values := range queryLike {
			fields := m.FilterValidFields(tableName, field)
			if len(fields) == 0 {
				continue
			}
			for _, value := range values {
				valueStr := fmt.Sprintf("%v", value)
				likeQuery = append(likeQuery, fmt.Sprintf("%s like ?", field))
				likeValues = append(likeValues, "%"+valueStr+"%")
			}
		}
		if len(likeQuery) > 0 {
			db = db.Where(strings.Join(likeQuery, " or "), likeValues...)
		}
	}
	return db
}

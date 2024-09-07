package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moredoc/conf"
	"strings"
	"sync"
	"time"

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

type TableIndex struct {
	Table        string         `gorm:"column:Table"` // 表名
	NonUnique    int            `gorm:"column:Non_unique"`
	KeyName      string         `gorm:"column:Key_name"` // 索引名称
	SeqInIndex   int            `gorm:"column:Seq_in_index"`
	ColumnName   string         `gorm:"column:Column_name"` // 索引字段名称
	Collation    string         `gorm:"column:Collation"`   // 字符集
	Cardinality  int            `gorm:"column:Cardinality"`
	SubPart      sql.NullInt64  `gorm:"column:Sub_part"`
	Packed       sql.NullString `gorm:"column:Packed"`
	Null         string         `gorm:"column:Null"`
	IndexType    string         `gorm:"column:Index_type"`
	Comment      string         `gorm:"column:Comment"` // 索引备注
	IndexComment string         `gorm:"column:Index_comment"`
}

// 默认表前缀
var (
	tablePrefix            string = "mnt_"
	convertDocumentRunning        = false
)

type DBModel struct {
	db             *gorm.DB
	tablePrefix    string
	logger         *zap.Logger
	tableFields    map[string][]string
	tableFieldsMap map[string]map[string]struct{}
	validToken     sync.Map // map[tokenUUID]struct{} 有效的token uuid
	invalidToken   sync.Map // map[tokenUUID]struct{} 存在，未过期但无效token，比如读者退出登录后的token
	ctx            context.Context
	cfg            *conf.Database
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
		cfg:            cfg,
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

	sqlDB.SetConnMaxLifetime(time.Hour)
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
		&AttachmentContent{},
		&Banner{},
		&Category{},
		&Config{},
		&Document{},
		&DocumentCategory{},
		&DocumentError{},
		&DocumentScore{},
		&DocumentRelate{},
		&Download{},
		&Friendlink{},
		&User{},
		&Group{},
		&UserGroup{},
		&Permission{},
		&GroupPermission{},
		&Logout{},
		&Article{},
		&ArticleCategory{},
		&Favorite{},
		&Comment{},
		&Dynamic{},
		&Sign{},
		&Report{},
		&Navigation{},
		&Punishment{},
		&EmailCode{},
		&Advertisement{},
		&SearchRecord{},
		&Language{},
		&ArticleRelate{},
	}

	m.alterTableBeforeSyncDB()
	if err = m.db.AutoMigrate(tableModels...); err != nil {
		m.logger.Fatal("SyncDB", zap.Error(err))
		return
	}
	m.alterTableAfterSyncDB()

	if err = m.initDatabase(); err != nil {
		m.logger.Fatal("SyncDB", zap.Error(err))
	}
	return
}

func (m *DBModel) RunTasks() {
	go m.loopCovertDocument()
	go m.cronUpdateSitemap()
	go m.cronMarkAttachmentDeleted()
	go m.cronCleanInvalidAttachment()
	go m.createSearchRecordFromQueue()
	go m.cronCheckLatestVersion()
	go m.checkAndStartSSR()
}

func (m *DBModel) GetDB() *gorm.DB {
	return m.db
}

// 用来承载事务的上下文
type contextTxKey struct{}

// DB 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (m *DBModel) DB() *gorm.DB {
	if m.ctx != nil {
		tx, ok := m.ctx.Value(contextTxKey{}).(*gorm.DB)
		if ok {
			return tx
		}
	}
	return m.db
}

func (m *DBModel) ShowTables() (tables []string, err error) {
	err = m.db.Raw("show tables").Scan(&tables).Error
	if err != nil {
		m.logger.Error("ShowTables", zap.Error(err))
	}
	return
}

func (m *DBModel) alterTableBeforeSyncDB() {
	// 查询mnt_user表，将email字段由唯一索引删掉，以便变更为普通索引
	tableUser := User{}.TableName()
	indexes := m.ShowIndexes(tableUser)
	m.logger.Debug("alterTableBeforeSyncDB", zap.String("table", tableUser), zap.Any("indexes", indexes))
	if len(indexes) > 0 {
		for _, index := range indexes {
			if index.ColumnName == "email" && index.NonUnique == 0 { // 唯一索引，需要删除原索引
				err := m.db.Exec(fmt.Sprintf("alter table %s drop index %s", tableUser, index.KeyName)).Error
				if err != nil {
					m.logger.Error("alterTableBeforeSyncDB", zap.Error(err))
				}
			}
		}
	}

	// 查询category表，将原本有title和parent_id的唯一索引删除
	tableCategory := Category{}.TableName()
	m.db.Exec(fmt.Sprintf("alter table %s drop index %s", tableCategory, "parent_id_title"))

	// 删除favorite表相关索引
	tableFavorite := Favorite{}.TableName()
	m.db.Exec(fmt.Sprintf("alter table %s drop index %s", tableFavorite, "idx_user_document"))
	m.db.Exec(fmt.Sprintf("alter table %s drop index %s", tableFavorite, "idx_created_at"))
}

func (m *DBModel) alterTableAfterSyncDB() {

}

func (m *DBModel) ShowIndexes(table string) (indexes []TableIndex) {
	sql := "show index from " + table
	err := m.db.Raw(sql).Find(&indexes).Error
	if err != nil {
		m.logger.Error("ShowIndexes", zap.Error(err))
	}
	return
}

// FilterValidFields 过滤掉不存在的字段
func (m *DBModel) FilterValidFields(tableName string, fields ...string) (validFields []string) {
	alias := ""
	slice := strings.Split(tableName, " ")
	if len(slice) == 2 {
		alias = slice[1] + "."
		tableName = slice[0]
	}
	fieldsMap, ok := m.tableFieldsMap[tableName]
	if ok {
		for _, field := range fields {
			field = strings.ToLower(strings.TrimSpace(field))
			if _, ok := fieldsMap[field]; ok {
				validFields = append(validFields, fmt.Sprintf("%s%s", alias, field))
			}
		}
	}
	return
}

// GetTableFields 查询指定表的所有字段
func (m *DBModel) GetTableFields(tableName string, ignoreField ...string) (fields []string) {
	ignoreFieldMap := make(map[string]struct{})
	for _, field := range ignoreField {
		ignoreFieldMap[field] = struct{}{}
	}

	slice := strings.Split(tableName, " ")
	alias := ""
	if len(slice) == 2 {
		tableName = slice[0]
		alias = slice[1] + "."
	}
	fieldsMap, ok := m.tableFieldsMap[tableName]
	if ok {
		for field := range fieldsMap {
			f := fmt.Sprintf("%s%s", alias, field)
			_, ok1 := ignoreFieldMap[f]
			_, ok2 := ignoreFieldMap[field]
			if ok1 || ok2 {
				continue
			}
			fields = append(fields, f)
		}
	}
	return
}

// 关闭数据库连接
func (m *DBModel) CloseDB() {
	sqlDB, err := m.db.DB()
	if err != nil {
		m.logger.Error("db.DB()", zap.Error(err))
		return
	}
	sqlDB.Close()
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
	// 初始化用户组及其权限
	if err = m.initGroupAndPermission(); err != nil {
		m.logger.Error("initGroupAndPermission", zap.Error(err))
		return
	}

	// 初始化用户
	if err = m.initUser(); err != nil {
		m.logger.Error("initUser", zap.Error(err))
		return
	}

	// 初始化配置
	if err = m.initConfig(); err != nil {
		m.logger.Error("initConfig", zap.Error(err))
		return
	}

	// 初始化文章
	m.initArticle()

	// 初始化友情链接
	if err = m.initFriendlink(); err != nil {
		m.logger.Error("initFriendlink", zap.Error(err))
		return
	}

	// 更新uuid
	m.checkAndUpdateDocumentUUID()

	// 初始化语言
	if err = m.initLanguage(); err != nil {
		m.logger.Error("initLanguage", zap.Error(err))
	}

	// 初始化导航栏
	m.initNavigation()

	// 初始化静态页面SEO
	m.InitSEO()
	return
}

func (m *DBModel) initLanguage() (err error) {
	// 初始化语言
	var lang Language
	m.db.Find(&lang)
	if lang.Id > 0 {
		return
	}
	langs := getLangs()
	err = m.db.Create(&langs).Error
	if err != nil {
		m.logger.Error("initLanguage", zap.Error(err))
	}
	return
}

// 初始化用户组
func (m *DBModel) initGroupAndPermission() (err error) {
	groups := []Group{
		{Id: 1, Title: "超级管理员", IsDisplay: true, Description: "系统超级管理员", UserCount: 0, Sort: 0, EnableUpload: true, EnableDocumentReview: false, EnableComment: true, EnableCommentApproval: false},
		{Id: 2, Title: "普通用户", IsDisplay: true, Description: "普通用户", UserCount: 0, Sort: 0, IsDefault: true, EnableUpload: false, EnableDocumentReview: true, EnableComment: true, EnableCommentApproval: true},
	}

	// 如果没有任何用户组，则初始化
	var existGroup Group
	m.db.First(&existGroup)

	sess := m.db.Begin()
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()

	if existGroup.Id == 0 {
		// 用户组还不存在，则创建初始用户组
		err = sess.Create(&groups).Error
		if err != nil {
			m.logger.Error("initGroup", zap.Error(err))
			return
		}
	}

	// 初始化权限
	for _, permission := range getPermissions() {
		var existPermission Permission
		sess.Where("method = ? and path = ?", permission.Method, permission.Path).First(&existPermission)
		if existPermission.Id == 0 {
			err = sess.Create(&permission).Error
			if err != nil {
				m.logger.Error("initPermission", zap.Error(err))
				return
			}
			continue
		}

		if existPermission.Title == "" {
			existPermission.Title = permission.Title
			err = sess.Save(&existPermission).Error
			if err != nil {
				m.logger.Error("initPermission", zap.Error(err))
				return
			}
		}
	}

	// 废弃了的API权限，需要删除
	deprecatedPermissions := []string{
		"/api.v1.AttachmentAPI/DeleteAttachment",
	}
	for _, path := range deprecatedPermissions {
		err = sess.Where("path = ?", path).Delete(&Permission{}).Error
		if err != nil {
			m.logger.Error("initPermission", zap.Error(err))
			return
		}
	}
	return
}

func (m *DBModel) initFriendlink() (err error) {
	var friendlink Friendlink
	m.db.Find(&friendlink)
	if friendlink.Id > 0 {
		return
	}

	// 默认友链
	var friendlinks = []Friendlink{
		{Title: "摩枫网络科技", Link: "https://mnt.ltd", Enable: true},
		{Title: "书栈网", Link: "https://www.bookstack.cn", Enable: true},
	}

	err = m.db.Create(&friendlinks).Error
	if err != nil {
		m.logger.Error("initFriendlink", zap.Error(err))
	}
	return
}

// generateQueryLike 生成like查询。Like 查询比较特殊，统一用or来拼接查询的字段
func (m *DBModel) generateQueryLike(db *gorm.DB, tableName string, queryLike map[string][]interface{}) *gorm.DB {
	alias := ""
	slice := strings.Split(tableName, " ")
	if len(slice) == 2 {
		tableName = slice[0]
		alias = slice[1] + "."
	}

	if len(queryLike) > 0 {
		var likeQuery []string
		var likeValues []interface{}
		for field, values := range queryLike {
			fields := m.FilterValidFields(tableName, field)
			if len(fields) == 0 {
				continue
			}
			for _, value := range values {
				valueStr := strings.TrimSpace(fmt.Sprintf("%v", value))
				if !strings.Contains(valueStr, "%") {
					valueStr = "%" + valueStr + "%"
				}
				valueStr = strings.ReplaceAll(valueStr, " ", "%")
				likeQuery = append(likeQuery, fmt.Sprintf("%s%s like ?", alias, field))
				likeValues = append(likeValues, valueStr)
			}
		}
		if len(likeQuery) > 0 {
			db = db.Where(strings.Join(likeQuery, " or "), likeValues...)
		}
	}
	return db
}

func (m *DBModel) generateQueryRange(db *gorm.DB, tableName string, queryRange map[string][2]interface{}, withNULL ...bool) *gorm.DB {
	alias := ""
	slice := strings.Split(tableName, " ")
	if len(slice) == 2 {
		tableName = slice[0]
		alias = slice[1] + "."
	}

	for field, rangeValue := range queryRange {
		fields := m.FilterValidFields(tableName, field)
		if len(fields) == 0 {
			continue
		}

		nullStr := ""
		if len(withNULL) > 0 && withNULL[0] {
			nullStr = fmt.Sprintf(" or %s%s is null", alias, field)
		}

		if rangeValue[0] != nil {
			db = db.Where(fmt.Sprintf("%s%s >= ?"+nullStr, alias, field), rangeValue[0])
		}
		if rangeValue[1] != nil {
			db = db.Where(fmt.Sprintf("%s%s <= ?"+nullStr, alias, field), rangeValue[1])
		}
	}
	return db
}

func (m *DBModel) generateQueryIn(db *gorm.DB, tableName string, queryIn map[string][]interface{}) *gorm.DB {
	alias := ""
	slice := strings.Split(tableName, " ")
	if len(slice) == 2 {
		tableName = slice[0]
		alias = slice[1] + "."
	}
	for field, values := range queryIn {
		fields := m.FilterValidFields(tableName, field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(fmt.Sprintf("%s%s in (?)", alias, field), values)
	}
	return db
}

func (m *DBModel) generateQuerySort(db *gorm.DB, tableName string, querySort []string) *gorm.DB {
	var sorts []string
	alias := ""
	slice := strings.Split(tableName, " ")
	if len(slice) == 2 {
		tableName = slice[0]
		alias = slice[1] + "."
	}
	for _, sort := range querySort {
		slice := strings.Split(sort, " ")
		if len(m.FilterValidFields(tableName, slice[0])) == 0 {
			continue
		}

		if len(slice) == 2 {
			item := strings.ToLower(slice[1])
			if item == "asc" || item == "desc" {
				sorts = append(sorts, fmt.Sprintf("%s%s %s", alias, slice[0], item))
			}
		} else {
			sorts = append(sorts, fmt.Sprintf("%s%s desc", alias, slice[0]))
		}
	}
	if len(sorts) > 0 {
		db = db.Order(strings.Join(sorts, ","))
	} else {
		db = db.Order(fmt.Sprintf("%sid desc", alias))
	}
	return db
}

// 检查MySQL数据库是否支持group by 查询
func (m *DBModel) IsSupportGroupBy() (yes bool, sqlMode string) {
	var variables struct {
		VariableName string `gorm:"column:Variable_name"`
		Value        string `gorm:"column:Value"`
	}
	err := m.db.Raw("SHOW VARIABLES LIKE 'sql_mode'").Scan(&variables).Error
	if err != nil {
		m.logger.Error("CheckMySQLGroupBy", zap.Error(err))
		return
	}
	m.logger.Debug("CheckMySQLGroupBy", zap.Any("variables", variables))
	yes = !strings.Contains(variables.Value, "ONLY_FULL_GROUP_BY")
	return yes, variables.Value
}

// 设置数据库的sql_mode，去掉 ONLY_FULL_GROUP_BY，使得支持group by查询
func (m *DBModel) SetSQLMode() (err error) {
	err = m.db.Exec("set global sql_mode=(select replace(@@sql_mode,'ONLY_FULL_GROUP_BY',''))").Error
	if err != nil {
		m.logger.Error("SetSQLMode", zap.Error(err))
		return
	}

	err = m.resetDB()
	if err != nil {
		m.logger.Error("resetDB", zap.Error(err))
	}
	return
}

func (m *DBModel) resetDB() (err error) {
	var (
		db    *gorm.DB
		sqlDB *sql.DB
	)

	sqlLogLevel := logger.Info
	if !m.cfg.ShowSQL {
		sqlLogLevel = logger.Silent
	}

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.cfg.DSN, // DSN data source name
		DefaultStringSize:         255,       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,     // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.cfg.Prefix, // 表名前缀，`User`表为`t_users`
			SingularTable: true,         // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		Logger: logger.Default.LogMode(sqlLogLevel),
	})
	if err != nil {
		m.logger.Error("NewDBModel", zap.Error(err), zap.Any("config", m.cfg))
		return
	}

	sqlDB, err = db.DB()
	if err != nil {
		m.logger.Error("db.DB()", zap.Error(err))
		return
	}

	if m.cfg.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(m.cfg.MaxIdle)
	}

	if m.cfg.MaxOpen > 0 {
		sqlDB.SetMaxIdleConns(m.cfg.MaxOpen)
	}

	sqlDB.SetConnMaxLifetime(time.Hour)
	m.CloseDB() // 关闭旧的数据库连接
	m.db = db   // 更新新的数据库连接
	return
}

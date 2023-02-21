package service

import (
	"database/sql"
	"fmt"
	"moredoc/conf"
	"moredoc/model"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func SyncDB(cfg *conf.Config, logger *zap.Logger) {
	err := checkAndCreateDatabase(cfg.Database.DSN, logger)
	if err != nil {
		logger.Fatal("checkAndCreateDatabase", zap.Error(err))
		return
	}

	lg := logger.Named("syncdb")
	lg.Info("start syncdb")
	dbModel, err := model.NewDBModel(&cfg.Database, logger)
	if err != nil {
		lg.Fatal("NewDBModel", zap.Error(err))
		return
	}

	err = dbModel.SyncDB()
	if err != nil {
		lg.Fatal("SyncDB", zap.Error(err))
		return
	}
	lg.Info("syncdb success")
}

func checkAndCreateDatabase(dsn string, loggger *zap.Logger) (err error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		loggger.Error("ParseDSN", zap.Error(err))
		return
	}

	dbName := cfg.DBName
	if dbName == "" {
		loggger.Error("ParseDSN", zap.String("database", "数据库名称不能为空"))
		return
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		loggger.Error("sql.Open", zap.Error(err))
		return
	}

	createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci", cfg.DBName)
	_, err = db.Exec(createDB)
	if err != nil {
		loggger.Error("db.Exec", zap.Error(err))
	}
	return
}

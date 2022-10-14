package service

import (
	"moredoc/conf"
	"moredoc/model"

	"go.uber.org/zap"
)

func SyncDB(cfg *conf.Config, logger *zap.Logger) {
	lg := logger.Named("syncdb")
	lg.Info("start syncdb")
	dbModel, err := model.NewDBModel(cfg, logger)
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

package service

import (
	"moredoc/conf"
	"moredoc/model"

	"go.uber.org/zap"
)

func Reconvert(cfg *conf.Config, logger *zap.Logger, ext string, documentId int64) {
	db, err := model.NewDBModel(&cfg.Database, logger)
	if err != nil {
		logger.Fatal("NewDBModel", zap.Error(err))
		return
	}
	logger.Info("Reconvert", zap.Int64("documentId", documentId), zap.String("ext", ext))
	db.ReconvertDocoument(documentId, ext)
	logger.Info("Reconvert", zap.Int64("documentId", documentId), zap.String("ext", ext), zap.String("status", "done!"))
}

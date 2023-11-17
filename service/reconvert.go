package service

import (
	"fmt"
	"moredoc/conf"

	"go.uber.org/zap"
)

func Reconvert(cfg *conf.Config, logger *zap.Logger, ext string, documentId int64) {
	fmt.Println("reconvert called")
}

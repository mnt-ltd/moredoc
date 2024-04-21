package service

import (
	"moredoc/conf"
	"moredoc/model"
	"moredoc/util"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func FixCover(cfg *conf.Config, logger *zap.Logger) {
	lg := logger.Named("FixCover")
	lg.Info("start...")
	dbModel, err := model.NewDBModel(&cfg.Database, logger)
	if err != nil {
		lg.Fatal("NewDBModel", zap.Error(err))
		return
	}

	page := 1
	size := 100
	for {
		var attachments []model.Attachment
		err := dbModel.DB().Where("type = ?", model.AttachmentTypeDocument).Select("hash", "path").Group("hash").Offset((page - 1) * size).Limit(size).Order("id asc").Find(&attachments).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			lg.Error("查询附件失败", zap.Error(err))
			break
		}
		if len(attachments) == 0 {
			break
		}

		for _, attachment := range attachments {
			coverBig := strings.TrimLeft(strings.TrimSuffix(attachment.Path, filepath.Ext(attachment.Path)), "./") + "/cover.big.png"
			cover := strings.TrimLeft(strings.TrimSuffix(attachment.Path, filepath.Ext(attachment.Path)), "./") + "/cover.png"
			os.Remove(cover)
			util.CopyFile(coverBig, cover)
			util.CropImage(cover, model.DocumentCoverWidth, model.DocumentCoverHeight, true)
			lg.Info("fixed", zap.String("cover", cover))
		}
		page++
	}
	lg.Info("done!")
}

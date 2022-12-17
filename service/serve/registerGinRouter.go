package serve

import (
	"moredoc/biz"
	"moredoc/middleware/auth"
	"moredoc/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterGinRouter 注册gin路由
func RegisterGinRouter(app *gin.Engine, dbModel *model.DBModel, logger *zap.Logger, auth *auth.Auth) (err error) {
	attachmentAPIService := biz.NewAttachmentAPIService(dbModel, logger)

	app.GET("/helloworld", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})
	app.GET("/view/page/:hash/:page", attachmentAPIService.ViewDocumentPages)
	app.GET("/view/cover/:hash", attachmentAPIService.ViewDocumentCover)
	app.GET("/download/:jwt", attachmentAPIService.DownloadDocument)

	checkPermissionGroup := app.Group("/api/v1/upload")
	checkPermissionGroup.Use(auth.AuthGin())
	{
		checkPermissionGroup.POST("avatar", attachmentAPIService.UploadAvatar)
		checkPermissionGroup.POST("config", attachmentAPIService.UploadConfig)
		checkPermissionGroup.POST("banner", attachmentAPIService.UploadBanner)
		checkPermissionGroup.POST("document", attachmentAPIService.UploadDocument)
		checkPermissionGroup.POST("category", attachmentAPIService.UploadCategory)
		checkPermissionGroup.POST("article", attachmentAPIService.UploadArticle)
	}

	return
}

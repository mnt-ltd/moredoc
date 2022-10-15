package serve

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterGinRouter 注册gin路由
func RegisterGinRouter(app *gin.Engine) (err error) {
	app.GET("/helloworld", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})
	return
}

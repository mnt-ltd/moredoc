package service

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"moredoc/conf"
	"moredoc/middleware/auth"
	"moredoc/middleware/jsonpb"
	"moredoc/model"
	"moredoc/service/serve"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip" // grpc gzip
)

// Run start server
func Run(cfg *conf.Config, logger *zap.Logger) {

	size := 100 * 1024 * 1024 // 100MB
	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(size)),
		grpc.WithDefaultCallOptions(grpc.UseCompressor("gzip")),
	}

	auth := auth.NewAuth(&cfg.JWT)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			auth.AuthUnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&jsonpb.JSONPb{OrigName: true, EmitDefaults: true, EnumsAsInts: true},
		),
	)

	dbModel, err := model.NewDBModel(&cfg.Database, logger)
	if err != nil {
		logger.Fatal("NewDBModel", zap.Error(err))
		return
	}
	defer dbModel.CloseDB()

	// 每次启动时，都对dist中的title进行一次处理，以替换掉关键字 moredoc
	go dbModel.InitSEO()
	dbModel.RunTasks()

	if yes, sqlMode := dbModel.IsSupportGroupBy(); !yes {
		logger.Warn("IsSupportGroupBy", zap.String("告警提示", "您的数据库不支持MySQL的group by 查询"), zap.String("sql_mode", sqlMode))
		logger.Warn("自动设置全局 sql_mode，但仍建议修改数据库的 sql_mode 配置，去掉 ONLY_FULL_GROUP_BY")
		err = dbModel.SetSQLMode()
		if err != nil {
			logger.Error("设置sql_mode失败", zap.Error(err))
		} else {
			logger.Info("设置sql_mode成功")
		}
	}

	if cfg.Level != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.New()
	app.Use(
		gzip.Gzip(gzip.BestCompression, gzip.WithExcludedExtensions([]string{".svg", ".png", ".gif", ".jpeg", ".jpg", ".ico"})), // gzip
		gin.Recovery(), // recovery
		// cors.Default(), // allows all origins
		cors.New(cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"*"},
			AllowOrigins:     []string{"*"}, //  Referrer Policy: strict-origin-when-cross-origin
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		}),
	)

	endpoint := fmt.Sprintf("localhost:%v", cfg.Port)
	err = serve.RegisterGRPCService(dbModel, logger, endpoint, auth, grpcServer, gwmux, dialOpts...)
	if err != nil {
		logger.Fatal("registerAPIService", zap.Error(err))
		return
	}

	serve.RegisterGinRouter(app, dbModel, logger, auth)

	// 根目录访问静态文件，要放在 grpc 服务的前面
	// 可以在 dist 目录下创建一个 index.html 文件并添加内容，然后访问 http://ip:port
	app.Use(static.Serve("/uploads", static.LocalFile("./uploads", true)))
	app.Use(static.Serve("/sitemap", static.LocalFile("./sitemap", true)))
	app.Use(static.Serve("/", static.LocalFile("./dist", true)))
	app.NoRoute(wrapH(grpcHandlerFunc(grpcServer, gwmux))) // grpcServer and grpcGatewayServer

	addr := fmt.Sprintf(":%v", cfg.Port)
	logger.Info("server start", zap.Int("port", cfg.Port))
	err = app.Run(addr)
	if err != nil {
		logger.Fatal(err.Error())
	}
}

// See: https://github.com/philips/grpc-gateway-example/issues/22
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
		if r.ProtoMajor == 2 { // grpc 请求
			grpcServer.ServeHTTP(w, r)
		} else if !strings.HasPrefix(r.URL.Path, "/api/") {
			http.ServeFile(w, r, "./dist/index.html")
		} else {
			// 如 /api/v1/xxxx ，/api/v2/xxxx
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

// wrapH overwrite gin.WrapH
func wrapH(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK) // reset 404
		h.ServeHTTP(c.Writer, c.Request)
	}
}

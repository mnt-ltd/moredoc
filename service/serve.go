package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	v1 "moredoc/api/v1"
	"moredoc/biz"
	"moredoc/conf"
	"moredoc/middleware/jsonpb"
	"moredoc/model"

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

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			// jwtil.UnaryServerInterceptor(jwtil.TokenAuthInterceptor),
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

	// =========================================================================
	// 【start】 在这里，注册您的API服务模块
	// =========================================================================

	endpoint := fmt.Sprintf("localhost:%v", cfg.Port)
	healthAPIService := biz.NewHealthAPIService(dbModel, logger)

	// 注册grpc服务，可以理解为类似给Web服务注册路由
	v1.RegisterHealthAPIServer(grpcServer, healthAPIService)

	// 注册 restful api 服务
	err = v1.RegisterHealthAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Fatal("RegisterHealthAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// =========================================================================
	// 【end】 在这里，注册您的API服务模块
	// =========================================================================

	if cfg.Level != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()
	app.Use(
		gzip.Gzip(gzip.DefaultCompression), // gzip
		gin.Recovery(),                     // recovery
		cors.Default(),                     // allows all origins
	)

	// Web router
	mountWebRouter(app)

	// 根目录访问静态文件，要放在 grpc 服务的前面
	// 可以在 dist 目录下创建一个 index.html 文件并添加内容，然后访问 http://ip:port
	app.Use(static.Serve("/", static.LocalFile("./dist", true)))
	app.NoRoute(func(ctx *gin.Context) {
		http.ServeFile(ctx.Writer, ctx.Request, "./dist/index.html")
	})

	// grpcServer and grpcGatewayServer
	app.NoRoute(wrapH(grpcHandlerFunc(grpcServer, gwmux)))

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
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
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

// mountWebRouter mount web router
func mountWebRouter(app *gin.Engine) {
	app.GET("/helloworld", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})
}

package service

import (
	"fmt"
	"net/http"
	"strings"

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

	if cfg.Level != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.New()
	app.Use(
		gzip.Gzip(gzip.DefaultCompression), // gzip
		gin.Recovery(),                     // recovery
		cors.Default(),                     // allows all origins
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

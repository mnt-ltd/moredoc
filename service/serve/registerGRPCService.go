package serve

import (
	"context"
	v1 "moredoc/api/v1"
	"moredoc/biz"
	"moredoc/middleware/auth"
	"moredoc/model"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RegisterGRPCService 注册grpc服务
func RegisterGRPCService(dbModel *model.DBModel, logger *zap.Logger, endpoint string, authMiddleWare *auth.Auth, grpcServer *grpc.Server, gwmux *runtime.ServeMux, dialOpts ...grpc.DialOption) (err error) {
	// 用户API接口服务
	userAPIService := biz.NewUserAPIService(dbModel, logger, authMiddleWare)
	v1.RegisterUserAPIServer(grpcServer, userAPIService)
	err = v1.RegisterUserAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterUserAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 分组API接口服务
	groupAPIService := biz.NewGroupAPIService(dbModel, logger)
	v1.RegisterGroupAPIServer(grpcServer, groupAPIService)
	err = v1.RegisterGroupAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterGroupAPIHandlerFromEndpoint", zap.Error(err))
		return
	}
	return
}

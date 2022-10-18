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

	// 友链API接口服务
	friendlinkAPIService := biz.NewFriendlinkAPIService(dbModel, logger)
	v1.RegisterFriendlinkAPIServer(grpcServer, friendlinkAPIService)
	err = v1.RegisterFriendlinkAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterFriendlinkAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 附件API接口服务
	attachmentAPIService := biz.NewAttachmentAPIService(dbModel, logger)
	v1.RegisterAttachmentAPIServer(grpcServer, attachmentAPIService)
	err = v1.RegisterAttachmentAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterAttachmentAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 横幅API接口服务
	bannerAPIService := biz.NewBannerAPIService(dbModel, logger)
	v1.RegisterBannerAPIServer(grpcServer, bannerAPIService)
	err = v1.RegisterBannerAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterBannerAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 权限API接口服务
	permissionAPIService := biz.NewPermissionAPIService(dbModel, logger)
	v1.RegisterPermissionAPIServer(grpcServer, permissionAPIService)
	err = v1.RegisterPermissionAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterPermissionAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	return
}

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

	// Config API接口服务
	configAPIService := biz.NewConfigAPIService(dbModel, logger)
	v1.RegisterConfigAPIServer(grpcServer, configAPIService)
	err = v1.RegisterConfigAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterConfigAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 注册分类服务
	categoryAPIService := biz.NewCategoryAPIService(dbModel, logger)
	v1.RegisterCategoryAPIServer(grpcServer, categoryAPIService)
	err = v1.RegisterCategoryAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterCategoryAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 注册文档服务
	documentAPIService := biz.NewDocumentAPIService(dbModel, logger)
	v1.RegisterDocumentAPIServer(grpcServer, documentAPIService)
	err = v1.RegisterDocumentAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterDocumentAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 文档回收站服务
	v1.RegisterRecycleAPIServer(grpcServer, documentAPIService)
	err = v1.RegisterRecycleAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterRecycleAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 文章服务
	articleAPIService := biz.NewArticleAPIService(dbModel, logger)
	v1.RegisterArticleAPIServer(grpcServer, articleAPIService)
	err = v1.RegisterArticleAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterArticleAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 收藏服务
	favoriteAPIService := biz.NewFavoriteAPIService(dbModel, logger)
	v1.RegisterFavoriteAPIServer(grpcServer, favoriteAPIService)
	err = v1.RegisterFavoriteAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterFavoriteAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 评论服务
	commentAPIService := biz.NewCommentAPIService(dbModel, logger)
	v1.RegisterCommentAPIServer(grpcServer, commentAPIService)
	err = v1.RegisterCommentAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterCommentAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 举报服务
	reportAPIService := biz.NewReportAPIService(dbModel, logger)
	v1.RegisterReportAPIServer(grpcServer, reportAPIService)
	err = v1.RegisterReportAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterReportAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 导航服务
	navgationAPIService := biz.NewNavigationAPIService(dbModel, logger)
	v1.RegisterNavigationAPIServer(grpcServer, navgationAPIService)
	err = v1.RegisterNavigationAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterNavigationAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 惩罚服务
	punishmentAPIService := biz.NewPunishmentAPIService(dbModel, logger)
	v1.RegisterPunishmentAPIServer(grpcServer, punishmentAPIService)
	err = v1.RegisterPunishmentAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterPunishmentAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 广告服务
	advertisementAPIService := biz.NewAdvertisementAPIService(dbModel, logger)
	v1.RegisterAdvertisementAPIServer(grpcServer, advertisementAPIService)
	err = v1.RegisterAdvertisementAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterAdvertisementAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	// 搜索记录服务
	searchRecordAPIService := biz.NewSearchRecordAPIService(dbModel, logger)
	v1.RegisterSearchRecordAPIServer(grpcServer, searchRecordAPIService)
	err = v1.RegisterSearchRecordAPIHandlerFromEndpoint(context.Background(), gwmux, endpoint, dialOpts)
	if err != nil {
		logger.Error("RegisterSearchRecordAPIHandlerFromEndpoint", zap.Error(err))
		return
	}

	return
}

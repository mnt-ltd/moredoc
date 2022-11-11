package biz

import (
	"context"
	"fmt"

	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ConfigAPIService struct {
	pb.UnimplementedConfigAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewConfigAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *ConfigAPIService) {
	return &ConfigAPIService{dbModel: dbModel, logger: logger.Named("ConfigAPIService")}
}

func (s *ConfigAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

// UpdateConfig 更新配置
func (s *ConfigAPIService) UpdateConfig(ctx context.Context, req *pb.Configs) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	var cfgs []*model.Config
	err = util.CopyStruct(req.Config, &cfgs)
	if err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("req", req), zap.Any("cfgs", cfgs), zap.Error(err))
		fmt.Println(err.Error())
	}

	err = s.dbModel.UpdateConfigs(cfgs, "value")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// ListConfig 查询配置
func (s *ConfigAPIService) ListConfig(ctx context.Context, req *pb.ListConfigRequest) (*pb.Configs, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	opt := &model.OptionGetConfigList{
		QueryIn: map[string][]interface{}{
			"category": util.Slice2Interface(req.Category),
		},
	}

	configs, err := s.dbModel.GetConfigList(opt)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbConfigs []*pb.Config
	util.CopyStruct(&configs, &pbConfigs)

	return &pb.Configs{Config: pbConfigs}, nil
}

// GetSettings 获取公开配置
func (s *ConfigAPIService) GetSettings(ctx context.Context, req *emptypb.Empty) (*pb.Settings, error) {
	res := &pb.Settings{
		// Captcha:  &pb.ConfigCaptcha{},
		System:   &pb.ConfigSystem{},
		Footer:   &pb.ConfigFooter{},
		Security: &pb.ConfigSecurity{},
	}

	// captcha := s.dbModel.GetConfigOfCaptcha()
	// util.CopyStruct(&captcha, res.Captcha)

	system := s.dbModel.GetConfigOfSystem()
	if err := util.CopyStruct(&system, res.System); err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("system", system), zap.Any("res.System", res.System), zap.Error(err))
	}

	footer := s.dbModel.GetConfigOfFooter()
	if err := util.CopyStruct(&footer, res.Footer); err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("footer", footer), zap.Any("res.Footer", res.Footer), zap.Error(err))
	}

	security := s.dbModel.GetConfigOfSecurity()
	if err := util.CopyStruct(&security, res.Security); err != nil {
		s.logger.Error("util.CopyStruct", zap.Any("security", security), zap.Any("res.Security", res.Security), zap.Error(err))
	}

	return res, nil
}

package biz

import (
	"context"
	"fmt"
	"runtime"

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

	isEmail := false
	for idx, cfg := range cfgs {
		if cfg.Category == model.ConfigCategoryEmail && cfg.Name == model.ConfigEmailPassword && cfg.Value == "******" {
			// 6个星号，不修改密码
			cfgs[idx].Value = s.dbModel.GetConfigOfEmail(model.ConfigEmailPassword).Password
		}
		isEmail = isEmail || cfg.Category == model.ConfigCategoryEmail
	}

	err = s.dbModel.UpdateConfigs(cfgs, "value")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if isEmail {
		cfgEmail := s.dbModel.GetConfigOfEmail(model.ConfigEmailEnable, model.ConfigEmailTestEmail)
		if cfgEmail.Enable && cfgEmail.TestEmail != "" {
			err = s.dbModel.SendMail("测试邮件", cfgEmail.TestEmail, "这是一封测试邮件")
			if err != nil {
				return nil, status.Error(codes.Internal, "邮件发送失败:"+err.Error())
			}
		}
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

	for idx, cfg := range pbConfigs {
		if cfg.Category == model.ConfigCategoryEmail && cfg.Name == model.ConfigEmailPassword {
			pbConfigs[idx].Value = "******"
		}
	}

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

func (s *ConfigAPIService) GetStats(ctx context.Context, req *emptypb.Empty) (res *pb.Stats, err error) {
	res = &pb.Stats{
		UserCount:       0,
		DocumentCount:   0,
		CategoryCount:   0,
		ArticleCount:    0,
		CommentCount:    0,
		BannerCount:     0,
		FriendlinkCount: 0,
		Os:              runtime.GOOS,
		Version:         util.Version,
		Hash:            util.Hash,
		BuildAt:         util.BuildAt,
	}
	res.UserCount, _ = s.dbModel.CountUser()
	res.DocumentCount, _ = s.dbModel.CountDocument()
	_, errPermission := s.checkPermission(ctx)
	if errPermission == nil {
		res.CategoryCount, _ = s.dbModel.CountCategory()
		res.ArticleCount, _ = s.dbModel.CountArticle()
		res.CommentCount, _ = s.dbModel.CountComment()
		res.BannerCount, _ = s.dbModel.CountBanner()
		res.FriendlinkCount, _ = s.dbModel.CountFriendlink()
		res.ReportCount, _ = s.dbModel.CountReport()
	}

	return
}

// UpdateSitemap 更新站点地图
func (s *ConfigAPIService) UpdateSitemap(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.UpdateSitemap()
	if err != nil {
		s.logger.Error("UpdateSitemap", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

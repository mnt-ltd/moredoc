package biz

import (
	"context"

	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LanguageAPIService struct {
	pb.UnimplementedLanguageAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewLanguageAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *LanguageAPIService) {
	return &LanguageAPIService{dbModel: dbModel, logger: logger.Named("LanguageAPIService")}
}

func (s *LanguageAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

func (s *LanguageAPIService) UpdateLanguageStatus(ctx context.Context, req *pb.UpdateLanguageStatusRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	if len(req.Id) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "语言ID不能为空")
	}

	err = s.dbModel.UpdateLanguageStatus(req.Id, req.Enable)
	if err != nil {
		s.logger.Error("UpdateLanguageStatus", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *LanguageAPIService) UpdateLanguage(ctx context.Context, req *pb.Language) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "语言ID不能为空")
	}

	lang := &model.Language{}
	util.CopyStruct(req, lang)
	fields := []string{
		"language",
		"enable",
		"sort",
	}
	err = s.dbModel.UpdateLanguage(lang, fields...)
	if err != nil {
		s.logger.Error("UpdateLanguage", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *LanguageAPIService) ListLanguage(ctx context.Context, req *pb.ListLanguageRequest) (*pb.ListLanguageReply, error) {
	// 未登录用户，智能查看启用的语言
	userClaims, _ := s.checkPermission(ctx)
	opt := &model.OptionGetLanguageList{
		WithCount:    true,
		QueryIn:      map[string][]interface{}{},
		SelectFields: req.Field,
		Page:         int(req.Page),
		Size:         int(req.Size_),
	}

	if opt.Page <= 0 {
		opt.Page = 1
	}
	if opt.Size <= 0 {
		opt.Size = 10
	}

	if userClaims != nil && userClaims.HaveAccess {
		if len(req.Enable) > 0 {
			opt.QueryIn["enable"] = util.Slice2Interface(req.Enable)
		}
		if req.Wd != "" {
			opt.QueryLike = map[string][]interface{}{
				"language": {req.Wd},
				"code":     {req.Wd},
			}
		}
	} else {
		opt.QueryIn["enable"] = []interface{}{true}
	}

	langs, total, err := s.dbModel.GetLanguageList(opt)
	if err != nil {
		s.logger.Error("ListLanguage", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListLanguageReply{}
	util.CopyStruct(&langs, &res.Language)
	res.Total = total
	return res, nil
}

func (s *LanguageAPIService) CreateLanguage(ctx context.Context, req *pb.Language) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	lang := &model.Language{}
	util.CopyStruct(req, lang)
	err = s.dbModel.CreateLanguage(lang)
	if err != nil {
		s.logger.Error("CreateLanguage", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *LanguageAPIService) DeleteLanguage(ctx context.Context, req *pb.DeleteLanguageRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	if len(req.Id) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "语言ID不能为空")
	}

	err = s.dbModel.DeleteLanguage(req.Id)
	if err != nil {
		s.logger.Error("DeleteLanguage", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

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

func (s *LanguageAPIService) ListLanguage(ctx context.Context, req *pb.ListLanguageRequest) (*pb.ListLanguageReply, error) {
	var userId int64
	userClaims, _ := s.checkPermission(ctx)
	if userClaims != nil {
		userId = userClaims.UserId
	}
	opt := &model.OptionGetLanguageList{
		WithCount: false,
		Size:      10000,
		Page:      1,
		QueryIn:   map[string][]interface{}{},
	}
	if s.dbModel.IsAdmin(userId) {
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

	langs, _, err := s.dbModel.GetLanguageList(opt)
	if err != nil {
		s.logger.Error("ListLanguage", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListLanguageReply{}
	util.CopyStruct(&langs, &res.Language)
	return res, nil
}

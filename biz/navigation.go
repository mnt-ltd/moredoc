package biz

import (
	"context"
	"strings"

	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type NavigationAPIService struct {
	pb.UnimplementedNavigationAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewNavigationAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *NavigationAPIService) {
	return &NavigationAPIService{dbModel: dbModel, logger: logger.Named("NavigationAPIService")}
}

func (s *NavigationAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

func (s *NavigationAPIService) CreateNavigation(ctx context.Context, req *pb.Navigation) (*pb.Navigation, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	nav := &model.Navigation{}
	err = util.CopyStruct(req, nav)
	if err != nil {
		s.logger.Error("CopyStruct", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	s.logger.Debug("CreateNavigation", zap.Any("nav", nav), zap.Any("req", req))

	err = s.dbModel.CreateNavigation(nav)
	if err != nil {
		s.logger.Error("CreateNavigation", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.Navigation{}
	err = util.CopyStruct(nav, res)
	if err != nil {
		s.logger.Error("CopyStruct", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *NavigationAPIService) UpdateNavigation(ctx context.Context, req *pb.Navigation) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	nav := &model.Navigation{}
	err = util.CopyStruct(req, nav)
	if err != nil {
		s.logger.Error("CopyStruct", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	s.logger.Debug("UpdateNavigation", zap.Any("nav", nav), zap.Any("req", req))

	err = s.dbModel.UpdateNavigation(nav)
	if err != nil {
		s.logger.Error("UpdateNavigation", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *NavigationAPIService) DeleteNavigation(ctx context.Context, req *pb.DeleteNavigationRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteNavigation(req.Id)
	if err != nil {
		s.logger.Error("DeleteNavigation", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *NavigationAPIService) GetNavigation(ctx context.Context, req *pb.GetNavigationRequest) (*pb.Navigation, error) {
	nav, err := s.dbModel.GetNavigation(req.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error("GetNavigation", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.Navigation{}
	err = util.CopyStruct(nav, res)
	if err != nil {
		s.logger.Error("CopyStruct", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *NavigationAPIService) ListNavigation(ctx context.Context, req *pb.ListNavigationRequest) (*pb.ListNavigationReply, error) {

	opt := &model.OptionGetNavigationList{
		Page:         int(req.Page),
		Size:         int(req.Size_),
		WithCount:    true,
		SelectFields: req.Field,
	}

	if req.Order != "" {
		opt.Sort = strings.Split(req.Order, ",")
	}

	userClaims, _ := s.checkPermission(ctx)
	if userClaims != nil && userClaims.HaveAccess {
		if req.Wd != "" {
			opt.QueryLike = map[string][]interface{}{
				"title":       {req.Wd},
				"description": {req.Wd},
				"href":        {req.Wd},
			}
		}
	}

	s.logger.Debug("ListNavigation", zap.Any("opt", opt), zap.Any("req", req))

	navs, total, err := s.dbModel.GetNavigationList(opt)
	if err != nil {
		s.logger.Error("GetNavigationList", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListNavigationReply{
		Total: total,
	}

	err = util.CopyStruct(navs, &res.Navigation)
	if err != nil {
		s.logger.Error("CopyStruct", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return res, nil
}

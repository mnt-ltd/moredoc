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

type FriendlinkAPIService struct {
	pb.UnimplementedFriendlinkAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewFriendlinkAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *FriendlinkAPIService) {
	return &FriendlinkAPIService{dbModel: dbModel, logger: logger.Named("FriendlinkAPIService")}
}

// checkPermission 检查用户权限
func (s *FriendlinkAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

// CreateFriendlink 创建友情链接，需要鉴权
func (s *FriendlinkAPIService) CreateFriendlink(ctx context.Context, req *pb.Friendlink) (*pb.Friendlink, error) {
	s.logger.Debug("CreateFriendlink", zap.Any("req", req))
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	friendlink := &model.Friendlink{}
	util.CopyStruct(req, friendlink)
	err = s.dbModel.CreateFriendlink(friendlink)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbFriendlink := &pb.Friendlink{}
	util.CopyStruct(friendlink, pbFriendlink)
	return pbFriendlink, nil
}

// UpdateFriendlink 更新友情链接，需要鉴权
func (s *FriendlinkAPIService) UpdateFriendlink(ctx context.Context, req *pb.Friendlink) (*pb.Friendlink, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "参数不正确")
	}

	friendlink := &model.Friendlink{}
	util.CopyStruct(req, friendlink)
	err = s.dbModel.UpdateFriendlink(friendlink)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbFriendlink := &pb.Friendlink{}
	util.CopyStruct(friendlink, pbFriendlink)
	return pbFriendlink, nil
}

// DeleteFriendlink 删除友情链接，需要鉴权
func (s *FriendlinkAPIService) DeleteFriendlink(ctx context.Context, req *pb.DeleteFriendlinkRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteFriendlink(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// GetFriendlink 查询友情链接
func (s *FriendlinkAPIService) GetFriendlink(ctx context.Context, req *pb.GetFriendlinkRequest) (*pb.Friendlink, error) {
	var fields []string
	_, err := s.checkPermission(ctx)
	if err != nil {
		fields = s.dbModel.GetFriendlinkPublicFields() // 非管理员可查询的字段
	}

	friendlink, err := s.dbModel.GetFriendlink(req.Id, fields...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbFriendlink := &pb.Friendlink{}
	util.CopyStruct(friendlink, pbFriendlink)
	return pbFriendlink, nil
}

func (s *FriendlinkAPIService) ListFriendlink(ctx context.Context, req *pb.ListFriendlinkRequest) (*pb.ListFriendlinkReply, error) {
	opt := &model.OptionGetFriendlinkList{
		WithCount: true,
		Page:      int(req.Page),
		Size:      int(req.Size_),
	}

	_, err := s.checkPermission(ctx)
	if err == nil {
		// 管理员可使用like查询
		if req.Wd != "" {
			wd := "%" + req.Wd + "%"
			opt.QueryLike = map[string][]interface{}{
				"title":       {wd},
				"description": {wd},
			}
		}
		// 管理员可查询指定状态的友链
		if len(req.Enable) > 0 {
			opt.QueryIn = map[string][]interface{}{"enable": util.Slice2Interface(req.Enable)}
		}
	} else {
		// 非管理员可查询的字段
		opt.SelectFields = s.dbModel.GetFriendlinkPublicFields()
		opt.QueryIn = map[string][]interface{}{"enable": {true}}
	}

	friendlink, total, err := s.dbModel.GetFriendlinkList(opt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbFriendlink []*pb.Friendlink
	util.CopyStruct(friendlink, &pbFriendlink)

	return &pb.ListFriendlinkReply{Friendlink: pbFriendlink, Total: total}, nil
}

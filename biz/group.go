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
	"gorm.io/gorm"
)

type GroupAPIService struct {
	pb.UnimplementedGroupAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewGroupAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *GroupAPIService) {
	return &GroupAPIService{dbModel: dbModel, logger: logger.Named("GroupAPIService")}
}

func (s *GroupAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

// CreateGroup 创建用户组
// 0. 检查用户权限
// 1. 检查用户组是否存在
// 2. 创建用户组
func (s *GroupAPIService) CreateGroup(ctx context.Context, req *pb.Group) (*pb.Group, error) {
	s.logger.Debug("CreateGroup", zap.Any("req", req))

	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	existGroup, err := s.dbModel.GetGroupByTitle(req.Title)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if existGroup.Id > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "分组名称已存在")
	}

	group := &model.Group{
		Title:       req.Title,
		Color:       req.Color,
		IsDefault:   req.IsDefault,
		IsDisplay:   req.IsDisplay,
		Sort:        int(req.Sort),
		Description: req.Description,
	}
	err = s.dbModel.CreateGroup(group)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbGroup := &pb.Group{}
	util.CopyStruct(group, pbGroup)
	return pbGroup, nil
}
func (s *GroupAPIService) UpdateGroup(ctx context.Context, req *pb.Group) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	var group model.Group
	util.CopyStruct(req, &group)
	err = s.dbModel.UpdateGroup(&group)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *GroupAPIService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteGroup(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *GroupAPIService) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.Group, error) {
	s.logger.Debug("GetGroup", zap.Any("req", req))
	group, err := s.dbModel.GetGroup(req.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbGroup := &pb.Group{}
	util.CopyStruct(group, pbGroup)

	return pbGroup, nil
}

// ListGroup 列出用户组。所有用户都可以查询
func (s *GroupAPIService) ListGroup(ctx context.Context, req *pb.ListGroupRequest) (*pb.ListGroupReply, error) {
	s.logger.Debug("ListGroup", zap.Any("req", req))
	opt := &model.OptionGetGroupList{
		Page:         int(req.Page),
		Size:         int(req.Size_),
		SelectFields: req.Field,
		WithCount:    true,
	}

	if req.Wd != "" {
		_, err := s.checkPermission(ctx)
		if err == nil {
			opt.QueryLike = map[string][]interface{}{
				"title": {req.Wd},
			}
		}
	}

	groups, total, err := s.dbModel.GetGroupList(opt)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbGroups []*pb.Group
	util.CopyStruct(&groups, &pbGroups)
	return &pb.ListGroupReply{Group: pbGroups, Total: total}, nil
}

// GetGroupPermission 获取用户组权限
func (s *GroupAPIService) GetGroupPermission(ctx context.Context, req *pb.GetGroupPermissionRequest) (*pb.GroupPermissions, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	groupPermissions, _ := s.dbModel.GetGroupPermissinsByGroupId(req.Id)
	pbGroupPermissions := &pb.GroupPermissions{}
	for _, item := range groupPermissions {
		pbGroupPermissions.PermissionId = append(pbGroupPermissions.PermissionId, item.PermissionId)
	}
	return pbGroupPermissions, nil
}

// UpdateGroupPermission 更新用户组权限
func (s *GroupAPIService) UpdateGroupPermission(ctx context.Context, req *pb.UpdateGroupPermissionRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.UpdateGroupPermissions(req.GroupId, req.PermissionId)
	if err != nil {
		s.logger.Error("UpdateGroupPermissions", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

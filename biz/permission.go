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

type PermissionAPIService struct {
	pb.UnimplementedPermissionAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewPermissionAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *PermissionAPIService) {
	return &PermissionAPIService{dbModel: dbModel, logger: logger.Named("PermissionAPIService")}
}

func (s *PermissionAPIService) checkPermission(ctx context.Context) (*auth.UserClaims, error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

// UpdatePermission 编辑权限信息。这里只能更新名称和描述，不能更新权限项。
func (s *PermissionAPIService) UpdatePermission(ctx context.Context, req *pb.Permission) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.UpdatePermission(&model.Permission{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
	}, "title", "description")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// GetPermission 查询单个权限信息
func (s *PermissionAPIService) GetPermission(ctx context.Context, req *pb.GetPermissionRequest) (*pb.Permission, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	permission, err := s.dbModel.GetPermission(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbPermission := &pb.Permission{}
	util.CopyStruct(&permission, pbPermission)

	return pbPermission, nil
}

// ListPermission 查询权限列表
func (s *PermissionAPIService) ListPermission(ctx context.Context, req *pb.ListPermissionRequest) (*pb.ListPermissionReply, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	opt := &model.OptionGetPermissionList{
		Page:      int(req.Page),
		Size:      int(req.Size_),
		WithCount: true,
		QueryLike: make(map[string][]interface{}),
		QueryIn:   make(map[string][]interface{}),
	}

	if req.Wd != "" {
		opt.QueryLike["title"] = []interface{}{req.Wd}
		opt.QueryLike["description"] = []interface{}{req.Wd}
	}

	if req.Path != "" {
		opt.QueryLike["path"] = []interface{}{req.Path}
	}

	if len(req.Method) > 0 {
		opt.QueryIn["method"] = util.Slice2Interface(req.Method)
	}

	permissions, total, err := s.dbModel.GetPermissionList(opt)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbPermissions []*pb.Permission
	util.CopyStruct(&permissions, &pbPermissions)

	return &pb.ListPermissionReply{
		Total:      total,
		Permission: pbPermissions,
	}, nil
}

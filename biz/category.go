package biz

import (
	"context"

	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryAPIService struct {
	pb.UnimplementedCategoryAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewCategoryAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *CategoryAPIService) {
	return &CategoryAPIService{dbModel: dbModel, logger: logger.Named("CategoryAPIService")}
}

func (s *CategoryAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

func (s *CategoryAPIService) CreateCategory(ctx context.Context, req *pb.Category) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *CategoryAPIService) UpdateCategory(ctx context.Context, req *pb.Category) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *CategoryAPIService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *CategoryAPIService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	return &pb.Category{}, nil
}
func (s *CategoryAPIService) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryReply, error) {
	return &pb.ListCategoryReply{}, nil
}

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

type FavoriteAPIService struct {
	pb.UnimplementedFavoriteAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewFavoriteAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *FavoriteAPIService) {
	return &FavoriteAPIService{dbModel: dbModel, logger: logger.Named("FavoriteAPIService")}
}

func (s *FavoriteAPIService) checkLogin(ctx context.Context) (*auth.UserClaims, error) {
	return checkGRPCLogin(s.dbModel, ctx)
}

func (s *FavoriteAPIService) CreateFavorite(ctx context.Context, req *pb.Favorite) (*pb.Favorite, error) {
	userClaims, err := s.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	favorite := &model.Favorite{
		UserId:     userClaims.UserId,
		DocumentId: req.DocumentId,
	}

	exsit, _ := s.dbModel.GetUserFavorite(favorite.UserId, favorite.DocumentId)
	if exsit.Id > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "您已经收藏过了")
	}

	err = s.dbModel.CreateFavorite(favorite)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "收藏失败:"+err.Error())
	}

	pbFavorite := &pb.Favorite{}
	util.CopyStruct(favorite, pbFavorite)

	return pbFavorite, nil
}

func (s *FavoriteAPIService) DeleteFavorite(ctx context.Context, req *pb.DeleteFavoriteRequest) (*emptypb.Empty, error) {
	userClaims, err := s.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteFavorite(userClaims.UserId, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除失败:"+err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *FavoriteAPIService) GetFavorite(ctx context.Context, req *pb.GetFavoriteRequest) (*pb.Favorite, error) {
	userClaims, err := s.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	favorite, _ := s.dbModel.GetUserFavorite(userClaims.UserId, req.DocumentId)

	pbFavorite := &pb.Favorite{}
	if favorite.Id > 0 {
		util.CopyStruct(&favorite, pbFavorite)
	}

	return pbFavorite, nil
}

// 获取用户自身的收藏列表
func (s *FavoriteAPIService) ListFavorite(ctx context.Context, req *pb.ListFavoriteRequest) (*pb.ListFavoriteReply, error) {
	userClaims, err := s.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	favorites, total, err := s.dbModel.GetFavoriteList(&model.OptionGetFavoriteList{
		Page:      int(req.Page),
		Size:      int(req.Size_),
		WithCount: true,
		QueryIn: map[string][]interface{}{
			"user_id": {userClaims.UserId},
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取失败:"+err.Error())
	}
	resp := &pb.ListFavoriteReply{
		Total:    total,
		Favorite: favorites,
	}
	return resp, nil
}

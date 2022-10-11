package biz

import (
	"context"

	pb "moredoc/api/v1"
	"moredoc/model"
	"moredoc/util/validate"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserAPIService struct {
	pb.UnimplementedUserAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewUserAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *UserAPIService) {
	return &UserAPIService{dbModel: dbModel, logger: logger.Named("UserAPIService")}
}

func (s *UserAPIService) getValidFieldMap() map[string]string {
	return map[string]string{"Username": "用户名", "Password": "密码"}
}

// Register 用户注册
// TODO: 1. 判断系统是否启用了注册
// TODO: 2. 如果系统启用了注册，判断是否需要管理员审核
// TODO: 3. 如果启用了验证码功能，则需要判断验证码是否正确
func (s *UserAPIService) Register(ctx context.Context, req *pb.RegisterAndLoginRequest) (*emptypb.Empty, error) {
	err := validate.ValidateStruct(req, s.getValidFieldMap())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	exist, _ := s.dbModel.GetUserByUsername(req.Username, "id")
	if exist.Id > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "用户名已存在")
	}

	user := &model.User{Username: req.Username, Password: req.Password}
	if err = s.dbModel.CreateUser(user); err != nil {
		s.logger.Error("CreateUser", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *UserAPIService) Login(ctx context.Context, req *pb.RegisterAndLoginRequest) (*pb.LoginReply, error) {
	errValidate := validate.ValidateStruct(req, s.getValidFieldMap())
	if errValidate != nil {
		return nil, status.Errorf(codes.InvalidArgument, errValidate.Error())
	}

	return &pb.LoginReply{}, nil
}

func (s *UserAPIService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return &pb.User{}, nil
}

func (s *UserAPIService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	return &pb.User{}, nil
}

func (s *UserAPIService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *UserAPIService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

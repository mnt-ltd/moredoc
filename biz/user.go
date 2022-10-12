package biz

import (
	"context"

	pb "moredoc/api/v1"
	"moredoc/model"
	"moredoc/util/validate"

	"moredoc/util/captcha"

	"github.com/alexandrevicenzi/unchained"
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
// TODO: 1. 如果系统启用了注册，判断是否需要管理员审核
func (s *UserAPIService) Register(ctx context.Context, req *pb.RegisterAndLoginRequest) (*emptypb.Empty, error) {
	err := validate.ValidateStruct(req, s.getValidFieldMap())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cfg := s.dbModel.GetConfigOfSecurity(
		model.ConfigSecurityEnableCaptchaRegister,
		model.ConfigSecurityEnableRegister,
		model.ConfigSecurityIsClose,
	)

	if !cfg.EnableRegister {
		return nil, status.Errorf(codes.InvalidArgument, "系统未开放注册")
	}

	if !cfg.IsClose {
		return nil, status.Errorf(codes.InvalidArgument, "网站已关闭，占时不允许注册")
	}

	if cfg.EnableCaptchaRegister && !captcha.VerifyCaptcha(req.CaptchaId, req.Captcha) {
		return nil, status.Errorf(codes.InvalidArgument, "验证码错误")
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

// Login 用户登录
// TODO: 1. 判断是否启用了验证码，如果启用了验证码，则需要进行验证码验证
func (s *UserAPIService) Login(ctx context.Context, req *pb.RegisterAndLoginRequest) (*pb.LoginReply, error) {
	errValidate := validate.ValidateStruct(req, s.getValidFieldMap())
	if errValidate != nil {
		return nil, status.Errorf(codes.InvalidArgument, errValidate.Error())
	}

	// 如果启用了验证码，则需要进行验证码验证
	cfg := s.dbModel.GetConfigOfSecurity(model.ConfigSecurityEnableCaptchaLogin)
	if cfg.EnableCaptchaLogin && !captcha.VerifyCaptcha(req.CaptchaId, req.Captcha) {
		return nil, status.Errorf(codes.InvalidArgument, "验证码错误")
	}

	user, err := s.dbModel.GetUserByUsername(req.Username, "id", "password")
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if ok, err := unchained.CheckPassword(req.Password, user.Password); !ok || err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "用户名或密码错误")
	}

	token, err := s.dbModel.CreateUserJWTToken(user.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.LoginReply{Token: token}, nil
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

//  GetUserCaptcha 获取用户验证码
func (s *UserAPIService) GetUserCaptcha(ctx context.Context, req *pb.GetUserCaptchaRequest) (res *pb.GetUserCaptchaReply, err error) {
	cfgCaptcha := s.dbModel.GetConfigOfCaptcha()
	cfgSecurity := s.dbModel.GetConfigOfSecurity()
	res = &pb.GetUserCaptchaReply{
		Enable: false,
		Type:   cfgCaptcha.Type,
	}
	switch req.Type {
	case "register":
		res.Enable = cfgSecurity.EnableCaptchaRegister
	case "login":
		res.Enable = cfgSecurity.EnableCaptchaLogin
	case "upload":
		res.Enable = cfgSecurity.EnableCaptchaUpload
	case "find_password":
		res.Enable = cfgSecurity.EnableCaptchaFindPassword
	case "comment":
		res.Enable = cfgSecurity.EnableCaptchaComment
	default:
		return nil, status.Errorf(codes.InvalidArgument, "不支持的验证码类型")
	}

	if res.Enable {
		res.Id, res.Captcha, err = captcha.GenerateCaptcha(cfgCaptcha.Type)
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return res, nil
}

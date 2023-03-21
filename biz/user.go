package biz

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	pb "moredoc/api/v1"
	v1 "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"
	"moredoc/util/validate"

	"moredoc/util/captcha"

	"github.com/alexandrevicenzi/unchained"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

const (
	ErrorMessageUsernameOrPasswordError = "用户名或密码不正确"
	ErrorMessageInvalidToken            = "您未登录或您的登录已过期，请重新登录"
	ErrorMessageUserNotExists           = "用户不存在"
	ErrorMessageInvalidOldPassword      = "原密码不正确"
	ErrorMessageUnsupportedCaptchaType  = "不支持的验证码类型"
)

type UserAPIService struct {
	pb.UnimplementedUserAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
	auth    *auth.Auth
}

func NewUserAPIService(dbModel *model.DBModel, logger *zap.Logger, auth *auth.Auth) (service *UserAPIService) {
	return &UserAPIService{dbModel: dbModel, logger: logger.Named("UserAPIService"), auth: auth}
}

func (s *UserAPIService) getValidFieldMap() map[string]string {
	return map[string]string{"Username": "用户名", "Password": "密码"}
}

func (s *UserAPIService) checkPermission(ctx context.Context) (*auth.UserClaims, error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

// Register 用户注册
func (s *UserAPIService) Register(ctx context.Context, req *pb.RegisterAndLoginRequest) (*pb.LoginReply, error) {
	err := validate.ValidateStruct(req, s.getValidFieldMap())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if yes := util.IsValidEmail(req.Email); !yes {
		return nil, status.Errorf(codes.InvalidArgument, "邮箱格式不正确")
	}

	cfg := s.dbModel.GetConfigOfSecurity(
		model.ConfigSecurityEnableCaptchaRegister,
		model.ConfigSecurityEnableRegister,
		model.ConfigSecurityIsClose,
	)

	if !cfg.EnableRegister {
		return nil, status.Errorf(codes.InvalidArgument, "系统未开放注册")
	}

	if cfg.IsClose {
		return nil, status.Errorf(codes.InvalidArgument, "网站已关闭，占时不允许注册")
	}

	if cfg.EnableCaptchaRegister && !captcha.VerifyCaptcha(req.CaptchaId, req.Captcha) {
		return nil, status.Errorf(codes.InvalidArgument, "验证码错误")
	}

	exist, _ := s.dbModel.GetUserByUsername(req.Username, "id")
	if exist.Id > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "用户名已存在")
	}
	exist, _ = s.dbModel.GetUserByEmail(req.Email, "id")
	if exist.Id > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "邮箱已存在")
	}

	user := &model.User{Username: req.Username, Password: req.Password, Email: req.Email}
	if ips, _ := util.GetGRPCRemoteIP(ctx); len(ips) > 0 {
		user.RegisterIp = ips[0]
		user.LastLoginIp = user.RegisterIp
	}
	now := time.Now()
	user.LoginAt = &now

	group, _ := s.dbModel.GetDefaultUserGroup()
	if group.Id <= 0 {
		return nil, status.Errorf(codes.Internal, "请联系管理员设置系统默认用户组")
	}

	// 用户积分
	cfgScore := s.dbModel.GetConfigOfScore(model.ConfigScoreRegister, model.ConfigScoreCreditName)
	user.CreditCount = int(cfgScore.Register)
	if err = s.dbModel.CreateUser(user, group.Id); err != nil {
		s.logger.Error("CreateUser", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if cfgScore.Register > 0 {
		// 积分记录
		s.dbModel.CreateDynamic(&model.Dynamic{
			UserId:  user.Id,
			Type:    model.DynamicTypeRegister,
			Content: fmt.Sprintf("成功注册成网站会员，获得 %d %s奖励", cfgScore.Register, cfgScore.CreditName),
		})
	}

	token, err := s.auth.CreateJWTToken(user.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbUser := &pb.User{}
	util.CopyStruct(&user, pbUser)

	return &pb.LoginReply{Token: token, User: pbUser}, nil
}

// Login 用户登录
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

	user, err := s.dbModel.GetUserByUsername(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if ok, err := unchained.CheckPassword(req.Password, user.Password); !ok || err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "用户名或密码错误")
	}

	token, err := s.auth.CreateJWTToken(user.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbUser := &pb.User{}
	util.CopyStruct(&user, pbUser)

	ip := ""
	if ips, _ := util.GetGRPCRemoteIP(ctx); len(ips) > 0 {
		ip = ips[0]
	}
	loginAt := time.Now()
	if e := s.dbModel.UpdateUser(&model.User{Id: user.Id, LoginAt: &loginAt, LastLoginIp: ip}, "login_at", "last_login_ip"); e != nil {
		s.logger.Error("UpdateUser", zap.Error(e))
	}

	return &pb.LoginReply{Token: token, User: pbUser}, nil
}

func (s *UserAPIService) Logout(ctx context.Context, req *emptypb.Empty) (res *emptypb.Empty, err error) {
	res = &emptypb.Empty{}
	userClaims, ok := ctx.Value(auth.CtxKeyUserClaims).(*auth.UserClaims)
	if !ok {
		return
	}

	// 标记退出的用户token
	s.dbModel.Logout(userClaims.UserId, userClaims.UUID, userClaims.ExpiresAt)
	return
}

// GetUser 根据ID获取用户信息
// 对于非管理员，只能获取公开字段
func (s *UserAPIService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	userId := req.Id
	fields := s.dbModel.GetUserPublicFields()
	userClaims, err := s.checkPermission(ctx)
	if err == nil || (userClaims != nil && (userClaims.UserId == userId || userId == 0)) {
		// 有权限或者查的是用户自己的资料
		fields = []string{}
	}

	if userId <= 0 {
		if userClaims == nil {
			return nil, status.Errorf(codes.InvalidArgument, "ID错误")
		}
		userId = userClaims.UserId
	}

	user, err := s.dbModel.GetUser(userId, fields...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	pbUser := &pb.User{}
	util.CopyStruct(&user, pbUser)
	return pbUser, nil
}

func (s *UserAPIService) SetUser(ctx context.Context, req *pb.SetUserRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	if len(req.GroupId) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "用户组不能为空")
	}

	err = s.dbModel.SetUserGroupAndPassword(req.Id, req.GroupId, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// UpdateUserProfile 更改用户信息
// 1. 用户更改自身信息
// 2. 管理员更改用户信息
func (s *UserAPIService) UpdateUserProfile(ctx context.Context, req *pb.User) (*emptypb.Empty, error) {
	userClaims, ok := ctx.Value(auth.CtxKeyUserClaims).(*auth.UserClaims)
	if !ok || s.dbModel.IsInvalidToken(userClaims.UUID) {
		return nil, status.Errorf(codes.Unauthenticated, ErrorMessageInvalidToken)
	}

	// 允许更改的字段
	fields := []string{"mobile", "email", "address", "signature", "avatar", "realname", "identity"}
	user := &model.User{
		Mobile: req.Mobile, Email: req.Email, Address: req.Address,
		Signature: req.Signature, Avatar: req.Avatar,
		Realname: req.Realname, Identity: req.Identity,
		Status: int8(req.Status), Id: req.Id,
	}

	// 更改用户自己的资料
	if req.Id <= 0 || req.Id == userClaims.UserId {
		user.Id = userClaims.UserId
		exist, _ := s.dbModel.GetUser(user.Id, "status")
		if exist.Status != model.UserStatusNormal {
			// 非正常的用户状态，禁止修改个人信息，以避免用户修改成非法信息等
			return nil, status.Errorf(codes.InvalidArgument, "您的用户状态异常，禁止修改个人信息")
		}

		if err := s.dbModel.UpdateUser(user, fields...); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		return &emptypb.Empty{}, nil
	}

	// 管理员更改用户资料，验证是否有权限
	fullMethod, _ := ctx.Value(auth.CtxKeyFullMethod).(string)
	if permission, yes := s.dbModel.CheckPermissionByUserId(userClaims.UserId, fullMethod); !yes {
		item := permission.Title
		if item == "" {
			item = permission.Path
		}
		return nil, status.Errorf(codes.PermissionDenied, errorMessagePermissionDeniedFormat, item)
	}

	// 对于管理员，允许该更用户状态
	fields = append(fields, "status")
	err := s.dbModel.UpdateUser(user, fields...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// UpdateUserPassword 更改用户密码
// 1. 用户更改自身密码：需要验证旧密码
// 2. 管理员更改用户密码：不需要验证旧密码
func (s *UserAPIService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*emptypb.Empty, error) {
	userClaims, ok := ctx.Value(auth.CtxKeyUserClaims).(*auth.UserClaims)
	if !ok || s.dbModel.IsInvalidToken(userClaims.UUID) {
		return nil, status.Errorf(codes.Unauthenticated, ErrorMessageInvalidToken)
	}

	err := validate.ValidateStruct(req, map[string]string{"NewPassword": "新密码"})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	// 更改自己的密码
	if req.Id <= 0 || req.Id == userClaims.UserId {
		existUser, _ := s.dbModel.GetUser(userClaims.UserId, "id", "password")
		if existUser.Id == 0 {
			return nil, status.Errorf(codes.Unauthenticated, ErrorMessageUserNotExists)
		}

		if ok, err := unchained.CheckPassword(req.OldPassword, existUser.Password); !ok || err != nil {
			return nil, status.Errorf(codes.InvalidArgument, ErrorMessageInvalidOldPassword)
		}

		err = s.dbModel.UpdateUserPassword(userClaims.UserId, req.NewPassword)
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		return &emptypb.Empty{}, nil
	}

	// 管理员更改用户密码
	fullMethod, _ := ctx.Value(auth.CtxKeyFullMethod).(string)
	if permission, yes := s.dbModel.CheckPermissionByUserId(userClaims.UserId, fullMethod); !yes {
		item := permission.Title
		if item == "" {
			item = permission.Path
		}
		return nil, status.Errorf(codes.PermissionDenied, errorMessagePermissionDeniedFormat, item)
	}

	err = s.dbModel.UpdateUserPassword(req.Id, req.NewPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *UserAPIService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteUser(req.Id)
	if err != nil {
		s.logger.Error("DeleteUser", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// ListUser 查询用户列表
// 1. 非管理员，只能查询公开信息
// 2. 管理员，可以查询全部信息
func (s *UserAPIService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	s.logger.Debug("ListUser", zap.Any("req", req), zap.Any("status", req.Status))

	var (
		userId        int64
		limitFileds   = model.UserPublicFields
		fullMethod, _ = ctx.Value(auth.CtxKeyFullMethod).(string)
	)

	userClaims, ok := ctx.Value(auth.CtxKeyUserClaims).(*auth.UserClaims)
	if ok && !s.dbModel.IsInvalidToken(userClaims.UUID) {
		userId = userClaims.UserId
	}

	opt := &model.OptionGetUserList{
		Page:      int(req.Page),
		Size:      int(req.Size_),
		WithCount: true,
	}
	if req.Limit > 0 {
		opt.Page = 1
		opt.Size = int(req.Limit)
		opt.WithCount = false
	}

	if len(req.Id) > 0 {
		var ids []interface{}
		for _, id := range req.Id {
			ids = append(ids, id)
		}
		opt.WithCount = false
		opt.Ids = ids
	}

	if len(req.GroupId) > 0 {
		opt.QueryIn = map[string][]interface{}{"group_id": util.Slice2Interface(req.GroupId)}
	}

	if len(req.Status) > 0 {
		opt.QueryIn = map[string][]interface{}{"status": util.Slice2Interface(req.Status)}
	}

	if req.Sort != "" {
		opt.Sort = strings.Split(req.Sort, ",")
	}

	if _, yes := s.dbModel.CheckPermissionByUserId(userId, fullMethod); yes {
		limitFileds = []string{} // 管理员，可以查询全部信息
		if req.Wd != "" {
			value := []interface{}{"%" + strings.TrimSpace(req.Wd) + "%"}
			opt.QueryLike = map[string][]interface{}{
				"username": value, "realname": value, "email": value, "mobile": value,
			}
		}
	}

	opt.SelectFields = limitFileds
	var pbUsers []*pb.User
	userList, total, err := s.dbModel.GetUserList(opt)
	if err == gorm.ErrRecordNotFound {
		err = nil
		return &pb.ListUserReply{}, nil
	}
	util.CopyStruct(&userList, &pbUsers)

	// 查询用户ID
	var (
		userIds   []interface{}
		userIndex = make(map[int64]int)
	)

	for index, pbUser := range pbUsers {
		userIds = append(userIds, pbUser.Id)
		userIndex[pbUser.Id] = index
	}

	userGroups, _, _ := s.dbModel.GetUserGroupList(&model.OptionGetUserGroupList{
		QueryIn: map[string][]interface{}{"user_id": userIds},
	})

	for _, userGroup := range userGroups {
		index := userIndex[userGroup.UserId]
		pbUsers[index].GroupId = append(pbUsers[index].GroupId, userGroup.GroupId)
	}

	s.logger.Debug("ListUser", zap.Any("userList", userList), zap.Any("pbUser", pbUsers), zap.Int64("total", total))
	return &pb.ListUserReply{Total: total, User: pbUsers}, nil
}

// GetUserCaptcha 获取用户验证码
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
	case "find_password":
		res.Enable = cfgSecurity.EnableCaptchaFindPassword
	case "comment":
		res.Enable = cfgSecurity.EnableCaptchaComment
	default:
		return nil, status.Errorf(codes.InvalidArgument, ErrorMessageUnsupportedCaptchaType)
	}

	if res.Enable {
		res.Id, res.Captcha, err = captcha.GenerateCaptcha(cfgCaptcha.Type, cfgCaptcha.Length, cfgCaptcha.Width, cfgCaptcha.Height)
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return res, nil
}

// GetUserPermissions 获取用户权限
func (s *UserAPIService) GetUserPermissions(ctx context.Context, req *emptypb.Empty) (*pb.GetUserPermissionsReply, error) {
	userClaims, ok := ctx.Value(auth.CtxKeyUserClaims).(*auth.UserClaims)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "您未登录或您的登录已过期")
	}

	permissions, err := s.dbModel.GetUserPermissinsByUserId(userClaims.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbPermissions []*pb.Permission
	util.CopyStruct(&permissions, &pbPermissions)

	return &pb.GetUserPermissionsReply{Permission: pbPermissions}, nil
}

// AddUser 新增用户
func (s *UserAPIService) AddUser(ctx context.Context, req *pb.SetUserRequest) (*emptypb.Empty, error) {
	s.logger.Debug("AddUser", zap.Any("req", req))

	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = validate.ValidateStruct(req, s.getValidFieldMap())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if len(req.GroupId) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "用户组不能为空")
	}

	if !util.IsValidEmail(req.Email) {
		return nil, status.Errorf(codes.InvalidArgument, "邮箱格式不正确")
	}

	existUser, _ := s.dbModel.GetUserByEmail(req.Email, "id")
	if existUser.Id > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "邮箱已存在")
	}

	existUser, _ = s.dbModel.GetUserByUsername(req.Username, "id")
	if existUser.Id > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "用户名已存在")
	}

	// 新增用户
	user := &model.User{Username: req.Username, Password: req.Password, Email: req.Email}
	err = s.dbModel.CreateUser(user, req.GroupId...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *UserAPIService) CanIUploadDocument(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	userClaims, err := checkGRPCLogin(s.dbModel, ctx)
	if err != nil {
		return nil, err
	}

	if !s.dbModel.CanIUploadDocument(userClaims.UserId) {
		return nil, status.Errorf(codes.PermissionDenied, "您没有上传文档的权限")
	}

	return &emptypb.Empty{}, nil
}

func (s *UserAPIService) GetSignedToday(ctx context.Context, req *emptypb.Empty) (*v1.Sign, error) {
	userClaims, err := checkGRPCLogin(s.dbModel, ctx)
	if err != nil {
		return nil, err
	}

	sign := s.dbModel.GetSignedToday(userClaims.UserId)
	if sign.Id == 0 {
		return nil, status.Errorf(codes.NotFound, "您今天还没有签到")
	}
	pbSign := &v1.Sign{}
	util.CopyStruct(&sign, pbSign)
	return pbSign, nil
}

func (s *UserAPIService) SignToday(ctx context.Context, req *emptypb.Empty) (*v1.Sign, error) {
	userClaims, err := checkGRPCLogin(s.dbModel, ctx)
	if err != nil {
		return nil, err
	}

	if sign := s.dbModel.GetSignedToday(userClaims.UserId); sign.Id > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "您今天已经签到过了")
	}
	ip := ""
	if ips, _ := util.GetGRPCRemoteIP(ctx); len(ips) > 0 {
		ip = ips[0]
	}

	sign, err := s.dbModel.CreateSign(userClaims.UserId, ip)
	if err != nil {
		s.logger.Error("签到失败", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	pbSign := &v1.Sign{}
	util.CopyStruct(sign, pbSign)
	return pbSign, nil
}

func (s *UserAPIService) ListUserDynamic(ctx context.Context, req *v1.ListUserDynamicRequest) (*v1.ListUserDynamicReply, error) {
	userId := req.Id
	if userId == 0 {
		userClaims, _ := checkGRPCLogin(s.dbModel, ctx)
		if userClaims == nil {
			return nil, status.Error(codes.Unauthenticated, "用户ID参数不正确")
		}
		userId = userClaims.UserId
	}

	opt := &model.OptionGetDynamicList{
		WithCount: true,
		Page:      int(req.Page),
		Size:      int(req.Size_),
		QueryIn: map[string][]interface{}{
			"user_id": {userId},
		},
	}

	opt.Size = util.LimitRange(opt.Size, 10, 100)
	opt.Page = util.LimitMin(opt.Page, 1)

	dynamics, total, err := s.dbModel.GetDynamicList(opt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbDynamics []*v1.Dynamic
	util.CopyStruct(&dynamics, &pbDynamics)

	return &v1.ListUserDynamicReply{Dynamic: pbDynamics, Total: total}, nil
}

// 找回密码：发送邮件
func (s *UserAPIService) FindPasswordStepOne(ctx context.Context, req *v1.FindPasswordRequest) (res *emptypb.Empty, err error) {
	if !util.IsValidEmail(req.Email) {
		return nil, status.Errorf(codes.InvalidArgument, "邮箱格式不正确")
	}

	cfgSec := s.dbModel.GetConfigOfSecurity(model.ConfigSecurityEnableCaptchaFindPassword)
	if cfgSec.EnableCaptchaFindPassword && !captcha.VerifyCaptcha(req.CaptchaId, req.Captcha) {
		return nil, status.Errorf(codes.InvalidArgument, "验证码错误")
	}

	user, _ := s.dbModel.GetUserByEmail(req.Email, "id")
	if user.Id == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	cfgEmail := s.dbModel.GetConfigOfEmail()
	if !cfgEmail.Enable {
		return nil, status.Errorf(codes.Internal, "邮件服务未启用，请联系管理员")
	}
	cfgSystem := s.dbModel.GetConfigOfSystem()
	cfgEmail.Duration = util.LimitRange(cfgEmail.Duration, 1, 60)
	// 使用JWT创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        req.Email,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(cfgEmail.Duration)).Unix(),
	})
	tokenString, err := token.SignedString([]byte(cfgEmail.Secret))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	link := fmt.Sprintf("%s/findpassword?token=%s&email=%s", strings.TrimRight(cfgSystem.Domain, "/"), tokenString, url.QueryEscape(req.Email))
	body := fmt.Sprintf(`
	<div class="wrapper" style="margin: 20px auto 0; width: 500px; padding-top:16px; padding-bottom:10px;">
        <div class="content" style="background: none repeat scroll 0 0 #FFFFFF; border: 1px solid #E9E9E9; margin: 2px 0 0; padding: 30px;">
            <p>您好: </p>
            <p>您在 <a href="%s">%s</a> 提交了找回密码申请。<br>如果您没有提交修改密码的申请, 请忽略本邮件</p>
            <p style="border-top: 1px solid #DDDDDD;margin: 15px 0 25px;padding: 15px;">
                请点击链接继续: <a href="%s" target="_blank">%s</a>
            </p>
            <p>
                好的密码，不但应该容易记住，还要尽量符合以下强度标准：
            <ul>
                <li>包含大小写字母、数字和符号</li>
                <li>不少于 6 位 </li>
                <li>不包含生日、手机号码等易被猜出的信息</li>
            </ul>
            </p>
        </div>
    </div>
	`,
		cfgSystem.Domain,
		cfgSystem.Sitename,
		link,
		link,
	)

	// 发送邮件
	err = s.dbModel.SendMail(
		"找回密码 - "+cfgSystem.Sitename,
		req.Email,
		body,
	)
	if err != nil {
		s.logger.Error("发送邮件失败", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// 找回密码：重置密码
func (s *UserAPIService) FindPasswordStepTwo(ctx context.Context, req *v1.FindPasswordRequest) (res *emptypb.Empty, err error) {
	if !util.IsValidEmail(req.Email) {
		return nil, status.Errorf(codes.InvalidArgument, "邮箱格式不正确")
	}

	if len(req.Password) < 6 {
		return nil, status.Errorf(codes.InvalidArgument, "密码长度不能小于6位")
	}

	cfgSec := s.dbModel.GetConfigOfSecurity(model.ConfigSecurityEnableCaptchaFindPassword)
	if cfgSec.EnableCaptchaFindPassword && !captcha.VerifyCaptcha(req.CaptchaId, req.Captcha) {
		return nil, status.Errorf(codes.InvalidArgument, "验证码错误")
	}

	// 验证token
	claims := &jwt.StandardClaims{}
	cfgEmail := s.dbModel.GetConfigOfEmail(model.ConfigEmailSecret)
	// 验证JWT是否合法
	jwtToken, err := jwt.ParseWithClaims(req.Token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfgEmail.Secret), nil
	})
	if err != nil || !jwtToken.Valid {
		return nil, status.Errorf(codes.InvalidArgument, "找回密码 token 无效")
	}

	user, _ := s.dbModel.GetUserByEmail(req.Email, "id")
	if user.Id == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	// 更新密码
	err = s.dbModel.UpdateUserPassword(user.Id, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

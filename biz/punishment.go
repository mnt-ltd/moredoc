package biz

import (
	"context"
	"strings"
	"time"

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

type PunishmentAPIService struct {
	pb.UnimplementedPunishmentAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewPunishmentAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *PunishmentAPIService) {
	return &PunishmentAPIService{dbModel: dbModel, logger: logger.Named("PunishmentAPIService")}
}

func (s *PunishmentAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

func (s *PunishmentAPIService) CreatePunishment(ctx context.Context, req *pb.CreatePunishmentRequest) (*emptypb.Empty, error) {
	userClaims, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	if len(req.UserId) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "请选择用户")
	}

	if len(req.Type) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "请选择处罚类型")
	}

	now := time.Now()
	startTime := &now
	if req.StartTime != nil {
		startTime = req.StartTime
	}
	for _, userId := range req.UserId {
		if userId == 1 {
			continue
		}
		for _, typ := range req.Type {
			punishment := &model.Punishment{
				UserId:    userId,
				Type:      int(typ),
				Enable:    req.Enable,
				Reason:    req.Reason,
				Remark:    req.Remark,
				StartTime: startTime,
				EndTime:   req.EndTime,
			}
			s.logger.Debug("CreatePunishment", zap.Any("punishment", punishment), zap.Any("req", req))
			punishment.Operators = s.dbModel.MakePunishmentOperators(userClaims.UserId, typ)
			err = s.dbModel.CreatePunishment(punishment)
			if err != nil {
				s.logger.Error("CreatePunishment", zap.Error(err))
				return nil, status.Errorf(codes.Internal, err.Error())
			}
		}
	}
	return &emptypb.Empty{}, nil
}

func (s *PunishmentAPIService) UpdatePunishment(ctx context.Context, req *pb.Punishment) (*emptypb.Empty, error) {
	userClaims, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	punishment := &model.Punishment{}
	err = util.CopyStruct(req, punishment)
	if err != nil {
		s.logger.Error("CopyStruct", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	s.logger.Debug("UpdatePunishment", zap.Any("punishment", punishment), zap.Any("req", req))

	if existPunishment, _ := s.dbModel.GetPunishment(punishment.Id, "id", "operators"); existPunishment.Id > 0 {
		punishment.Operators = s.dbModel.MakePunishmentOperators(userClaims.UserId, req.Type, existPunishment.Operators)
	}

	err = s.dbModel.UpdatePunishment(punishment)
	if err != nil {
		s.logger.Error("UpdatePunishment", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *PunishmentAPIService) GetPunishment(ctx context.Context, req *pb.GetPunishmentRequest) (*pb.Punishment, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	punishment, err := s.dbModel.GetPunishment(req.Id)
	if err != nil {
		s.logger.Error("GetPunishment", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.Punishment{}
	err = util.CopyStruct(punishment, res)
	if err != nil {
		s.logger.Error("CopyStruct", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if punishment.UserId > 0 {
		user, _ := s.dbModel.GetUser(punishment.UserId, "id", "username")
		if user.Id > 0 {
			res.Username = user.Username
		}
	}
	return res, nil
}

func (s *PunishmentAPIService) ListPunishment(ctx context.Context, req *pb.ListPunishmentRequest) (*pb.ListPunishmentReply, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	opt := &model.OptionGetPunishmentList{
		Page:         int(req.Page),
		Size:         int(req.Size_),
		WithCount:    true,
		SelectFields: req.Field,
		QueryLike:    make(map[string][]interface{}),
		QueryIn:      make(map[string][]interface{}),
	}

	if len(req.UserId) > 0 {
		opt.QueryIn["user_id"] = util.Slice2Interface(req.UserId)
	}

	if len(req.Type) > 0 {
		opt.QueryIn["type"] = util.Slice2Interface(req.Type)
	}

	if len(req.Enable) > 0 {
		opt.QueryIn["enable"] = util.Slice2Interface(req.Enable)
	}

	if req.Order != "" {
		opt.Sort = strings.Split(req.Order, ",")
	}

	if req.Wd != "" {
		wd := strings.TrimSpace(req.Wd)
		opt.QueryLike["reason"] = []interface{}{wd}
		opt.QueryLike["remark"] = []interface{}{wd}
	}

	data, total, err := s.dbModel.GetPunishmentList(opt)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error("GetPunishmentList", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListPunishmentReply{
		Total: total,
	}

	err = util.CopyStruct(data, &res.Punishment)
	if err != nil {
		s.logger.Error("CopyStruct", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if len(res.Punishment) > 0 {
		var (
			userIds          []int64
			userIdMapIndexes = make(map[int64][]int)
		)

		for i, v := range res.Punishment {
			userIds = append(userIds, v.UserId)
			userIdMapIndexes[v.UserId] = append(userIdMapIndexes[v.UserId], i)
		}

		users, _, _ := s.dbModel.GetUserList(&model.OptionGetUserList{
			Ids:       userIds,
			WithCount: false,
			SelectFields: []string{
				"id",
				"username",
			},
		})

		for _, v := range users {
			if indexes, ok := userIdMapIndexes[v.Id]; ok {
				for _, index := range indexes {
					res.Punishment[index].Username = v.Username
				}
			}
		}

	}

	return res, nil
}

// 取消惩罚
func (s *PunishmentAPIService) CancelPunishment(ctx context.Context, req *pb.CancelPunishmentRequest) (*emptypb.Empty, error) {
	userCliams, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("CancelPunishment", zap.Any("req", req))

	data, _, err := s.dbModel.GetPunishmentList(&model.OptionGetPunishmentList{
		Ids: req.Id,
	})

	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error("GetPunishmentList", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, item := range data {
		item.Enable = false
		item.Operators = s.dbModel.MakePunishmentOperators(userCliams.UserId, 0, item.Operators)
		err = s.dbModel.UpdatePunishment(&item, "enable", "operators")
		if err != nil {
			s.logger.Error("UpdatePunishment", zap.Error(err))
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &emptypb.Empty{}, nil
}

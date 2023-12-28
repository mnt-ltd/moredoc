package biz

import (
	"context"
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

type AdvertisementAPIService struct {
	pb.UnimplementedAdvertisementAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewAdvertisementAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *AdvertisementAPIService) {
	return &AdvertisementAPIService{dbModel: dbModel, logger: logger.Named("AdvertisementAPIService")}
}

func (s *AdvertisementAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

func (s *AdvertisementAPIService) CreateAdvertisement(ctx context.Context, req *pb.Advertisement) (*pb.Advertisement, error) {
	userCliams, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	adv := &model.Advertisement{}
	if err = util.CopyStruct(req, adv); err != nil {
		s.logger.Error("CreateAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "创建广告失败:"+err.Error())
	}

	if adv.Position == "" || adv.Content == "" {
		return nil, status.Errorf(codes.InvalidArgument, "广告位和广告内容均不能为空")
	}

	adv.UserId = userCliams.UserId
	if err = s.dbModel.CreateAdvertisement(adv); err != nil {
		s.logger.Error("CreateAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "创建广告失败:"+err.Error())
	}

	res := &pb.Advertisement{}
	if err = util.CopyStruct(adv, res); err != nil {
		s.logger.Error("CreateAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return res, nil
}

func (s *AdvertisementAPIService) UpdateAdvertisement(ctx context.Context, req *pb.Advertisement) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	adv := &model.Advertisement{}
	if err = util.CopyStruct(req, adv); err != nil {
		s.logger.Error("UpdateAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "更新广告失败:"+err.Error())
	}

	if err = s.dbModel.UpdateAdvertisement(adv); err != nil {
		s.logger.Error("UpdateAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "更新广告失败:"+err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *AdvertisementAPIService) DeleteAdvertisement(ctx context.Context, req *pb.DeleteAdvertisementRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	if len(req.Id) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "广告id不能为空")
	}

	if err = s.dbModel.DeleteAdvertisement(req.Id); err != nil {
		s.logger.Error("DeleteAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "删除广告失败:"+err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *AdvertisementAPIService) GetAdvertisementByPosition(ctx context.Context, req *pb.GetAdvertisementByPositionRequest) (*pb.ListAdvertisementReply, error) {
	// 根据广告位获取广告
	now := time.Now()
	opt := &model.OptionGetAdvertisementList{
		WithCount: false,
		Page:      1,
		Size:      100000,
		QueryIn: map[string][]interface{}{
			"enable": {true},
		},
		QueryRange: map[string][2]interface{}{
			"start_time": {nil, now},
			"end_time":   {now, nil},
		},
	}

	if len(req.Position) > 0 {
		opt.QueryIn["position"] = util.Slice2Interface(req.Position)
	}

	advs, _, err := s.dbModel.GetAdvertisementList(opt)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error("GetAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "获取广告失败:"+err.Error())
	}

	res := &pb.ListAdvertisementReply{}
	err = util.CopyStruct(advs, &res.Advertisement)
	if err != nil {
		s.logger.Error("GetAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for idx := range res.Advertisement {
		// 去除备注
		res.Advertisement[idx].Remark = ""
	}

	return res, nil
}

func (s *AdvertisementAPIService) GetAdvertisement(ctx context.Context, req *pb.GetAdvertisementRequest) (*pb.Advertisement, error) {
	if req.Id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "广告id不能为空")
	}

	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	ad, err := s.dbModel.GetAdvertisement(req.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error("GetAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "获取广告失败:"+err.Error())
	}

	res := &pb.Advertisement{}
	err = util.CopyStruct(ad, res)
	if err != nil {
		s.logger.Error("GetAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *AdvertisementAPIService) ListAdvertisement(ctx context.Context, req *pb.ListAdvertisementRequest) (*pb.ListAdvertisementReply, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	opt := &model.OptionGetAdvertisementList{
		WithCount: true,
		Page:      int(req.Page),
		Size:      int(req.Size_),
		QueryIn:   make(map[string][]interface{}),
		QueryLike: make(map[string][]interface{}),
	}

	if len(req.Position) > 0 {
		opt.QueryIn["position"] = util.Slice2Interface(req.Position)
	}

	if len(req.Enable) > 0 {
		opt.QueryIn["enable"] = util.Slice2Interface(req.Enable)
	}

	if req.Wd != "" {
		fields := []string{"title", "content", "remark"}
		for _, field := range fields {
			opt.QueryLike[field] = []interface{}{req.Wd}
		}
	}

	advs, total, err := s.dbModel.GetAdvertisementList(opt)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error("ListAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "获取广告列表失败:"+err.Error())
	}

	res := &pb.ListAdvertisementReply{
		Total: total,
	}
	err = util.CopyStruct(advs, &res.Advertisement)
	if err != nil {
		s.logger.Error("ListAdvertisement", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return res, nil
}

package biz

import (
	"context"
	"strings"

	pb "moredoc/api/v1"
	"moredoc/model"
	"moredoc/util"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type SearchRecordAPIService struct {
	pb.UnimplementedSearchRecordAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewSearchRecordAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *SearchRecordAPIService) {
	return &SearchRecordAPIService{dbModel: dbModel, logger: logger.Named("SearchRecordAPIService")}
}

func (s *SearchRecordAPIService) DeleteSearchRecord(ctx context.Context, req *pb.DeleteSearchRecordRequest) (*emptypb.Empty, error) {
	_, err := checkGRPCPermission(s.dbModel, ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteSearchRecord(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除失败:"+err.Error())
	}

	return &emptypb.Empty{}, nil
}
func (s *SearchRecordAPIService) ListSearchRecord(ctx context.Context, req *pb.ListSearchRecordRequest) (*pb.ListSearchRecordReply, error) {
	_, err := checkGRPCPermission(s.dbModel, ctx)
	if err != nil {
		return nil, err
	}

	opt := &model.OptionGetSearchRecordList{
		Size:      int(req.Size_),
		Page:      int(req.Page),
		WithCount: true,
		Sort:      []string{"id desc"},
		QueryIn:   map[string][]interface{}{},
		QueryLike: map[string][]interface{}{},
	}
	if req.Order != "" {
		opt.Sort = strings.Split(req.Order, ",")
	}

	if req.Keywords != "" {
		opt.QueryLike["keyword"] = []interface{}{req.Keywords}
	}

	if len(req.UserId) > 0 {
		opt.QueryIn["user_id"] = util.Slice2Interface(req.UserId)
	}

	if req.Ip != "" {
		opt.QueryLike["ip"] = []interface{}{req.Ip}
	}

	records, total, err := s.dbModel.GetSearchRecordList(opt)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, "获取列表失败:"+err.Error())
	}

	res := &pb.ListSearchRecordReply{
		Total: total,
	}

	err = util.CopyStruct(&records, &res.SearchRecord)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取列表失败:"+err.Error())
	}

	return res, nil
}

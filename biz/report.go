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

type ReportAPIService struct {
	pb.UnimplementedReportAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewReportAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *ReportAPIService) {
	return &ReportAPIService{dbModel: dbModel, logger: logger.Named("ReportAPIService")}
}

func (s *ReportAPIService) checkLogin(ctx context.Context) (*auth.UserClaims, error) {
	return checkGRPCLogin(s.dbModel, ctx)
}

func (s *ReportAPIService) checkPermission(ctx context.Context) (*auth.UserClaims, error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

func (s *ReportAPIService) CreateReport(ctx context.Context, req *pb.Report) (*emptypb.Empty, error) {
	UserClaims, err := s.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	if req.DocumentId == 0 {
		return nil, status.Error(codes.InvalidArgument, "文档参数不正确")
	}

	report, _ := s.dbModel.GetReportByDocUser(req.DocumentId, UserClaims.UserId)
	if report.Id > 0 {
		return nil, status.Error(codes.AlreadyExists, "您已举报过当前文档")
	}

	util.CopyStruct(req, &report)
	report.UserId = UserClaims.UserId
	err = s.dbModel.CreateReport(&report)
	if err != nil {
		return nil, status.Error(codes.Internal, "创建举报失败")
	}

	return &emptypb.Empty{}, nil
}

func (s *ReportAPIService) UpdateReport(ctx context.Context, req *pb.Report) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	report := &model.Report{}
	util.CopyStruct(req, report)
	err = s.dbModel.UpdateReport(report, "status", "remark")
	if err != nil {
		return nil, status.Error(codes.Internal, "更新举报失败")
	}

	return &emptypb.Empty{}, nil
}

func (s *ReportAPIService) DeleteReport(ctx context.Context, req *pb.DeleteReportRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteReport(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "删除举报失败")
	}

	return &emptypb.Empty{}, nil
}

func (s *ReportAPIService) ListReport(ctx context.Context, req *pb.ListReportRequest) (*pb.ListReportReply, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	opt := &model.OptionGetReportList{
		WithCount: true,
		Page:      int(req.Page),
		Size:      int(req.Size_),
		QueryLike: make(map[string][]interface{}),
		QueryIn:   make(map[string][]interface{}),
	}

	if req.Wd != "" {
		opt.QueryLike["document_title"] = []interface{}{req.Wd}
	}

	if len(req.Status) > 0 {
		opt.QueryIn["status"] = util.Slice2Interface(req.Status)
	}

	reports, total, err := s.dbModel.GetReportList(opt)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Error(codes.Internal, "获取举报列表失败")
	}

	pbReport := &pb.ListReportReply{
		Total:  total,
		Report: reports,
	}
	return pbReport, nil
}

package biz

import (
	"context"

	pb "moredoc/api/v1"
	"moredoc/model"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DocumentAPIService struct {
	pb.UnimplementedDocumentAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewDocumentAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *DocumentAPIService) {
	return &DocumentAPIService{dbModel: dbModel, logger: logger.Named("DocumentAPIService")}
}

func (s *DocumentAPIService) CreateDocument(ctx context.Context, req *pb.CreateDocumentRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *DocumentAPIService) UpdateDocument(ctx context.Context, req *pb.Document) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *DocumentAPIService) DeleteDocument(ctx context.Context, req *pb.DeleteDocumentRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *DocumentAPIService) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (*pb.Document, error) {
	return &pb.Document{}, nil
}

func (s *DocumentAPIService) ListDocument(ctx context.Context, req *pb.ListDocumentRequest) (*pb.ListDocumentReply, error) {
	return &pb.ListDocumentReply{}, nil
}

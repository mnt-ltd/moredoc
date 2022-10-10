package biz

import (
	"context"
	v1 "moredoc/api/v1"
	"moredoc/model"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthAPIService struct {
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewHealthAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *HealthAPIService) {
	return &HealthAPIService{
		dbModel: dbModel,
		logger:  logger.Named("biz"),
	}
}

// Health is health check
func (p *HealthAPIService) Health(ctx context.Context, in *empty.Empty) (out *empty.Empty, err error) {
	out = &emptypb.Empty{}
	return
}

// Ping ping pong
func (p *HealthAPIService) Ping(ctx context.Context, in *v1.PingRequest) (out *v1.PongReply, err error) {
	createdAt := time.Now()
	out = &v1.PongReply{
		Name:      in.Name,
		CreatedAt: &createdAt,
	}
	return
}

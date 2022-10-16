package biz

import (
	"context"
	"strings"

	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AttachmentAPIService struct {
	pb.UnimplementedAttachmentAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewAttachmentAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *AttachmentAPIService) {
	return &AttachmentAPIService{dbModel: dbModel, logger: logger.Named("AttachmentAPIService")}
}

// checkPermission 检查用户权限
func (s *AttachmentAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	var ok bool
	userClaims, ok = ctx.Value(auth.CtxKeyUserClaims).(*auth.UserClaims)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, ErrorMessageInvalidToken)
	}

	fullMethod, _ := ctx.Value(auth.CtxKeyFullMethod).(string)
	if yes := s.dbModel.CheckPermissionByUserId(userClaims.UserId, fullMethod); !yes {
		return nil, status.Errorf(codes.PermissionDenied, ErrorMessagePermissionDenied)
	}
	return
}

func (s *AttachmentAPIService) UpdateAttachment(ctx context.Context, req *pb.Attachment) (*pb.Attachment, error) {
	return &pb.Attachment{}, nil
}

func (s *AttachmentAPIService) DeleteAttachment(ctx context.Context, req *pb.DeleteAttachmentRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteAttachment(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *AttachmentAPIService) GetAttachment(ctx context.Context, req *pb.GetAttachmentRequest) (*pb.Attachment, error) {
	return &pb.Attachment{}, nil
}

func (s *AttachmentAPIService) ListAttachment(ctx context.Context, req *pb.ListAttachmentRequest) (*pb.ListAttachmentReply, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	opt := &model.OptionGetAttachmentList{
		Page:      int(req.Page),
		Size:      int(req.Size_),
		WithCount: true,
		QueryIn:   make(map[string][]interface{}),
	}

	if len(req.UserId) > 0 {
		opt.QueryIn["user_id"] = util.Slice2Interface(req.UserId)
	}

	if len(req.IsApproved) > 0 {
		opt.QueryIn["is_approved"] = util.Slice2Interface(req.IsApproved)
	}

	if len(req.Type) > 0 {
		opt.QueryIn["type"] = util.Slice2Interface(req.Type)
	}

	req.Wd = strings.TrimSpace(req.Wd)
	if req.Wd != "" {
		wd := "%" + req.Wd + "%"
		opt.QueryLike = map[string][]interface{}{"name": {wd}}
	}

	attachments, total, err := s.dbModel.GetAttachmentList(opt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbAttachments []*pb.Attachment
	util.CopyStruct(&attachments, &pbAttachments)

	var (
		userIds        []interface{}
		userIdIndexMap = make(map[int64][]int)
	)

	for idx, attchment := range pbAttachments {
		attchment.TypeName = s.dbModel.GetAttachmentTypeName(int(attchment.Type))
		userIds = append(userIds, attchment.UserId)
		userIdIndexMap[attchment.UserId] = append(userIdIndexMap[attchment.UserId], idx)
		pbAttachments[idx] = attchment
	}

	if size := len(userIds); size > 0 {
		users, _, _ := s.dbModel.GetUserList(&model.OptionGetUserList{Ids: userIds, Page: 1, Size: size, SelectFields: []string{"id", "username"}})
		s.logger.Debug("GetUserList", zap.Any("users", users))
		for _, user := range users {
			if indexes, ok := userIdIndexMap[user.Id]; ok {
				for _, idx := range indexes {
					pbAttachments[idx].Username = user.Username
				}
			}
		}
	}

	return &pb.ListAttachmentReply{Total: total, Attachment: pbAttachments}, nil
}

// 上传头像
func (s *AttachmentAPIService) UploadAvatar() {

}

// 上传横幅
func (s *AttachmentAPIService) UploadBanner() {

}

// 上传文档
func (s *AttachmentAPIService) UploadDocument() {

}

// 上传文档分类封面
func (s *AttachmentAPIService) UploadCategoryCover() {

}

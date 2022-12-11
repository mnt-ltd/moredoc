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

type DocumentAPIService struct {
	pb.UnimplementedDocumentAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewDocumentAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *DocumentAPIService) {
	return &DocumentAPIService{dbModel: dbModel, logger: logger.Named("DocumentAPIService")}
}

func (s *DocumentAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

func (s *DocumentAPIService) checkLogin(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCLogin(s.dbModel, ctx)
}

// CreateDocument 创建文档
// 0. 判断是否有权限
// 1. 同名覆盖：找到该作者上传的相同title和ext的文档，然后用新文件覆盖，同时文档状态改为待转换
// 2. 相同hash的文档如果已经被转换了，则该文档的状态直接改为已转换
// 3. 判断附件ID是否与用户ID匹配，不匹配则跳过该文档
func (s *DocumentAPIService) CreateDocument(ctx context.Context, req *pb.CreateDocumentRequest) (*emptypb.Empty, error) {
	userCliams, err := s.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	if !s.dbModel.CanIUploadDocument(userCliams.UserId) {
		return nil, status.Error(codes.PermissionDenied, "没有权限上传文档")
	}

	var (
		attachmentIds []interface{}
		attachmentMap = make(map[int64]model.Attachment)
	)

	for _, item := range req.Document {
		attachmentIds = append(attachmentIds, item.AttachmentId)
	}

	attachments, _, _ := s.dbModel.GetAttachmentList(&model.OptionGetAttachmentList{
		Ids:     attachmentIds,
		QueryIn: map[string][]interface{}{"user_id": {userCliams.UserId}},
	})
	if len(attachments) == 0 {
		return nil, status.Error(codes.InvalidArgument, "文档文件参数attachment_id不正确")
	}

	for _, attachment := range attachments {
		attachmentMap[attachment.Id] = attachment
	}

	var (
		documents        []model.Document
		docMapAttachment = make(map[int]int64)
		jieba            = util.NewJieba()
	)
	for idx, doc := range req.Document {
		attachment, ok := attachmentMap[doc.AttachmentId]
		if !ok {
			continue
		}

		doc := model.Document{
			Title:    doc.Title,
			Keywords: strings.Join(jieba.SegWords(doc.Title, 10), ","),
			UserId:   userCliams.UserId,
			// UUID:     uuid.Must(uuid.NewV4()).String(),
			Score:  300,
			Price:  int(doc.Price),
			Size:   attachment.Size,
			Ext:    attachment.Ext,
			Status: model.DocumentStatusPending,
		}
		docMapAttachment[idx] = attachment.Id
		documents = append(documents, doc)
	}

	docs, err := s.dbModel.CreateDocuments(documents, req.CategoryId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	attachIdTypeIdMap := make(map[int64]int64)
	for idx, doc := range docs {
		if attachmentId, ok := docMapAttachment[idx]; ok {
			attachIdTypeIdMap[attachmentId] = doc.Id
		}
	}

	s.dbModel.SetAttachmentTypeId(attachIdTypeIdMap)

	return &emptypb.Empty{}, nil
}

// UpdateDocument 更新文档
// 1. 对于普通用户，可以更新自己创建的文档
// 2. 对于管理员，可以更新所有文档
func (s *DocumentAPIService) UpdateDocument(ctx context.Context, req *pb.Document) (*emptypb.Empty, error) {
	s.logger.Debug("UpdateDocument", zap.Any("req", req))
	userClaims, err := s.checkPermission(ctx)
	if userClaims == nil { // 未登录
		return nil, err
	}

	fields := []string{"id", "title", "keywords", "description", "price"}
	doc := &model.Document{}
	util.CopyStruct(req, doc)

	if err != nil { // 普通用户，只能更新自己的文档
		existDoc, _ := s.dbModel.GetDocument(req.Id, "id", "user_id")
		if existDoc.UserId != userClaims.UserId {
			return nil, status.Error(codes.PermissionDenied, "文档不存在或没有权限")
		}
	} else { // 管理员，可以更新所有文档
		fields = append(fields, "status")
	}

	err = s.dbModel.UpdateDocument(doc, req.CategoryId, fields...)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// DeleteDocument 删除文档
// 1. 对于普通用户，可以删除自己创建的文档
// 2. 对于管理员，可以删除所有文档
func (s *DocumentAPIService) DeleteDocument(ctx context.Context, req *pb.DeleteDocumentRequest) (*emptypb.Empty, error) {
	userClaims, err := s.checkPermission(ctx)
	s.logger.Info("DeleteDocument", zap.Any("userClaims", userClaims), zap.Error(err))
	if err != nil && userClaims == nil { // 未登录
		return nil, err
	}

	ids := req.Id
	if err != nil { // 普通用户，只能删除自己创建的文档
		userDocs, _, _ := s.dbModel.GetDocumentList(&model.OptionGetDocumentList{
			WithCount:    false,
			SelectFields: []string{"id"},
			Ids:          util.Slice2Interface(req.Id),
		})

		if len(userDocs) == 0 {
			return &emptypb.Empty{}, nil
		}

		for _, doc := range userDocs {
			ids = append(ids, doc.Id)
		}
	}

	err = s.dbModel.DeleteDocument(ids, userClaims.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除文档失败：%v", err)
	}

	return &emptypb.Empty{}, nil
}

// GetDocument 获取文档（任何人都可以调用）。当文档禁用之后，只有管理员可以查看
func (s *DocumentAPIService) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (*pb.Document, error) {
	doc, _ := s.dbModel.GetDocument(req.Id)
	if doc.Id == 0 {
		return nil, status.Error(codes.NotFound, "文档不存在")
	}

	_, err := s.checkPermission(ctx)
	if err != nil && doc.Status == model.DocumentStatusDisabled {
		return nil, status.Error(codes.NotFound, "文档不存在或没有权限")
	}

	pbDoc := &pb.Document{}
	util.CopyStruct(doc, pbDoc)
	docCates, _, _ := s.dbModel.GetDocumentCategoryList(
		&model.OptionGetDocumentCategoryList{
			WithCount:    false,
			SelectFields: []string{"category_id"},
			QueryIn:      map[string][]interface{}{"document_id": {doc.Id}},
		},
	)

	for _, dc := range docCates {
		pbDoc.CategoryId = append(pbDoc.CategoryId, dc.CategoryId)
	}

	return pbDoc, nil
}

// ListDocument 查询文档列表
// 1. 对于普通用户，只能查询未禁用的文档，且最多只能查询100页
// 2. 对于管理员，可以查询所有文档，可以根据关键字进行查询
func (s *DocumentAPIService) ListDocument(ctx context.Context, req *pb.ListDocumentRequest) (*pb.ListDocumentReply, error) {
	opt := &model.OptionGetDocumentList{
		WithCount:    req.Limit <= 0,
		Page:         int(req.Page),
		Size:         int(req.Size_),
		SelectFields: req.Field,
		QueryIn:      make(map[string][]interface{}),
		QueryLike:    make(map[string][]interface{}),
		IsRecommend:  req.IsRecommend,
	}

	if len(req.Order) > 0 {
		opt.Sort = []string{req.Order}
	}

	if len(req.CategoryId) > 0 {
		opt.QueryIn["category_id"] = util.Slice2Interface(req.CategoryId)
	}

	if len(req.UserId) > 0 {
		opt.QueryIn["user_id"] = []interface{}{req.UserId[0]}
	}

	_, err := s.checkPermission(ctx)
	if err == nil { // 有权限，则不限页数
		if req.Wd != "" {
			opt.QueryLike["title"] = []interface{}{req.Wd}
			opt.QueryLike["keywords"] = []interface{}{req.Wd}
			opt.QueryLike["description"] = []interface{}{req.Wd}
		}

		if len(req.Status) > 0 {
			opt.QueryIn["status"] = util.Slice2Interface(req.Status)
		}
	} else {
		opt.Size = util.LimitRange(opt.Size, 1, 24)
		opt.Page = util.LimitRange(opt.Page, 1, 100)
		opt.QueryIn["status"] = []interface{}{
			model.DocumentStatusPending, model.DocumentStatusConverting,
			model.DocumentStatusConverted, model.DocumentStatusFailed,
		}
	}

	if req.Limit > 0 {
		opt.Size = int(req.Limit)
		opt.Page = 1
	}

	s.logger.Debug("ListDocument", zap.Any("opt", opt))
	return s.listDocument(opt)
}

// ListRecycleDocument 回收站文档
func (s *DocumentAPIService) ListRecycleDocument(ctx context.Context, req *pb.ListDocumentRequest) (*pb.ListDocumentReply, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	opt := &model.OptionGetDocumentList{
		WithCount:    true,
		Page:         int(req.Page),
		Size:         int(req.Size_),
		SelectFields: req.Field,
		QueryIn:      make(map[string][]interface{}),
		QueryLike:    make(map[string][]interface{}),
		IsRecycle:    true,
	}

	if len(req.CategoryId) > 0 {
		opt.QueryIn["category_id"] = util.Slice2Interface(req.CategoryId)
	}

	if len(req.UserId) > 0 {
		opt.QueryIn["user_id"] = []interface{}{req.UserId[0]}
	}

	if req.Wd != "" {
		opt.QueryLike["title"] = []interface{}{req.Wd}
		opt.QueryLike["keywords"] = []interface{}{req.Wd}
		opt.QueryLike["description"] = []interface{}{req.Wd}
	}

	if len(req.Status) > 0 {
		opt.QueryIn["status"] = util.Slice2Interface(req.Status)
	}

	return s.listDocument(opt)
}

// RecoverRecycleDocument 恢复回收站文档
func (s *DocumentAPIService) RecoverRecycleDocument(ctx context.Context, req *pb.RecoverRecycleDocumentRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	s.logger.Debug("RecoverRecycleDocument", zap.Any("req", req))

	if len(req.Id) == 0 {
		return &emptypb.Empty{}, nil
	}

	err = s.dbModel.RecoverRecycleDocument(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// DeleteRecycleDocument 删除回收站文档
func (s *DocumentAPIService) DeleteRecycleDocument(ctx context.Context, req *pb.DeleteDocumentRequest) (*emptypb.Empty, error) {
	userClaims, err := s.checkPermission(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	err = s.dbModel.DeleteDocument(req.Id, userClaims.UserId, true)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// ClearRecycleDocument 清空回收站文档
func (s *DocumentAPIService) ClearRecycleDocument(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	err = s.dbModel.ClearRecycleDocument()
	if err != nil {
		s.logger.Error("ClearRecycleDocument", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *DocumentAPIService) listDocument(opt *model.OptionGetDocumentList) (*pb.ListDocumentReply, error) {
	docs, total, err := s.dbModel.GetDocumentList(opt)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbDocs []*pb.Document
	err = util.CopyStruct(&docs, &pbDocs)
	if err != nil {
		s.logger.Error("CopyStruct failed", zap.Error(err))
	}

	var (
		docCates            []model.DocumentCategory
		docUsers            []model.User
		docIndexMap         = make(map[int64]int)
		userIndexesMap      = make(map[int64][]int)
		deletedUserIndexMap = make(map[int64][]int)
		docIds              []int64
		userIds             []int64
	)

	for i, doc := range pbDocs {
		docIndexMap[doc.Id] = i
		userIndexesMap[doc.UserId] = append(userIndexesMap[doc.UserId], i)
		docIds = append(docIds, doc.Id)
		if doc.UserId > 0 {
			userIds = append(userIds, doc.UserId)
		}
		if doc.DeletedUserId > 0 {
			userIds = append(userIds, doc.DeletedUserId)
			deletedUserIndexMap[doc.DeletedUserId] = append(deletedUserIndexMap[doc.DeletedUserId], i)
		}
	}

	if len(pbDocs) > 0 {
		docCates, _, _ = s.dbModel.GetDocumentCategoryList(&model.OptionGetDocumentCategoryList{
			WithCount:    false,
			SelectFields: []string{"document_id", "category_id"},
			QueryIn:      map[string][]interface{}{"document_id": util.Slice2Interface(docIds)},
		})
		for _, docCate := range docCates {
			pbDocs[docIndexMap[docCate.DocumentId]].CategoryId = append(pbDocs[docIndexMap[docCate.DocumentId]].CategoryId, docCate.CategoryId)
		}

		docUsers, _, _ = s.dbModel.GetUserList(&model.OptionGetUserList{
			WithCount:    false,
			SelectFields: []string{"id", "username"},
			QueryIn:      map[string][]interface{}{"id": util.Slice2Interface(userIds)},
		})

		// 查找文档相关联的附件。对于列表，只返回hash和id，不返回其他字段
		attachments, _, _ := s.dbModel.GetAttachmentList(&model.OptionGetAttachmentList{
			WithCount:    false,
			SelectFields: []string{"hash", "id", "type_id"},
			QueryIn: map[string][]interface{}{
				"type_id": util.Slice2Interface(docIds),
				"type":    {model.AttachmentTypeDocument},
			},
		})

		for _, attachment := range attachments {
			index := docIndexMap[attachment.TypeId]
			pbDocs[index].Attachment = &pb.Attachment{
				Hash: attachment.Hash,
			}
		}

		for _, docUser := range docUsers {
			indexes := userIndexesMap[docUser.Id]
			for _, index := range indexes {
				pbDocs[index].Username = docUser.Username
			}

			indexes = deletedUserIndexMap[docUser.Id]
			for _, index := range indexes {
				pbDocs[index].DeletedUsername = docUser.Username
			}
		}
	}

	return &pb.ListDocumentReply{
		Total:    total,
		Document: pbDocs,
	}, nil
}

// SetDocumentRecommend 推荐文档
func (s *DocumentAPIService) SetDocumentRecommend(ctx context.Context, req *pb.SetDocumentRecommendRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	err = s.dbModel.SetDocumentRecommend(req.Id, req.Type)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

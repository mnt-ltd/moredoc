package biz

import (
	"context"
	"strings"

	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"
	"moredoc/util/validate"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryAPIService struct {
	pb.UnimplementedCategoryAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewCategoryAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *CategoryAPIService) {
	return &CategoryAPIService{dbModel: dbModel, logger: logger.Named("CategoryAPIService")}
}

func (s *CategoryAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

func (s *CategoryAPIService) CreateCategory(ctx context.Context, req *pb.Category) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = validate.ValidateStruct(req, map[string]string{"Title": "名称"})
	if err != nil {
		return nil, err
	}

	cate := &model.Category{}
	util.CopyStruct(req, &cate)
	titles := strings.Split(cate.Title, "\n")
	for _, title := range titles {
		title = strings.TrimSpace(title)
		if title == "" {
			continue
		}

		cate.Title = title
		exist, _ := s.dbModel.GetCategoryByParentIdTitle(cate.ParentId, cate.Title, "id")
		if exist.Id > 0 {
			continue
		}

		cate.Id = 0
		err = s.dbModel.CreateCategory(cate)
		if err != nil {
			s.logger.Error("CreateCategory", zap.Error(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *CategoryAPIService) UpdateCategory(ctx context.Context, req *pb.Category) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = validate.ValidateStruct(req, map[string]string{"Title": "名称"})
	if err != nil {
		return nil, err
	}

	cate := &model.Category{}
	util.CopyStruct(req, &cate)

	exist, _ := s.dbModel.GetCategoryByParentIdTitle(cate.ParentId, cate.Title, "id")
	if exist.Id > 0 && exist.Id != cate.Id {
		return nil, status.Error(codes.Internal, "分类名称已存在")
	}

	err = s.dbModel.UpdateCategory(cate)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *CategoryAPIService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteCategory(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *CategoryAPIService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	cate, err := s.dbModel.GetCategory(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbCategory := &pb.Category{}
	util.CopyStruct(cate, &pbCategory)

	return pbCategory, nil
}

func (s *CategoryAPIService) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryReply, error) {
	opt := &model.OptionGetCategoryList{
		WithCount:    false,
		QueryIn:      make(map[string][]interface{}),
		SelectFields: req.Field,
		Page:         int(req.Page),
		Size:         int(req.Size_),
	}

	if len(req.ParentId) > 0 {
		opt.QueryIn["parent_id"] = util.Slice2Interface(req.ParentId)
	}

	// 管理员，可以通过关键字搜索
	if _, err := s.checkPermission(ctx); err == nil {
		if req.Wd != "" {
			opt.QueryLike = map[string][]interface{}{"title": {req.Wd}}
		}

		if len(req.Enable) > 0 {
			opt.QueryIn["enable"] = util.Slice2Interface(req.Enable)
		}
	} else {
		// 非管理员，只能查询启用的
		opt.QueryIn["enable"] = []interface{}{true}
	}

	cates, total, err := s.dbModel.GetCategoryList(opt)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var pbCates []*pb.Category
	util.CopyStruct(&cates, &pbCates)
	return &pb.ListCategoryReply{Total: total, Category: pbCates}, nil
}

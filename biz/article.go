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

type ArticleAPIService struct {
	pb.UnimplementedArticleAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewArticleAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *ArticleAPIService) {
	return &ArticleAPIService{dbModel: dbModel, logger: logger.Named("ArticleAPIService")}
}

func (s *ArticleAPIService) checkPermission(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

// CreateArticle 创建文章
func (s *ArticleAPIService) CreateArticle(ctx context.Context, req *pb.Article) (*pb.Article, error) {
	userClaims, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	if req.Identifier == "" {
		req.Identifier = util.GenDocumentMD5UUID()
	}

	existArticle, _ := s.dbModel.GetArticleByIdentifier(req.Identifier, "id")
	if existArticle.Id > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "文章标识符已存在")
	}

	article := &model.Article{}
	err = util.CopyStruct(req, article)
	if err != nil {
		s.logger.Error("CreateArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "创建文章失败")
	}

	if req.IsRecommend && req.RecommendAt == nil {
		now := time.Now()
		article.RecommendAt = &now
	}

	article.UserId = userClaims.UserId
	err = s.dbModel.CreateArticle(article)
	if err != nil {
		s.logger.Error("CreateArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "创建文章失败")
	}

	res := &pb.Article{}
	err = util.CopyStruct(article, res)
	if err != nil {
		s.logger.Error("CreateArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return res, nil
}

// UpdateArticle 更新文章。注意：不支持更新identifier
func (s *ArticleAPIService) UpdateArticle(ctx context.Context, req *pb.Article) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	article := &model.Article{}
	err = util.CopyStruct(req, article)
	if err != nil {
		s.logger.Error("UpdateArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "更新文章失败")
	}

	if req.IsRecommend && req.RecommendAt == nil {
		now := time.Now()
		article.RecommendAt = &now
	} else {
		article.RecommendAt = nil
	}

	err = s.dbModel.UpdateArticle(article)
	if err != nil {
		s.logger.Error("UpdateArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "更新文章失败")
	}

	return &emptypb.Empty{}, nil
}

// DeleteArticle 删除文章
func (s *ArticleAPIService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteArticle(req.Id)
	if err != nil {
		s.logger.Error("DeleteArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "删除文章失败")
	}

	return &emptypb.Empty{}, nil
}

// GetArticle 根据id或文章标识获取文章
// 不需要鉴权
func (s *ArticleAPIService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.Article, error) {
	var (
		article model.Article
		err     error
	)

	if req.Id > 0 {
		article, err = s.dbModel.GetArticle(req.Id)
		if err != nil && err != gorm.ErrRecordNotFound {
			s.logger.Error("GetArticle", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "获取文章失败")
		}
	} else {
		article, err = s.dbModel.GetArticleByIdentifier(req.Identifier)
		if err != nil && err != gorm.ErrRecordNotFound {
			s.logger.Error("GetArticle", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "获取文章失败")
		}
		article.ViewCount += 1
		s.dbModel.UpdateArticleViewCount(article.Id, article.ViewCount)
	}

	if article.Id == 0 {
		return nil, status.Errorf(codes.NotFound, "文章不存在")
	}

	pbArticle := &pb.Article{}
	util.CopyStruct(article, pbArticle)
	if pbArticle.UserId > 0 {
		user, _ := s.dbModel.GetUser(pbArticle.UserId, s.dbModel.GetUserPublicFields()...)
		pbArticle.User = &pb.User{}
		util.CopyStruct(&user, pbArticle.User)
	}
	return pbArticle, nil
}

// ListArticles 获取文章列表。
// 如果是有权限，则可以根据关键字来查询，否则只能简单查询
func (s *ArticleAPIService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	opt := &model.OptionGetArticleList{
		Page:        int(req.Page),
		Size:        int(req.Size_),
		WithCount:   true,
		QueryLike:   make(map[string][]interface{}),
		QueryIn:     make(map[string][]interface{}),
		Sort:        []string{req.Order},
		IsRecommend: req.IsRecommend,
	}

	_, err := s.checkPermission(ctx)
	if err == nil {
		if req.Wd != "" {
			opt.QueryLike["title"] = []interface{}{req.Wd}
			opt.QueryLike["keywords"] = []interface{}{req.Wd}
			opt.QueryLike["description"] = []interface{}{req.Wd}
		}
		if len(req.Status) > 0 {
			opt.QueryIn["status"] = util.Slice2Interface(req.Status)
		}
	} else {
		opt.QueryLike["status"] = []interface{}{1}
	}

	if len(req.CategoryId) > 0 {
		opt.QueryIn["category_id"] = util.Slice2Interface(req.CategoryId)
	}

	articles, total, err := s.dbModel.GetArticleList(opt)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error("ListArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "获取文章列表失败")
	}

	var pbArticle []*pb.Article
	err = util.CopyStruct(articles, &pbArticle)
	if err != nil {
		s.logger.Error("ListArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "获取文章列表失败"+err.Error())
	}

	var (
		userIds         []int64
		userMapArticles = make(map[int64][]int)
	)
	for idx, article := range articles {
		userIds = append(userIds, article.UserId)
		userMapArticles[article.UserId] = append(userMapArticles[article.UserId], idx)
	}

	if len(userIds) > 0 {
		users, _, _ := s.dbModel.GetUserList(&model.OptionGetUserList{
			QueryIn:      map[string][]interface{}{"id": util.Slice2Interface(userIds)},
			SelectFields: []string{"id", "username", "avatar"},
		})

		for _, user := range users {
			for _, idx := range userMapArticles[user.Id] {
				pbArticle[idx].User = &pb.User{
					Id:       user.Id,
					Username: user.Username,
					Avatar:   user.Avatar,
				}
			}
		}
	}

	return &pb.ListArticleReply{
		Total:   total,
		Article: pbArticle,
	}, nil
}

// SetArticlesCategory
func (s *ArticleAPIService) SetArticlesCategory(ctx context.Context, req *pb.SetArticlesCategoryRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.SetArticlesCategory(req.ArticleId, req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "设置文档分类失败：%s", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *ArticleAPIService) ListRecycleArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	opt := &model.OptionGetArticleList{
		Page:      int(req.Page),
		Size:      int(req.Size_),
		WithCount: true,
		QueryLike: make(map[string][]interface{}),
		Sort:      []string{req.Order},
		IsRecycle: true,
	}

	_, err := s.checkPermission(ctx)
	if err == nil && req.Wd != "" {
		opt.QueryLike["title"] = []interface{}{req.Wd}
		opt.QueryLike["keywords"] = []interface{}{req.Wd}
		opt.QueryLike["description"] = []interface{}{req.Wd}
		opt.QueryLike["content"] = []interface{}{req.Wd}
	}

	articles, total, err := s.dbModel.GetArticleList(opt)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error("ListArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "获取文章列表失败")
	}

	var pbArticle []*pb.Article
	err = util.CopyStruct(articles, &pbArticle)
	if err != nil {
		s.logger.Error("ListArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "获取文章列表失败"+err.Error())
	}

	return &pb.ListArticleReply{
		Total:   total,
		Article: pbArticle,
	}, nil
}

func (s *ArticleAPIService) RestoreRecycleArticle(ctx context.Context, req *pb.RestoreArticleRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.RestoreArticle(req.Id)
	if err != nil {
		s.logger.Error("RestoreRecycleArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "恢复文章失败："+err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *ArticleAPIService) DeleteRecycleArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteArticle(req.Id, true)
	if err != nil {
		s.logger.Error("DeleteRecycleArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "删除文章失败")
	}

	return &emptypb.Empty{}, nil
}

func (s *ArticleAPIService) RecommendArticles(ctx context.Context, req *pb.RecommendArticlesRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.RecommendArticles(req.ArticleId, req.IsRecommend)
	if err != nil {
		s.logger.Error("RecommendArticles", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "推荐文章失败")
	}

	return &emptypb.Empty{}, nil
}

func (s *ArticleAPIService) EmptyRecycleArticle(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	opt := &model.OptionGetArticleList{
		IsRecycle:    true,
		SelectFields: []string{"id"},
	}

	for {
		articles, _, err := s.dbModel.GetArticleList(opt)
		if err != nil && err != gorm.ErrRecordNotFound {
			s.logger.Error("EmptyRecycleArticle", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "清空回收站失败")
		}

		var ids []int64
		for _, article := range articles {
			ids = append(ids, article.Id)
		}

		if len(ids) > 0 {
			err = s.dbModel.DeleteArticle(ids, true)
			if err != nil {
				s.logger.Error("EmptyRecycleArticle", zap.Error(err))
				return nil, status.Errorf(codes.Internal, "清空回收站失败")
			}
		}

		if len(articles) == 0 {
			break
		}
	}

	return &emptypb.Empty{}, nil
}

// 批量审核文档
func (s *ArticleAPIService) CheckArticles(ctx context.Context, req *pb.CheckArticlesRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.CheckArticles(req.ArticleId, req.Status, req.RejeactReason)
	if err != nil {
		s.logger.Error("CheckArticles", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "审核文章失败")
	}
	return &emptypb.Empty{}, nil
}

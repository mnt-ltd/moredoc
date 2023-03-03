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
func (s *ArticleAPIService) CreateArticle(ctx context.Context, req *pb.Article) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
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

	err = s.dbModel.CreateArticle(article)
	if err != nil {
		s.logger.Error("CreateArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "创建文章失败")
	}

	return &emptypb.Empty{}, nil
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
	return pbArticle, nil
}

// ListArticles 获取文章列表。
// 如果是有权限，则可以根据关键字来查询，否则只能简单查询
func (s *ArticleAPIService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	opt := &model.OptionGetArticleList{
		Page:      int(req.Page),
		Size:      int(req.Size_),
		WithCount: true,
		QueryLike: make(map[string][]interface{}),
		Sort:      []string{req.Order},
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

package biz

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"
	"moredoc/util/segword/jieba"

	"github.com/araddon/dateparse"
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

func (s *ArticleAPIService) checkLogin(ctx context.Context) (userClaims *auth.UserClaims, err error) {
	return checkGRPCLogin(s.dbModel, ctx)
}

// CreateArticle 创建文章
func (s *ArticleAPIService) CreateArticle(ctx context.Context, req *pb.Article) (*pb.Article, error) {
	userClaims, err := s.checkPermission(ctx)
	if userClaims == nil { // 未登录
		return nil, err
	}

	isAdmin := err == nil // err为nil表示管理员，否则为普通用户
	if !(isAdmin || s.dbModel.CanIAccessPublishArticle(userClaims.UserId)) {
		return nil, status.Errorf(codes.PermissionDenied, "您没有权限发布文章")
	}

	if !isAdmin && req.Identifier != "" { // 只有管理员才可以设置uuid
		req.Identifier = ""
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
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// 管理员在创建文章的时候，可以直接设置为推荐文章
	if isAdmin && req.IsRecommend && req.RecommendAt == nil {
		now := time.Now()
		article.RecommendAt = &now
	}

	article.UserId = userClaims.UserId
	if isAdmin {
		article.Status = model.ArticleStatusPass
	} else {
		article.Status = s.dbModel.GetDefaultArticleStatus(userClaims.UserId)
	}

	s.logger.Debug("CreateArticle", zap.Any("article", article))
	s.fixKeywordsAndDescription(article)
	err = s.dbModel.CreateArticle(article)
	if err != nil {
		s.logger.Error("CreateArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
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
	userClaims, err := s.checkPermission(ctx)
	if userClaims == nil {
		return nil, err
	}
	isAdmin := err == nil
	article := &model.Article{}
	err = util.CopyStruct(req, article)
	if err != nil {
		s.logger.Error("UpdateArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	existArticle, _ := s.dbModel.GetArticle(article.Id)
	if !isAdmin && existArticle.UserId != userClaims.UserId {
		return nil, status.Errorf(codes.PermissionDenied, "您没有权限修改此文章")
	}

	article.RecommendAt = existArticle.RecommendAt
	if isAdmin {
		if req.IsRecommend && existArticle.RecommendAt == nil {
			now := time.Now()
			article.RecommendAt = &now
		}

		// 如果文章之前被拒绝，则修改后状态变为待审核
		if existArticle.Status == model.ArticleStatusReject {
			article.Status = model.ArticleStatusPending
		} else {
			article.Status = existArticle.Status
		}

	} else {
		// 如果文章之前被拒绝，则修改后状态变为待审核
		if existArticle.Status == model.ArticleStatusReject {
			article.Status = model.ArticleStatusPending
		} else {
			article.Status = existArticle.Status
		}
	}

	s.fixKeywordsAndDescription(article)
	err = s.dbModel.UpdateArticle(article)
	if err != nil {
		s.logger.Error("UpdateArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// DeleteArticle 删除文章
func (s *ArticleAPIService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*emptypb.Empty, error) {
	userClaims, err := s.checkPermission(ctx)
	if userClaims == nil {
		return nil, err
	}

	isAdmin := err == nil
	ids := req.Id
	if len(ids) == 0 {
		return &emptypb.Empty{}, nil
	}

	if !isAdmin { // 非管理员，只能删除自己的文章
		articles, _, _ := s.dbModel.GetArticleList(&model.OptionGetArticleList{
			Ids:          util.Slice2Interface(req.Id),
			WithCount:    false,
			SelectFields: []string{"id"},
			QueryIn: map[string][]interface{}{
				"user_id": {userClaims.UserId},
			},
		})
		if len(articles) == 0 {
			return nil, errors.New("您没有权限删除指定文章")
		}
		for _, article := range articles {
			ids = append(ids, article.Id)
		}
	}

	err = s.dbModel.DeleteArticle(ids)
	if err != nil {
		s.logger.Error("DeleteArticle", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "删除文章失败:"+err.Error())
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
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		article.ViewCount += 1
	}

	if article.Id == 0 {
		return nil, status.Errorf(codes.NotFound, "文章不存在")
	}

	userClaims, err := s.checkPermission(ctx)
	s.logger.Debug("GetArticle", zap.Any("userClaims", userClaims), zap.Any("article", article))
	isAdmin := err == nil
	// 管理员或者是作者，可以查看未通过的文章
	if article.Status != model.ArticleStatusPass && !(isAdmin || userClaims.UserId == article.UserId) {
		err = errors.New("文章不存在")
		return nil, err
	}

	s.dbModel.UpdateArticleViewCount(article.Id, article.ViewCount)
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

	userClaims, err := s.checkPermission(ctx)
	isAdmin := err == nil
	if isAdmin || (userClaims != nil && len(req.UserId) > 0 && userClaims.UserId == req.UserId[0]) {
		// 管理员或者是作者，可以查询相关状态的文档
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

	if len(req.UserId) > 0 {
		opt.QueryIn["user_id"] = util.Slice2Interface(req.UserId)
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

// 相关文章
func (s *ArticleAPIService) GetRelatedArticles(ctx context.Context, req *pb.GetArticleRequest) (*pb.ListArticleReply, error) {
	if req.Identifier == "" {
		return nil, status.Errorf(codes.InvalidArgument, "参数错误:文章标识为空")
	}
	articles, _ := s.dbModel.GetRelatedArticles(req.Identifier)
	if len(articles) == 0 {
		return nil, status.Errorf(codes.NotFound, "相关文章不存在")
	}

	var pbArticle []*pb.Article
	err := util.CopyStruct(articles, &pbArticle)
	if err != nil {
		s.logger.Error("GetRelatedArticles", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "获取相关文章失败")
	}
	return &pb.ListArticleReply{Article: pbArticle}, nil
}

// 搜索文章
func (s *ArticleAPIService) SearchArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.SearchArticleReply, error) {
	res := &pb.SearchArticleReply{}
	now := time.Now()
	opt := &model.OptionGetArticleList{
		WithCount:  true,
		Page:       int(req.Page),
		Size:       int(req.Size_),
		QueryIn:    make(map[string][]interface{}),
		QueryRange: make(map[string][2]interface{}),
	}

	opt.Size = util.LimitRange(opt.Size, 10, 10)
	opt.Page = util.LimitRange(opt.Page, 1, 10000) // 最大默认1w页，等同于不限制页数
	maxPages := s.dbModel.GetConfigOfDisplay(model.ConfigDisplayMaxSearchPages).MaxSearchPages
	if maxPages > 0 {
		opt.Page = util.LimitRange(opt.Page, 1, int(maxPages))
	}
	if req.Wd == "" {
		return res, nil
	}
	opt.QueryLike = map[string][]interface{}{
		"title":       util.Slice2Interface(strings.Split(req.Wd, " ")),
		"keywords":    util.Slice2Interface(strings.Split(req.Wd, " ")),
		"description": util.Slice2Interface(strings.Split(req.Wd, " ")),
	}

	if len(req.CategoryId) > 0 {
		opt.QueryIn = map[string][]interface{}{
			"category_id": util.Slice2Interface(req.CategoryId),
		}
	}

	if l := len(req.CreatedAt); l > 0 {
		end := time.Now()
		start, _ := dateparse.ParseLocal(req.CreatedAt[0])
		if l > 1 {
			end, _ = dateparse.ParseLocal(req.CreatedAt[1])
		}
		opt.QueryRange["created_at"] = [2]interface{}{start, end}
	}

	if len(req.UserId) > 0 {
		opt.QueryIn["user_id"] = util.Slice2Interface(req.UserId)
	}

	if req.Sort != "" {
		if req.Sort == "latest" {
			opt.Sort = []string{"id"}
		} else {
			opt.Sort = []string{req.Sort}
		}
	}

	articles, total, err := s.dbModel.GetArticleList(opt)
	if err != nil {
		return res, status.Errorf(codes.Internal, "搜索文章失败：%s", err)
	}
	util.CopyStruct(&articles, &res.Article)

	if maxPages > 0 && total > int64(maxPages)*int64(opt.Size) {
		total = int64(maxPages) * int64(opt.Size)
	}

	res.Total = total
	spendTime := time.Since(now).Seconds()
	res.Spend = fmt.Sprintf("%.3f", spendTime)

	retentionDays := s.dbModel.GetConfigOfSecurity(model.ConfigSecuritySearchRecordRetentionDays).SearchRecordRetentionDays
	if retentionDays > 0 {
		var userId int64
		userCliaims, _ := s.checkLogin(ctx)
		if userCliaims != nil {
			userId = userCliaims.UserId
		}
		s.dbModel.CreateSearchRecord(&model.SearchRecord{
			Ip:        util.GetGRPCRemoteIP(ctx),
			Total:     int(total),
			Page:      int(opt.Page),
			UserAgent: util.GetGRPCUserAgent(ctx),
			UserId:    userId,
			SpendTime: spendTime,
			Keywords:  req.Wd,
			Type:      1, // 搜文章
		})
	}
	return res, nil
}

func (s *ArticleAPIService) fixKeywordsAndDescription(article *model.Article) {
	if strings.TrimSpace(article.Description) == "" {
		text := strings.TrimSpace(util.GetTextFromHTML(article.Content))
		article.Description = util.Substr(text, 200)
	}

	if strings.TrimSpace(article.Keywords) == "" {
		article.Keywords = strings.Join(jieba.SegWords(article.Title), ", ")
	}
}

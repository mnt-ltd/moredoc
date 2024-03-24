package model

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	// 审核状态：0，待审核; 1, 审核通过; 2, 审核拒绝
	ArticleStatusPending = 0
	ArticleStatusPass    = 1
	ArticleStatusReject  = 2
)

type Article struct {
	Id            int64          `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Identifier    string         `form:"identifier" json:"identifier,omitempty" gorm:"column:identifier;type:varchar(64);size:64;index:identifier,unique;comment:文章标识，唯一;"`
	UserId        int64          `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint;comment:用户ID;index:user_id"`
	ViewCount     int            `form:"view_count" json:"view_count,omitempty" gorm:"column:view_count;type:int(11);size:11;default:0;comment:阅读;"`
	FavoriteCount int            `form:"favorite_count" json:"favorite_count,omitempty" gorm:"column:favorite_count;type:int(11);size:11;default:0;comment:收藏;"`
	CommentCount  int            `form:"comment_count" json:"comment_count,omitempty" gorm:"column:comment_count;type:int(11);size:11;default:0;comment:评论;"`
	Title         string         `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(255);size:255;comment:文章标题;"`
	Keywords      string         `form:"keywords" json:"keywords,omitempty" gorm:"column:keywords;type:varchar(255);size:255;comment:关键字;"`
	Description   string         `form:"description" json:"description,omitempty" gorm:"column:description;type:varchar(255);size:255;comment:摘要;"`
	Content       string         `form:"content" json:"content,omitempty" gorm:"column:content;type:longtext;comment:内容;"`
	CreatedAt     time.Time      `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt     time.Time      `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
	DeletedAt     gorm.DeletedAt `form:"deleted_at" json:"deleted_at,omitempty" gorm:"column:deleted_at;type:datetime;comment:删除时间;index:idx_deleted_at"`
	RecommendAt   *time.Time     `form:"recommend_at" json:"recommend_at,omitempty" gorm:"column:recommend_at;type:datetime;comment:推荐时间;index:idx_recommend_at;default:null;"`
	CategoryId    []int64        `form:"category_id" json:"category_id,omitempty" gorm:"-"`
	Status        int32          `form:"status" json:"status,omitempty" gorm:"column:status;type:int(11);size:11;default:0;comment:状态;index:idx_status"`
	RejectReason  string         `form:"reject_reason" json:"reject_reason,omitempty" gorm:"column:reject_reason;type:varchar(2048);size:2048;comment:审核拒绝信息;"`
}

func (Article) TableName() string {
	return tablePrefix + "article"
}

func (m *DBModel) initArticle() (err error) {
	articles := []Article{
		{
			Identifier: "about",
			Title:      "关于我们",
			Content:    "请输入【关于我们】的内容",
		},
		{
			Identifier: "agreement",
			Title:      "文库协议",
			Content:    "请输入【文库协议】的内容",
		},
		{
			Identifier: "contact",
			Title:      "联系我们",
			Content:    "请输入【联系我们】的内容",
		},
		{
			Identifier: "feedback",
			Title:      "意见反馈",
			Content:    "请输入【意见反馈】的内容",
		},
		{
			Identifier: "copyright",
			Title:      "免责声明",
			Content:    "请输入【免责声明】的内容",
		},
		{
			Identifier: "help",
			Title:      "使用帮助",
			Content:    "请输入【使用帮助】的内容",
		},
	}
	for _, article := range articles {
		exist, _ := m.GetArticleByIdentifier(article.Identifier, "id")
		if exist.Id == 0 {
			err = m.CreateArticle(&article)
			if err != nil {
				m.logger.Error("initArticle", zap.Error(err), zap.Any("article", article))
				return
			}
		}
	}
	return
}

// CreateArticle 创建Article
func (m *DBModel) CreateArticle(article *Article) (err error) {
	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Create(article).Error
	if err != nil {
		m.logger.Error("CreateArticle", zap.Error(err))
		return
	}

	if len(article.CategoryId) > 0 {
		err = tx.Model(&Category{}).Where("id in ?", article.CategoryId).Update("doc_count", gorm.Expr("doc_count + 1")).Error
		if err != nil {
			m.logger.Error("CreateArticle", zap.Error(err))
			return
		}

		// 增加文档与分类的关联
		var (
			articleCategories []ArticleCategory
			existMap          = make(map[int64]bool)
		)

		for _, categoryId := range article.CategoryId {
			if existMap[categoryId] {
				continue
			}
			articleCategories = append(articleCategories, ArticleCategory{
				ArticleId:  article.Id,
				CategoryId: categoryId,
			})
			existMap[categoryId] = true
		}

		err = tx.Create(&articleCategories).Error
		if err != nil {
			m.logger.Error("CreateArticle", zap.Error(err))
			return
		}
	}

	m.checkArticleFile(article)
	return
}

// UpdateArticle 更新Article，如果需要更新指定字段，则请指定updateFields参数
// 注意：不支持更新identifier
func (m *DBModel) UpdateArticle(article *Article, updateFields ...string) (err error) {
	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var (
		existArticleCategories []ArticleCategory
		existCategoryIds       []int64
	)
	tx.Model(&ArticleCategory{}).Where("article_id = ?", article.Id).Find(&existArticleCategories)
	for _, cate := range existArticleCategories {
		existCategoryIds = append(existCategoryIds, cate.CategoryId)
	}

	if len(existCategoryIds) > 0 {
		err = tx.Model(&Category{}).Where("id in ?", existCategoryIds).Update("doc_count", gorm.Expr("doc_count - 1")).Error
		if err != nil {
			m.logger.Error("UpdateArticle", zap.Error(err))
			return
		}

		err = tx.Where("article_id = ?", article.Id).Delete(&ArticleCategory{}).Error
		if err != nil {
			m.logger.Error("UpdateArticle", zap.Error(err))
			return
		}
	}

	if len(article.CategoryId) > 0 {
		err = tx.Model(&Category{}).Where("id in ?", article.CategoryId).Update("doc_count", gorm.Expr("doc_count + 1")).Error
		if err != nil {
			m.logger.Error("UpdateArticle", zap.Error(err))
			return
		}

		// 增加文档与分类的关联
		var (
			articleCategories []ArticleCategory
			exist             = make(map[int64]bool)
		)
		for _, categoryId := range article.CategoryId {
			if exist[categoryId] {
				continue
			}
			exist[categoryId] = true
			articleCategories = append(articleCategories, ArticleCategory{
				ArticleId:  article.Id,
				CategoryId: categoryId,
			})
		}

		err = tx.Create(&articleCategories).Error
		if err != nil {
			m.logger.Error("UpdateArticle", zap.Error(err))
			return
		}
	}

	tableName := Article{}.TableName()
	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) == 0 { // 更新全部字段，包括零值字段
		updateFields = m.GetTableFields(tableName)
	}
	ignoreFields := []string{"identifier", "view_count", "favorite_count", "comment_count", "user_id"}
	err = tx.Model(article).Select(updateFields).Where("id = ?", article.Id).Omit(ignoreFields...).Updates(article).Error
	if err != nil {
		m.logger.Error("UpdateArticle", zap.Error(err))
		return
	}

	m.checkArticleFile(article)
	return
}

// UpdateArticleViewCount 更新浏览量
func (m *DBModel) UpdateArticleViewCount(id int64, viewCount int) (err error) {
	sql := fmt.Sprintf("update %s set view_count=? where id=?", Article{}.TableName())
	err = m.db.Exec(sql, viewCount, id).Error
	if err != nil {
		m.logger.Error("UpdateArticleViewCount", zap.Error(err))
	}
	return
}

// GetArticle 根据id获取Article
func (m *DBModel) GetArticle(id interface{}, fields ...string) (article Article, err error) {
	db := m.db

	fields = m.FilterValidFields(Article{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetArticle", zap.Error(err))
		return
	}
	cates, _ := m.GetArticleCategories(article.Id)
	article.CategoryId = make([]int64, 0, len(cates))
	for _, cate := range cates {
		article.CategoryId = append(article.CategoryId, cate.CategoryId)
	}
	return
}

// GetArticleByIdentifier(identifier string, fields ...string) 根据唯一索引获取Article
func (m *DBModel) GetArticleByIdentifier(identifier string, fields ...string) (article Article, err error) {
	db := m.db

	fields = m.FilterValidFields(Article{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	db = db.Where("identifier = ?", identifier)

	err = db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetArticleByIdentifier", zap.Error(err))
		return
	}

	cates, _ := m.GetArticleCategories(article.Id)
	article.CategoryId = make([]int64, 0, len(cates))
	for _, cate := range cates {
		article.CategoryId = append(article.CategoryId, cate.CategoryId)
	}

	return
}

type OptionGetArticleList struct {
	Page         int
	Size         int
	WithCount    bool                      // 是否返回总数
	Ids          []interface{}             // id列表
	SelectFields []string                  // 查询字段
	QueryRange   map[string][2]interface{} // map[field][]{min,max}
	QueryIn      map[string][]interface{}  // map[field][]{value1,value2,...}
	QueryLike    map[string][]interface{}  // map[field][]{value1,value2,...}
	Sort         []string
	IsRecycle    bool   // 是否是回收站模式查询
	IsRecommend  []bool // 是否是推荐模式查询
}

// GetArticleList 获取Article列表
func (m *DBModel) GetArticleList(opt *OptionGetArticleList) (articleList []Article, total int64, err error) {
	tableName := Article{}.TableName() + " a"
	db := m.db.Table(tableName).Unscoped()
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("a.id in (?)", opt.Ids)
	}

	if categoryIds, ok := opt.QueryIn["category_id"]; ok && len(categoryIds) > 0 {
		tableCategory := ArticleCategory{}.TableName()
		db = db.Joins("left join "+tableCategory+" ac on ac.article_id = a.id").Where("ac.category_id in (?)", categoryIds)
	}

	if opt.IsRecycle {
		db = db.Where("a.deleted_at is not null")
		// 回收站模式下，按删除时间倒序
		opt.Sort = []string{"a.deleted_at desc"}
	} else {
		db = db.Where("a.deleted_at is null")
	}

	if len(opt.IsRecommend) == 1 {
		if opt.IsRecommend[0] {
			db = db.Where("a.recommend_at is not null")
		} else {
			db = db.Where("a.recommend_at is null")
		}
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetArticleList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	} else {
		db = db.Select(m.GetTableFields(tableName, "a.content"))
	}

	if len(opt.Sort) > 0 {
		db = m.generateQuerySort(db, tableName, opt.Sort)
	} else {
		db = db.Order("a.id desc")
	}

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetArticleList", zap.Error(err))
		return
	}

	if len(articleList) == 0 {
		return
	}

	var (
		articleIds        []int64
		articleIdMapIndex = make(map[int64]int)
	)

	for index, article := range articleList {
		articleIds = append(articleIds, article.Id)
		articleIdMapIndex[article.Id] = index
	}

	articleCategories, _ := m.GetArticleCategories(articleIds...)
	for _, articleCategory := range articleCategories {
		if index, ok := articleIdMapIndex[articleCategory.ArticleId]; ok {
			articleList[index].CategoryId = append(articleList[index].CategoryId, articleCategory.CategoryId)
		}
	}

	return
}

// DeleteArticle 删除数据
func (m *DBModel) DeleteArticle(ids []int64, deepDelete ...bool) (err error) {
	if len(ids) == 0 {
		return
	}

	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	deep := false
	if len(deepDelete) > 0 {
		deep = deepDelete[0]
	}

	if !deep {
		// 软删除：删除文章、文章关联分类数量-1
		err = tx.Where("id in (?)", ids).Delete(&Article{}).Error
		if err != nil {
			m.logger.Error("DeleteArticle", zap.Error(err))
			return
		}

		// 查找文章已存在的分类，减少分类的文档数量
		var (
			articleCategories      []ArticleCategory
			articleIdMapCategories = make(map[int64][]int64)
		)

		tx.Where("article_id in (?)", ids).Find(&articleCategories)
		for _, cate := range articleCategories {
			articleIdMapCategories[cate.ArticleId] = append(articleIdMapCategories[cate.ArticleId], cate.CategoryId)
		}

		for _, cateIds := range articleIdMapCategories {
			err = tx.Model(&Category{}).Where("id in (?)", cateIds).Update("doc_count", gorm.Expr("doc_count - 1")).Error
			if err != nil {
				m.logger.Error("DeleteArticle", zap.Error(err))
				return
			}
		}
		return
	}

	// 文章删除
	err = tx.Unscoped().Where("id in (?)", ids).Delete(&Article{}).Error
	if err != nil {
		m.logger.Error("DeleteArticle", zap.Error(err))
		return
	}

	// 文章分类关联删除
	err = tx.Where("article_id in (?)", ids).Delete(&ArticleCategory{}).Error
	if err != nil {
		m.logger.Error("DeleteArticle", zap.Error(err))
		return
	}

	// 附件标记删除
	err = tx.Where("type = ? and type_id in (?)", AttachmentTypeArticle, ids).Delete(&Attachment{}).Error
	if err != nil {
		m.logger.Error("DeleteArticle", zap.Error(err))
		return
	}
	return
}

func (m *DBModel) SetArticlesCategory(articleId, categoryId []int64) (err error) {
	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	for _, id := range articleId {
		// 1. 旧的文档分类，减少计数
		var articleCates []ArticleCategory
		m.db.Model(&ArticleCategory{}).Where("article_id = ?", id).Find(&articleCates)
		for _, cate := range articleCates {
			err = tx.Model(&Category{}).Where("id = ?", cate.CategoryId).Update("doc_count", gorm.Expr("doc_count - ?", 1)).Error
			if err != nil {
				m.logger.Error("SetArticlesCategory", zap.Error(err))
				return
			}
		}

		// 2. 删除旧的分类
		err = tx.Model(&ArticleCategory{}).Where("article_id = ?", id).Delete(&DocumentCategory{}).Error
		if err != nil {
			m.logger.Error("SetArticlesCategory", zap.Error(err))
			return
		}

		// 3. 添加新的分类
		articleCates = []ArticleCategory{}
		for _, cid := range categoryId {
			articleCates = append(articleCates, ArticleCategory{
				ArticleId:  id,
				CategoryId: cid,
			})
		}
		err = tx.Create(&articleCates).Error
		if err != nil {
			m.logger.Error("SetArticlesCategory", zap.Error(err))
			return
		}

		// 4. 更新文档分类统计
		err = tx.Model(&Category{}).Where("id in (?)", categoryId).Update("doc_count", gorm.Expr("doc_count + ?", 1)).Error
		if err != nil {
			m.logger.Error("SetArticlesCategory", zap.Error(err))
			return
		}
	}
	return
}

// checkArticleFile 检查文章中的文件，包括音频视频和图片等
func (m *DBModel) checkArticleFile(article *Article) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(article.Content))
	if err != nil {
		m.logger.Error("checkArticleFile", zap.Error(err))
		return
	}

	var (
		hashes []string
		tags   = []string{"img", "video", "audio"}
	)

	for _, tag := range tags {
		doc.Find(tag).Each(func(i int, selection *goquery.Selection) {
			src, ok := selection.Attr("src")
			if !ok {
				src, ok = selection.Find("source").Attr("src")
			}
			if ok && strings.HasPrefix(src, "/uploads/") {
				src = strings.Split(src, "?")[0]
				hashes = append(hashes, strings.TrimSuffix(filepath.Base(src), filepath.Ext(src)))
			}
		})
	}

	if len(hashes) > 0 { // 更新内容ID
		err = m.db.Model(&Attachment{}).Where("hash in (?) and type = ? and type_id = 0", hashes, AttachmentTypeArticle).Update("type_id", article.Id).Error
		if err != nil {
			m.logger.Error("checkArticleFile", zap.Error(err))
		}
	}
}

func (m *DBModel) CountArticle() (count int64, err error) {
	err = m.db.Model(&Article{}).Count(&count).Error
	return
}

// 从回收站中恢复选中的文章
func (m *DBModel) RestoreArticle(ids []int64) (err error) {
	if len(ids) == 0 {
		return
	}

	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Model(&Article{}).Unscoped().Where("id in (?)", ids).Update("deleted_at", nil).Error
	if err != nil {
		m.logger.Error("RestoreArticle", zap.Error(err))
		return
	}

	// 查找文章已存在的分类，增加分类的文档数量
	var (
		articleCategories      []ArticleCategory
		articleIdMapCategories = make(map[int64][]int64)
	)

	tx.Where("article_id in (?)", ids).Find(&articleCategories)
	for _, cate := range articleCategories {
		articleIdMapCategories[cate.ArticleId] = append(articleIdMapCategories[cate.ArticleId], cate.CategoryId)
	}

	for _, cateIds := range articleIdMapCategories {
		err = tx.Model(&Category{}).Where("id in (?)", cateIds).Update("doc_count", gorm.Expr("doc_count + 1")).Error
		if err != nil {
			m.logger.Error("RestoreArticle", zap.Error(err))
			return
		}
	}

	return
}

func (m *DBModel) RecommendArticles(articleIds []int64, isRecommend bool) (err error) {
	var val interface{} = time.Now()
	if !isRecommend {
		val = nil
	}
	err = m.db.Model(&Article{}).Where("id in (?)", articleIds).Update("recommend_at", val).Error
	if err != nil {
		m.logger.Error("RecommendArticles", zap.Error(err))
	}
	return
}

func (m *DBModel) CheckArticles(ids []int64, status int32, reason ...string) (err error) {
	if len(ids) == 0 {
		return
	}
	r := ""
	if len(reason) > 0 {
		r = reason[0]
	}
	err = m.db.Model(&Article{}).Where("id in (?)", ids).Updates(map[string]interface{}{
		"status":        status,
		"reject_reason": r,
	}).Error
	if err != nil {
		m.logger.Error("CheckArticles", zap.Error(err))
	}
	return
}

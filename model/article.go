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

type Article struct {
	Id          int64     `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Identifier  string    `form:"identifier" json:"identifier,omitempty" gorm:"column:identifier;type:varchar(64);size:64;index:identifier,unique;comment:文章标识，唯一;"`
	Author      string    `form:"author" json:"author,omitempty" gorm:"column:author;type:varchar(64);size:64;comment:作者;"`
	ViewCount   int       `form:"view_count" json:"view_count,omitempty" gorm:"column:view_count;type:int(11);size:11;default:0;comment:阅读;"`
	Title       string    `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(255);size:255;comment:文章标题;"`
	Keywords    string    `form:"keywords" json:"keywords,omitempty" gorm:"column:keywords;type:varchar(255);size:255;comment:关键字;"`
	Description string    `form:"description" json:"description,omitempty" gorm:"column:description;type:varchar(255);size:255;comment:摘要;"`
	Content     string    `form:"content" json:"content,omitempty" gorm:"column:content;type:longtext;comment:内容;"`
	CreatedAt   time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Article) TableName() string {
	return tablePrefix + "article"
}

func (m *DBModel) initArticle() (err error) {
	// 初始化文章:
	// about 关于我们
	// agreement 文库协议
	// contact 联系我们
	// feedback 意见反馈
	// copyright 免责声明
	// help 使用帮助
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
	err = m.db.Create(article).Error
	if err != nil {
		m.logger.Error("CreateArticle", zap.Error(err))
		return
	}
	m.checkArticleFile(article)
	return
}

// UpdateArticle 更新Article，如果需要更新指定字段，则请指定updateFields参数
// 注意：不支持更新identifier
func (m *DBModel) UpdateArticle(article *Article, updateFields ...string) (err error) {
	db := m.db.Model(article)
	tableName := Article{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", article.Id).Omit("identifier").Updates(article).Error
	if err != nil {
		m.logger.Error("UpdateArticle", zap.Error(err))
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
}

// GetArticleList 获取Article列表
func (m *DBModel) GetArticleList(opt *OptionGetArticleList) (articleList []Article, total int64, err error) {
	tableName := Article{}.TableName()
	db := m.db.Model(&Article{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
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
		db = db.Omit("content")
	}

	if len(opt.Sort) > 0 {
		db = m.generateQuerySort(db, tableName, opt.Sort)
	} else {
		db = db.Order("id desc")
	}

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetArticleList", zap.Error(err))
	}
	return
}

// DeleteArticle 删除数据
func (m *DBModel) DeleteArticle(ids []int64) (err error) {
	sess := m.db.Begin()
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()

	err = sess.Where("id in (?)", ids).Delete(&Article{}).Error
	if err != nil {
		m.logger.Error("DeleteArticle", zap.Error(err))
		return
	}

	err = sess.Where("type = ? and type_id in (?)", AttachmentTypeArticle, ids).Delete(&Attachment{}).Error
	if err != nil {
		m.logger.Error("DeleteArticle", zap.Error(err))
		return
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

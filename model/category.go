package model

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Category struct {
	Id        int       `form:"id" json:"id,omitempty" gorm:"column:id;type:int(11);size:11;default:0;primarykey;comment:;"`
	ParentId  int       `form:"parent_id" json:"parent_id,omitempty" gorm:"column:parent_id;type:int(11);size:11;default:0;index:parent_id_title,unique;index:parent_id;comment:上级ID;"`
	Title     string    `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(64);size:64;default:;index:parent_id_title,unique;comment:分类名称;"`
	Cover     string    `form:"cover" json:"cover,omitempty" gorm:"column:cover;type:varchar(255);size:255;default:;comment:分类封面;"`
	DocCount  int       `form:"doc_count" json:"doc_count,omitempty" gorm:"column:doc_count;type:int(11);size:11;default:0;comment:文档统计;"`
	Sort      int       `form:"sort" json:"sort,omitempty" gorm:"column:sort;type:int(11);size:11;default:0;comment:排序，值越大越靠前;"`
	Alias     string    `form:"alias" json:"alias,omitempty" gorm:"column:alias;type:varchar(64);size:64;default:;comment:别名，限英文和数字等组成;"`
	Status    int8      `form:"status" json:"status,omitempty" gorm:"column:status;type:tinyint(1);size:1;default:0;comment:状态：0 正常，1 禁用;"`
	CreatedAt time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;default:;comment:创建时间;"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;default:;comment:更新时间;"`
}

func (Category) TableName() string {
	return tablePrefix + "category"
}

// CreateCategory 创建Category
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreateCategory(category *Category) (err error) {
	err = m.db.Create(category).Error
	if err != nil {
		m.logger.Error("CreateCategory", zap.Error(err))
		return
	}
	return
}

// UpdateCategory 更新Category，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateCategory(category *Category, updateFields ...string) (err error) {
	db := m.db.Model(category)

	updateFields = m.FilterValidFields(Category{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	}

	err = db.Where("id = ?", category.Id).Updates(category).Error
	if err != nil {
		m.logger.Error("UpdateCategory", zap.Error(err))
	}
	return
}

// GetCategory 根据id获取Category
func (m *DBModel) GetCategory(id interface{}, fields ...string) (category Category, err error) {
	db := m.db

	fields = m.FilterValidFields(Category{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&category).Error
	return
}

// GetCategoryByParentIdTitle(parentId int, title string, fields ...string) 根据唯一索引获取Category
func (m *DBModel) GetCategoryByParentIdTitle(parentId int, title string, fields ...string) (category Category, err error) {
	db := m.db

	fields = m.FilterValidFields(Category{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	db = db.Where("parent_id = ?", parentId)

	db = db.Where("title = ?", title)

	err = db.First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetCategoryByParentIdTitle", zap.Error(err))
		return
	}
	return
}

type OptionGetCategoryList struct {
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

// GetCategoryList 获取Category列表
func (m *DBModel) GetCategoryList(opt OptionGetCategoryList) (categoryList []Category, total int64, err error) {
	db := m.db.Model(&Category{})

	for field, rangeValue := range opt.QueryRange {
		fields := m.FilterValidFields(Category{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		if rangeValue[0] != nil {
			db = db.Where(fmt.Sprintf("%s >= ?", field), rangeValue[0])
		}
		if rangeValue[1] != nil {
			db = db.Where(fmt.Sprintf("%s <= ?", field), rangeValue[1])
		}
	}

	for field, values := range opt.QueryIn {
		fields := m.FilterValidFields(Category{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(fmt.Sprintf("%s in (?)", field), values)
	}

	for field, values := range opt.QueryLike {
		fields := m.FilterValidFields(Category{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(strings.TrimSuffix(fmt.Sprintf(strings.Join(make([]string, len(values)+1), "%s like ? or"), field), "or"), values...)
	}

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetCategoryList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(Category{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) > 0 {
		var sorts []string
		for _, sort := range opt.Sort {
			slice := strings.Split(sort, " ")
			if len(m.FilterValidFields(Category{}.TableName(), slice[0])) == 0 {
				continue
			}

			if len(slice) == 2 {
				sorts = append(sorts, fmt.Sprintf("%s %s", slice[0], slice[1]))
			} else {
				sorts = append(sorts, fmt.Sprintf("%s desc", slice[0]))
			}
		}
		if len(sorts) > 0 {
			db = db.Order(strings.Join(sorts, ","))
		}
	}

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&categoryList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetCategoryList", zap.Error(err))
	}
	return
}

// DeleteCategory 删除数据
// TODO: 删除数据之后，存在 category_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeleteCategory(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Category{}).Error
	if err != nil {
		m.logger.Error("DeleteCategory", zap.Error(err))
	}
	return
}

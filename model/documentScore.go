package model

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DocumentScore struct {
	Id         int64     `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	DocumentId int64     `form:"document_id" json:"document_id,omitempty" gorm:"column:document_id;type:bigint(20);size:20;default:0;comment:文档ID;"`
	UserId     int64     `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;comment:用户ID;"`
	Score      int       `form:"score" json:"score,omitempty" gorm:"column:score;type:int(11);size:11;default:0;comment:文档评分值，3位数，如500表示5分;"`
	CreatedAt  time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt  time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

// 这里是proto文件中的结构体，可以根据需要删除或者调整
//message DocumentScore {
// int64 id = 1;
// int64 document_id = 2;
// int64 user_id = 3;
// int32 score = 4;
// google.protobuf.Timestamp created_at = 5 [ (gogoproto.stdtime) = true ];
// google.protobuf.Timestamp updated_at = 6 [ (gogoproto.stdtime) = true ];
//}

func (DocumentScore) TableName() string {
	return tablePrefix + "document_score"
}

// CreateDocumentScore 创建DocumentScore
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreateDocumentScore(documentScore *DocumentScore) (err error) {
	err = m.db.Create(documentScore).Error
	if err != nil {
		m.logger.Error("CreateDocumentScore", zap.Error(err))
		return
	}
	return
}

// UpdateDocumentScore 更新DocumentScore，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateDocumentScore(documentScore *DocumentScore, updateFields ...string) (err error) {
	db := m.db.Model(documentScore)

	updateFields = m.FilterValidFields(DocumentScore{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	}

	err = db.Where("id = ?", documentScore.Id).Updates(documentScore).Error
	if err != nil {
		m.logger.Error("UpdateDocumentScore", zap.Error(err))
	}
	return
}

// GetDocumentScore 根据id获取DocumentScore
func (m *DBModel) GetDocumentScore(id interface{}, fields ...string) (documentScore DocumentScore, err error) {
	db := m.db

	fields = m.FilterValidFields(DocumentScore{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&documentScore).Error
	return
}

type OptionGetDocumentScoreList struct {
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

// GetDocumentScoreList 获取DocumentScore列表
func (m *DBModel) GetDocumentScoreList(opt OptionGetDocumentScoreList) (documentScoreList []DocumentScore, total int64, err error) {
	db := m.db.Model(&DocumentScore{})

	for field, rangeValue := range opt.QueryRange {
		fields := m.FilterValidFields(DocumentScore{}.TableName(), field)
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
		fields := m.FilterValidFields(DocumentScore{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(fmt.Sprintf("%s in (?)", field), values)
	}

	for field, values := range opt.QueryLike {
		fields := m.FilterValidFields(DocumentScore{}.TableName(), field)
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
			m.logger.Error("GetDocumentScoreList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(DocumentScore{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) > 0 {
		var sorts []string
		for _, sort := range opt.Sort {
			slice := strings.Split(sort, " ")
			if len(m.FilterValidFields(DocumentScore{}.TableName(), slice[0])) == 0 {
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

	err = db.Find(&documentScoreList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetDocumentScoreList", zap.Error(err))
	}
	return
}

// DeleteDocumentScore 删除数据
// TODO: 删除数据之后，存在 document_score_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeleteDocumentScore(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&DocumentScore{}).Error
	if err != nil {
		m.logger.Error("DeleteDocumentScore", zap.Error(err))
	}
	return
}

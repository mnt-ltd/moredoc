package model

import (
	v1 "moredoc/api/v1"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Report struct {
	Id            int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	DocumentId    int64      `form:"document_id" json:"document_id,omitempty" gorm:"column:document_id;type:bigint(20);size:20;default:0;comment:文档ID;index:idx_document_id;"`
	DocumentTitle string     `form:"document_title" json:"document_title,omitempty" gorm:"column:document_title;type:varchar(255);size:255;default:'';comment:文档标题;"`
	UserId        int64      `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;comment:用户ID;index:idx_user_id;"`
	Username      string     `form:"username" json:"username,omitempty" gorm:"column:username;type:varchar(64);size:64;default:'';comment:用户名;"`
	Reason        int        `form:"reason" json:"reason,omitempty" gorm:"column:reason;type:int(11);size:11;default:0;comment:举报原因;"`
	Status        bool       `form:"status" json:"status,omitempty" gorm:"column:status;type:tinyint(4);size:4;default:0;comment:是否已处理;index:idx_status;"`
	Remark        string     `form:"remark" json:"remark,omitempty" gorm:"column:remark;type:varchar(255);size:255;default:'';comment:备注;"`
	CreatedAt     *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt     *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Report) TableName() string {
	return tablePrefix + "report"
}

// CreateReport 创建Report
func (m *DBModel) CreateReport(report *Report) (err error) {
	doc, _ := m.GetDocument(report.DocumentId, "id", "title")
	report.DocumentTitle = doc.Title
	user, _ := m.GetUser(report.UserId, "id", "username")
	report.Username = user.Username
	err = m.db.Create(report).Error
	if err != nil {
		m.logger.Error("CreateReport", zap.Error(err))
		return
	}
	return
}

// UpdateReport 更新Report，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateReport(report *Report, updateFields ...string) (err error) {
	db := m.db.Model(report)
	tableName := Report{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", report.Id).Updates(report).Error
	if err != nil {
		m.logger.Error("UpdateReport", zap.Error(err))
	}
	return
}

// GetReport 根据id获取Report
func (m *DBModel) GetReport(id interface{}, fields ...string) (report Report, err error) {
	db := m.db

	fields = m.FilterValidFields(Report{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&report).Error
	return
}

type OptionGetReportList struct {
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

// GetReportList 获取Report列表
func (m *DBModel) GetReportList(opt *OptionGetReportList) (reportList []*v1.Report, total int64, err error) {
	tableName := Report{}.TableName()
	db := m.db.Model(&Report{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetReportList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&reportList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetReportList", zap.Error(err))
	}
	return
}

// DeleteReport 删除数据
func (m *DBModel) DeleteReport(ids []int64) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Report{}).Error
	if err != nil {
		m.logger.Error("DeleteReport", zap.Error(err))
	}
	return
}

func (m *DBModel) GetReportByDocUser(docId, userId int64) (report Report, err error) {
	err = m.db.Where("doc_id = ? and user_id = ?", docId, userId).First(&report).Error
	return
}

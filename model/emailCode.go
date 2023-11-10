package model

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EmailCode struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Email     string     `form:"email" json:"email,omitempty" gorm:"column:email;type:varchar(64);size:64;index:idx_email;comment:邮箱;"`
	Ip        string     `form:"ip" json:"ip,omitempty" gorm:"column:ip;type:varchar(64);size:64;comment:IP地址;"`
	Code      string     `form:"code" json:"code,omitempty" gorm:"column:code;type:varchar(16);size:16;comment:邮箱验证码;"`
	CodeType  int        `form:"code_type" json:"code_type,omitempty" gorm:"column:code_type;type:smallint(6);size:6;default:0;comment:验证码类型,0注册,1登录;"`
	Success   bool       `form:"success" json:"success,omitempty" gorm:"column:success;type:tinyint(1);size:1;default:0;comment:是否发送成功;"`
	Error     string     `form:"error" json:"error,omitempty" gorm:"column:error;type:text;comment:错误信息;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:;"`
	IsUsed    bool       `form:"is_used" json:"is_used,omitempty" gorm:"column:is_used;type:tinyint(1);size:1;default:0;comment:是否已使用;"`
}

func (EmailCode) TableName() string {
	return tablePrefix + "email_code"
}

// CreateEmailCode 创建EmailCode
func (m *DBModel) CreateEmailCode(emailCode *EmailCode) (err error) {
	err = m.db.Create(emailCode).Error
	if err != nil {
		m.logger.Error("CreateEmailCode", zap.Error(err))
		return
	}
	return
}

// UpdateEmailCode 更新EmailCode，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdateEmailCode(emailCode *EmailCode, updateFields ...string) (err error) {
	db := m.db.Model(emailCode)
	tableName := EmailCode{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", emailCode.Id).Updates(emailCode).Error
	if err != nil {
		m.logger.Error("UpdateEmailCode", zap.Error(err))
	}
	return
}

// GetEmailCode 根据id获取EmailCode
func (m *DBModel) GetEmailCode(id int64, fields ...string) (emailCode EmailCode, err error) {
	db := m.db

	fields = m.FilterValidFields(EmailCode{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&emailCode).Error
	return
}

type OptionGetEmailCodeList struct {
	Page         int
	Size         int
	WithCount    bool                      // 是否返回总数
	Ids          []int64                   // id列表
	SelectFields []string                  // 查询字段
	QueryRange   map[string][2]interface{} // map[field][]{min,max}
	QueryIn      map[string][]interface{}  // map[field][]{value1,value2,...}
	QueryLike    map[string][]interface{}  // map[field][]{value1,value2,...}
	Sort         []string
}

// GetEmailCodeList 获取EmailCode列表
func (m *DBModel) GetEmailCodeList(opt *OptionGetEmailCodeList) (emailCodeList []EmailCode, total int64, err error) {
	tableName := EmailCode{}.TableName()
	db := m.db.Model(&EmailCode{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetEmailCodeList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&emailCodeList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetEmailCodeList", zap.Error(err))
	}
	return
}

// DeleteEmailCode 删除数据
func (m *DBModel) DeleteEmailCode(ids []int64) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&EmailCode{}).Error
	if err != nil {
		m.logger.Error("DeleteEmailCode", zap.Error(err))
	}
	return
}

// 获取最新一条邮箱验证码
func (m *DBModel) GetLatestEmailCode(email string, codeType int32) (code EmailCode) {
	err := m.db.Where("email = ? and code_type = ?", email, codeType).Order("id desc").Find(&code).Error
	if err != nil {
		m.logger.Error("GetLatestEmailCode", zap.Error(err))
	}
	return
}

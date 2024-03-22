package model

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	PunishmentTypeDisabled              = 1 // 禁用账户：禁止登录、禁止评论、禁止上传、禁止下载、禁止收藏
	PunishmentTypeCommentLimited        = 2 // 禁止评论
	PunishmentTypeUploadLimited         = 3 // 禁止上传
	PunishmentTypeDownloadLimited       = 4 // 禁止下载
	PunishmentTypeFavoriteLimited       = 5 // 禁止收藏
	PunishmentTypePublishArticleLimited = 6 // 禁止发布文章
)

type Punishment struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:自增主键;"`
	UserId    int64      `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;index:idx_user_id;comment:用户ID;"`
	Type      int        `form:"type" json:"type,omitempty" gorm:"column:type;type:int(11);size:11;default:0;comment:惩罚类型，对应user表的status;"`
	Enable    bool       `form:"enable" json:"enable,omitempty" gorm:"column:enable;type:tinyint(1);size:1;default:0;index:idx_enable;comment:0 关闭，1启用;"`
	Operators string     `form:"operators" json:"operators,omitempty" gorm:"column:operators;type:text;comment:操作信息;"`
	Reason    string     `form:"reason" json:"reason,omitempty" gorm:"column:reason;type:text;comment:惩罚原因;"`
	Remark    string     `form:"remark" json:"remark,omitempty" gorm:"column:remark;type:text;comment:惩罚备注;"`
	EndTime   *time.Time `form:"end_time" json:"end_time,omitempty" gorm:"column:end_time;type:datetime;comment:惩罚结束时间，没有结束时间，则表示永久;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

type PunishmentOperator struct {
	UserId    int64 `json:"u"`
	Type      int32 `json:"t"`
	Timestamp int64 `json:"ts"`
}

func (Punishment) TableName() string {
	return tablePrefix + "punishment"
}

func (m *DBModel) MakePunishmentOperators(userId int64, punishmentType int32, operaterStr ...string) string {
	var operators []PunishmentOperator

	if len(operaterStr) > 0 && operaterStr[0] != "" {
		err := jsoniter.Unmarshal([]byte(operaterStr[0]), &operators)
		if err != nil {
			m.logger.Error("FormatPunishmentOperators", zap.Error(err))
			return operaterStr[0]
		}
	}

	operators = append(operators, PunishmentOperator{
		UserId:    userId,
		Type:      punishmentType,
		Timestamp: time.Now().Unix(),
	})

	operatersByte, err := jsoniter.Marshal(operators)
	if err != nil {
		m.logger.Error("FormatPunishmentOperators", zap.Error(err))
		if len(operaterStr) > 0 {
			return operaterStr[0]
		}
		return ""
	}
	return string(operatersByte)
}

// CreatePunishment 创建Punishment
func (m *DBModel) CreatePunishment(punishment *Punishment) (err error) {
	err = m.db.Create(punishment).Error
	if err != nil {
		m.logger.Error("CreatePunishment", zap.Error(err))
		return
	}
	return
}

// UpdatePunishment 更新Punishment，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdatePunishment(punishment *Punishment, updateFields ...string) (err error) {
	db := m.db.Model(punishment)
	tableName := Punishment{}.TableName()

	updateFields = m.FilterValidFields(tableName, updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	} else { // 更新全部字段，包括零值字段
		db = db.Select(m.GetTableFields(tableName))
	}

	err = db.Where("id = ?", punishment.Id).Updates(punishment).Error
	if err != nil {
		m.logger.Error("UpdatePunishment", zap.Error(err))
	}
	return
}

// GetPunishment 根据id获取Punishment
func (m *DBModel) GetPunishment(id interface{}, fields ...string) (punishment Punishment, err error) {
	db := m.db

	fields = m.FilterValidFields(Punishment{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&punishment).Error
	return
}

type OptionGetPunishmentList struct {
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

// GetPunishmentList 获取Punishment列表
func (m *DBModel) GetPunishmentList(opt *OptionGetPunishmentList) (punishmentList []Punishment, total int64, err error) {
	tableName := Punishment{}.TableName()
	db := m.db.Model(&Punishment{})
	db = m.generateQueryRange(db, tableName, opt.QueryRange)
	db = m.generateQueryIn(db, tableName, opt.QueryIn)
	db = m.generateQueryLike(db, tableName, opt.QueryLike)

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetPunishmentList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(tableName, opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	db = m.generateQuerySort(db, tableName, opt.Sort)

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&punishmentList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetPunishmentList", zap.Error(err))
	}
	return
}

// DeletePunishment 删除数据
func (m *DBModel) DeletePunishment(ids []int64) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Punishment{}).Error
	if err != nil {
		m.logger.Error("DeletePunishment", zap.Error(err))
	}
	return
}

func (m *DBModel) isInPunishing(userId int64, types []int) (yes bool, err error) {
	if userId <= 1 {
		return false, nil
	}

	punishment := &Punishment{}
	err = m.db.Model(punishment).Select("id").
		Where(
			"user_id = ? and enable = ? and type in ? and (end_time IS NULL or end_time > ?)",
			userId, true, types, time.Now(),
		).Find(&punishment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		m.logger.Error("isInPunishing", zap.Error(err))
		return
	}
	return punishment.Id > 0, nil
}

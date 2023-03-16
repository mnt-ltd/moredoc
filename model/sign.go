package model

import (
	"fmt"
	"strconv"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Sign struct {
	Id        int64      `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	UserId    int64      `form:"user_id" json:"user_id,omitempty" gorm:"column:user_id;type:bigint(20);size:20;default:0;index:idx_user_id;index:idx_user_sign_at,unique;comment:签到的用户ID;"`
	Ip        string     `form:"ip" json:"ip,omitempty" gorm:"column:ip;type:varchar(64);size:64;comment:签到的用户IP;"`
	SignAt    int        `form:"sign_at" json:"sign_at,omitempty" gorm:"column:sign_at;type:int(11);size:11;default:0;index:idx_user_sign_at,unique;comment:签到时间，格式20060102;"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt *time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
	Award     int32      `form:"award" json:"award,omitempty" gorm:"column:award;type:int(11);size:11;default:0;comment:奖励积分;"`
}

func (Sign) TableName() string {
	return tablePrefix + "sign"
}

// CreateSign 用户签到
// 1. 创建签到记录
// 2. 更新用户签积分
func (m *DBModel) CreateSign(userId int64, ip string) (sign *Sign, err error) {
	now := time.Now()
	signAt, _ := strconv.Atoi(now.Format("20060102"))
	sign = &Sign{
		UserId: userId,
		Ip:     ip,
		SignAt: signAt,
	}

	cfg := m.GetConfigOfScore(ConfigScoreSignIn, ConfigScoreCreditName)
	sign.Award = cfg.SignIn

	tx := m.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Create(sign).Error
	if err != nil {
		m.logger.Error("CreateSign Create", zap.Error(err))
		return
	}

	content := "完成了每日签到"
	if cfg.SignIn > 0 {
		// 1. 更新用户积分
		err = tx.Model(&User{}).Where("id=?", userId).Update("credit_count", gorm.Expr("credit_count + ?", cfg.SignIn)).Error
		if err != nil {
			m.logger.Error("CreateSign Update", zap.Error(err))
			return
		}
		content = fmt.Sprintf("签到成功，获得 %d %s奖励", cfg.SignIn, cfg.CreditName)
	}

	dynamic := Dynamic{
		UserId:  userId,
		Type:    DynamicTypeSign,
		Content: content,
	}
	err = tx.Create(&dynamic).Error
	if err != nil {
		m.logger.Error("CreateSign Create Dynamic", zap.Error(err))
		return
	}
	return
}

// 用户今日是否已签到
func (m *DBModel) GetSignedToday(userId int64) (sign Sign) {
	signAt, _ := strconv.Atoi(time.Now().Format("20060102"))
	err := m.db.Where("user_id=? and sign_at=?", userId, signAt).First(&sign).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetSignedToday", zap.Error(err))
	}
	return
}

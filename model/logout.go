package model

import (
	"time"

	"go.uber.org/zap"
)

type Logout struct {
	Id        int64 `gorm:"primaryKey;autoIncrement"`
	UserId    int64
	UUID      string     `gorm:"column:uuid;type:varchar(36);size:36;not null;uniqueIndex;comment:jwt的uuid"`
	ExpiredAt int64      `gorm:"column:expired_at;type:bigint;not null;comment:过期时间，超过这个时间之后，可以从当前数据表中删除;index"`
	CreatedAt *time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
}

func (Logout) TableName() string {
	return tablePrefix + "logout"
}

// SetLogout 设置退出登录
func (m *DBModel) Logout(userId int64, uuid string, expiredAt int64) {
	// 将token加入到退出登录表中
	m.invalidToken.Store(uuid, struct{}{})
	m.validToken.Delete(uuid)
	logout := &Logout{
		UserId:    userId,
		UUID:      uuid,
		ExpiredAt: expiredAt,
	}
	if err := m.db.Create(logout).Error; err != nil {
		m.logger.Error("SetLogout", zap.Error(err))
	}
}

// IsInvalidToken 根据token对应的uuid，判断token是否有效
func (m *DBModel) IsInvalidToken(uuid string) bool {
	if _, ok := m.invalidToken.Load(uuid); ok {
		return ok
	}

	if _, ok := m.validToken.Load(uuid); ok {
		return !ok
	}

	logout := &Logout{}
	m.db.Select("id").Where("uuid = ?", uuid).First(logout)
	if logout.Id > 0 {
		m.invalidToken.Store(uuid, struct{}{})
		m.validToken.Delete(uuid)
		return true
	}
	m.validToken.Store(uuid, struct{}{})
	return false
}

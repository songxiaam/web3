package entity

import (
	"time"

	"github.com/google/uuid"
)

// User 数据库模型
// 对应 users 表
// 适用于 GORM/SQLX 等 ORM 框架

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Address   string    `gorm:"type:varchar(42);uniqueIndex;not null" json:"address"`
	Nonce     string    `gorm:"type:varchar(64)" json:"nonce"`
	Nickname  string    `gorm:"type:varchar(32)" json:"nickname"`
	Avatar    string    `gorm:"type:varchar(255)" json:"avatar"`
	Status    int       `gorm:"type:int;default:1" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}

package entity

import (
	"time"

	"github.com/google/uuid"
)

// Token 数据库模型
// 对应 tokens 表
// 适用于 GORM/SQLX 等 ORM 框架

type Token struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Chain     string    `gorm:"type:varchar(32);index;not null" json:"chain"`
	ChainId   int64     `gorm:"type:bigint;index;not null" json:"chain_id"`
	Symbol    string    `gorm:"type:varchar(32);index;not null" json:"symbol"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Address   string    `gorm:"type:varchar(64);index" json:"address"`
	Decimals  int       `gorm:"type:int" json:"decimals"`
	Logo      string    `gorm:"type:varchar(255)" json:"logo"`
	IsNative  bool      `gorm:"type:boolean" json:"is_native"`
	IsStable  bool      `gorm:"type:boolean" json:"is_stable"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Token) TableName() string {
	return "token"
}

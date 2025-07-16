package model

import (
	"time"

	"github.com/google/uuid"
)

// Asset 数据库模型
// 对应 assets 表（或 routes 表，视业务而定）
// 适用于 GORM/SQLX 等 ORM 框架

type Asset struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Chain     string    `gorm:"type:varchar(32);index;not null" json:"chain"`
	Symbol    string    `gorm:"type:varchar(32);index;not null" json:"symbol"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Contract  string    `gorm:"type:varchar(64);index" json:"contract"`
	Decimals  int       `gorm:"type:int" json:"decimals"`
	Logo      string    `gorm:"type:varchar(255)" json:"logo"`
	IsNative  bool      `gorm:"type:boolean" json:"is_native"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Asset) TableName() string {
	return "asset"
}

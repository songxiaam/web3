package entity

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Username  string    `gorm:"type:varchar(64);uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Role      string    `gorm:"type:varchar(32);index;not null" json:"role"`
	Group     string    `gorm:"type:varchar(32);index" json:"group"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Admin) TableName() string {
	return "admin"
}

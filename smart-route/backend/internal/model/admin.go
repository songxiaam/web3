package model

import (
	"github.com/google/uuid"
	"time"
)

// Admin 管理员业务模型结构体
type Admin struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Status    int       `json:"status"` // 1: 启用, 0: 禁用
	Group     string    `json:"group"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

package model

import (
	"time"

	"github.com/google/uuid"
)

// User 用户业务模型结构体
type User struct {
	ID        uuid.UUID `json:"id"`
	Address   string    `json:"address"`
	Nonce     string    `json:"nonce"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

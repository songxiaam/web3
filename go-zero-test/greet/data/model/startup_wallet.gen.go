// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameStartupWallet = "startup_wallet"

// StartupWallet mapped from table <startup_wallet>
type StartupWallet struct {
	ID            int64     `gorm:"column:id;primaryKey" json:"id"`
	ComerID       int64     `gorm:"column:comer_id;not null;comment:comer_id" json:"comer_id"`                   // comer_id
	StartupID     int64     `gorm:"column:startup_id;not null;comment:startup_id" json:"startup_id"`             // startup_id
	WalletName    string    `gorm:"column:wallet_name;not null;comment:wallet name" json:"wallet_name"`          // wallet name
	WalletAddress string    `gorm:"column:wallet_address;not null;comment:wallet address" json:"wallet_address"` // wallet address
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	IsDeleted     bool      `gorm:"column:is_deleted;not null;comment:Is Deleted" json:"is_deleted"` // Is Deleted
}

// TableName StartupWallet's table name
func (*StartupWallet) TableName() string {
	return TableNameStartupWallet
}

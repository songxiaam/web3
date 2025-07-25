// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGovernanceProposal = "governance_proposal"

// GovernanceProposal mapped from table <governance_proposal>
type GovernanceProposal struct {
	ID                  int32     `gorm:"column:id;primaryKey" json:"id"`
	StartupID           int64     `gorm:"column:startup_id;not null" json:"startup_id"`
	AuthorComerID       int64     `gorm:"column:author_comer_id;not null" json:"author_comer_id"`
	AuthorWalletAddress string    `gorm:"column:author_wallet_address;not null" json:"author_wallet_address"`
	ChainID             int64     `gorm:"column:chain_id;not null" json:"chain_id"`
	BlockNumber         int64     `gorm:"column:block_number;not null" json:"block_number"`
	ReleaseTimestamp    time.Time `gorm:"column:release_timestamp;not null;default:0000-00-00 00:00:00" json:"release_timestamp"`
	IpfsHash            string    `gorm:"column:ipfs_hash;not null" json:"ipfs_hash"`
	Title               string    `gorm:"column:title;not null" json:"title"`
	Description         string    `gorm:"column:description;not null" json:"description"`
	DiscussionLink      string    `gorm:"column:discussion_link;not null" json:"discussion_link"`
	VoteSystem          string    `gorm:"column:vote_system;not null" json:"vote_system"`
	StartTime           time.Time `gorm:"column:start_time;not null;default:0000-00-00 00:00:00" json:"start_time"`
	EndTime             time.Time `gorm:"column:end_time;not null;default:0000-00-00 00:00:00" json:"end_time"`
	Status              bool      `gorm:"column:status;not null;comment:0:pending 1:upcoming 2:active 3:ended" json:"status"` // 0:pending 1:upcoming 2:active 3:ended
	CreatedAt           time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	IsDeleted           bool      `gorm:"column:is_deleted;not null" json:"is_deleted"`
}

// TableName GovernanceProposal's table name
func (*GovernanceProposal) TableName() string {
	return TableNameGovernanceProposal
}

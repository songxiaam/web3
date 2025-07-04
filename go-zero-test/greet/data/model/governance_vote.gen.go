// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGovernanceVote = "governance_vote"

// GovernanceVote mapped from table <governance_vote>
type GovernanceVote struct {
	ID                 int64     `gorm:"column:id;primaryKey" json:"id"`
	ProposalID         int64     `gorm:"column:proposal_id;not null" json:"proposal_id"`
	VoterComerID       int64     `gorm:"column:voter_comer_id;not null" json:"voter_comer_id"`
	VoterWalletAddress string    `gorm:"column:voter_wallet_address;not null" json:"voter_wallet_address"`
	ChoiceItemID       int64     `gorm:"column:choice_item_id;not null" json:"choice_item_id"`
	ChoiceItemName     string    `gorm:"column:choice_item_name;not null" json:"choice_item_name"`
	Votes              float64   `gorm:"column:votes;not null" json:"votes"`
	IpfsHash           string    `gorm:"column:ipfs_hash;not null" json:"ipfs_hash"`
	CreatedAt          time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	IsDeleted          bool      `gorm:"column:is_deleted;not null" json:"is_deleted"`
}

// TableName GovernanceVote's table name
func (*GovernanceVote) TableName() string {
	return TableNameGovernanceVote
}

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTag = "tag"

// Tag mapped from table <tag>
type Tag struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null;comment:name" json:"name"`             // name
	IsIndex   bool      `gorm:"column:is_index;not null;comment:Is index" json:"is_index"` // Is index
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	IsDeleted bool      `gorm:"column:is_deleted;not null;comment:Is Deleted" json:"is_deleted"` // Is Deleted
	Category  string    `gorm:"column:category;not null" json:"category"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}

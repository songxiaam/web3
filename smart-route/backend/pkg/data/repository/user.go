package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"smart-route/pkg/data/entity"
)

// GetUserByAddress 根据地址查找用户
func GetUserByAddress(ctx context.Context, db *gorm.DB, address string) (*entity.User, error) {
	var user *entity.User
	err := db.WithContext(ctx).Where("address = ?", address).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser 支持通过 id 或 address 查询，id/address 不能同时为空
func GetUser(ctx context.Context, db *gorm.DB, id *uuid.UUID, address string) (*entity.User, error) {
	if (id == nil || *id == uuid.Nil) && address == "" {
		return nil, errors.New("id and address cannot both be empty")
	}
	var user *entity.User
	query := db.WithContext(ctx)
	if id != nil && *id != uuid.Nil {
		query = query.Where("id = ?", *id)
	}
	if address != "" {
		query = query.Where("address = ?", address)
	}
	err := query.First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser 创建新用户
func CreateUser(ctx context.Context, db *gorm.DB, address string) (*entity.User, error) {
	user := &entity.User{
		ID:      uuid.New(),
		Address: address,
	}
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

package service

import (
	"context"
	"smart-route/internal/model"
	"smart-route/pkg/convert"
	"smart-route/pkg/data/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUserByAddress(ctx context.Context, address string) (*model.User, error) {
	userEntity, err := repository.GetUserByAddress(ctx, s.db, address)
	if err != nil {
		return nil, err
	}
	var user *model.User
	convert.CopyStructFields(user, userEntity)
	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, id *uuid.UUID, address string) (*model.User, error) {
	userEntity, err := repository.GetUser(ctx, s.db, id, address)
	if err != nil {
		return nil, err
	}
	var user *model.User
	convert.CopyStructFields(user, userEntity)
	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, address string) (*model.User, error) {
	userEntity, err := repository.CreateUser(ctx, s.db, address)
	if err != nil {
		return nil, err
	}
	var user *model.User
	convert.CopyStructFields(user, userEntity)
	return user, nil
}

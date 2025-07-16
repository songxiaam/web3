package repository

import (
	"context"
	"gorm.io/gorm"
	"smart-route/pkg/data/entity"
)

// GetAdminByUsername 根据用户名查找管理员
func GetAdminByUsername(ctx context.Context, db *gorm.DB, username string) (*entity.Admin, error) {
	var admin *entity.Admin
	err := db.WithContext(ctx).Where("username = ?", username).First(admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// GetAdminByUsernameAndPassword 根据用户名和密码查找管理员
func GetAdminByUsernameAndPassword(ctx context.Context, db *gorm.DB, username, password string) (*entity.Admin, error) {
	var admin *entity.Admin
	err := db.WithContext(ctx).Where("username = ? AND password = ?", username, password).First(admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// CreateAdmin 创建管理员
func CreateAdmin(ctx context.Context, db *gorm.DB, admin *entity.Admin) error {
	return db.WithContext(ctx).Create(admin).Error
}

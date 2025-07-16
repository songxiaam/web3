package data

import (
	"fmt"
	"smart-route/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DataService 提供数据库访问入口
// 可扩展为 DDD 仓储实现

type DataService struct {
	DB *gorm.DB
}

// NewDB 根据配置初始化数据库连接
func NewDataService(pgCfg config.PostgreSQLConfig) (*DataService, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		pgCfg.Host,
		pgCfg.Port,
		pgCfg.User,
		pgCfg.Password,
		pgCfg.DBName,
		pgCfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return &DataService{DB: db}, nil
}

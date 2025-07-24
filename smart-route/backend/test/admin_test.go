package test

import (
	"context"
	"log"
	"os"
	"smart-route/internal/model"
	"smart-route/internal/service"
	"smart-route/pkg/config"
	"smart-route/pkg/data"
	"testing"
)

func TestInsertAdminUser(t *testing.T) {
	// 使用内存数据库进行测试
	// 加载配置文件
	configPath := "config/config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// 如果当前目录没有配置文件，尝试从项目根目录加载
		configPath = "../config/config.yaml"
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// 初始化数据库
	dataService, err := data.NewDataService(cfg.Database.PostgreSQL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	adminService := service.NewAdminService(dataService.DB)

	password := "123456"

	admin := &model.Admin{
		Username: "admin",
		Role:     "super",
		Group:    "default",
	}

	ctx := context.Background()
	if err := adminService.CreateAdmin(ctx, admin, password); err != nil {
		t.Fatalf("failed to insert admin user: %v", err)
	}

	// 验证插入
	got, err := adminService.GetAdminByUsername(ctx, "admin")
	if err != nil {
		t.Fatalf("failed to get admin user: %v", err)
	}
	if got.Username != admin.Username {
		t.Errorf("expected username %s, got %s", admin.Username, got.Username)
	}
}

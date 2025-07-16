// @title Smart Route API
// @version 1.0
// @description 聚合器+最优路径项目的 Go 后端 API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
package main

import (
	"fmt"
	"log"
	"os"
	_ "smart-route/docs"
	"smart-route/internal/router"
	"smart-route/pkg/config"
	"smart-route/pkg/data"
	"smart-route/pkg/data/entity"
	"smart-route/pkg/redis"
)

// @Summary 启动服务
// @Description 启动 smart-route 后端服务
func main() {
	// 加载配置文件
	configPath := "config/config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// 如果当前目录没有配置文件，尝试从项目根目录加载
		configPath = "../../config/config.yaml"
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

	// 自动迁移表结构
	if err := dataService.DB.AutoMigrate(
		&entity.Admin{},
		&entity.Token{},
		&entity.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	// 初始化 Redis
	redisClient := redis.NewRedisClientByConfig(cfg.Redis)
	if err := redis.RedisPing(redisClient); err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	port := os.Getenv("SMART_ROUTE_PORT")
	if port == "" {
		port = fmt.Sprintf("%d", cfg.Server.Port)
	}

	r := router.NewRouter(cfg, dataService, redisClient)
	fmt.Printf("Starting smart-route backend (gin) on :%s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

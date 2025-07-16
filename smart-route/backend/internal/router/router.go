package router

import (
	"smart-route/pkg/config"
	"smart-route/pkg/data"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter 初始化 gin 路由和基础配置
func NewRouter(cfg *config.Config, ds *data.DataService, redis *redis.Client) *gin.Engine {
	r := gin.Default()

	// 加入跨域中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"}, // 允许访问的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 你的路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 注册 Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	RegisterRoutes(r, cfg, ds, redis)

	return r
}

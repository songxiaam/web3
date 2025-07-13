package router

import (
	"smart-route/internal/auth"
	"smart-route/pkg/config"

	"github.com/gin-gonic/gin"
)

// NewRouter 初始化 gin 路由和基础配置
func NewRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	authService := auth.NewAuthService(cfg)

	RegisterRoutes(r, authService)

	return r
}

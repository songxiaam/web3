package router

import (
	"smart-route/internal/auth"
	"smart-route/internal/handler"
	"smart-route/internal/service"
	"smart-route/pkg/config"
	"smart-route/pkg/data"

	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有 API 路由
func RegisterRoutes(r *gin.Engine, cfg *config.Config, ds *data.DataService, redis *redis.Client) {

	jwtAuth := auth.NewJwtAuth(cfg, ds.DB)

	userService := service.NewUserService(ds.DB)

	// 创建 handler
	userHandler := handler.NewUserHandler(jwtAuth, userService)

	// 用户接口
	r.POST("/login", userHandler.Login)
	user := r.Group("/user", jwtAuth.AuthMiddleware())
	{
		user.Use(jwtAuth.AuthMiddleware())
		user.GET("/profile", userHandler.GetProfile)
	}

	tokenService := service.NewTokenService(ds.DB)
	tokenHandler := handler.NewTokenHandler(tokenService)
	token := r.Group("/token")
	{
		token.GET("/list", tokenHandler.ListTokens)
	}

	// 管理员接口,用于后台管理系统
	adminUserService := service.NewAdminService(ds.DB)
	adminUserHandler := handler.NewAdminHandler(jwtAuth, adminUserService)

	r.POST("/admin/login", adminUserHandler.Login)
	admin := r.Group("/admin", jwtAuth.AuthMiddleware())
	{
		admin.POST("/token/create", tokenHandler.CreateToken)
		admin.GET("/token/list", tokenHandler.ListTokens)
	}
}

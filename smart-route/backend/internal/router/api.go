package router

import (
	"net/http"
	"smart-route/internal/auth"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有 API 路由
func RegisterRoutes(r *gin.Engine, authService *auth.AuthService) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/login", loginHandler)

	// 示例：受保护路由
	api := r.Group("/api")
	api.Use(authService.AuthMiddleware())
	api.GET("/profile", profileHandler)

}

// loginHandler 用户登录接口（示例）
func loginHandler(c *gin.Context) {
	var req struct {
		Address   string `json:"address" binding:"required"`
		Signature string `json:"signature" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: 验证签名、生成 token
	c.JSON(http.StatusOK, gin.H{"token": "mock-jwt-token", "address": req.Address})
}

// profileHandler 示例受保护接口
func profileHandler(c *gin.Context) {
	address, _ := c.Get("address")
	c.JSON(http.StatusOK, gin.H{"address": address, "message": "Profile retrieved successfully"})
}

package api

import (
	"github.com/gin-gonic/gin"
)

// NewRouter 创建 gin 路由
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 可在此添加更多 API 路由

	return r
}

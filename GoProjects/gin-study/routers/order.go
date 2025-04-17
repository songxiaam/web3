package routers

import (
	"gin-study/controllers"
	"github.com/gin-gonic/gin"
)

func InitOrderRoutes(r *gin.Engine) {
	orderGroup := r.Group("/pt/order")
	{
		orderGroup.POST("/login", controllers.Order.OrderDetail)
	}
}

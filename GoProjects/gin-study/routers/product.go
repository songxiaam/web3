package routers

import (
	"gin-study/controllers"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(r *gin.Engine) {
	productGroup := r.Group("/pt/product")
	{
		productGroup.POST("/list", controllers.User.LoginAccount)
	}
}

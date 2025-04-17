package routers

import (
	"gin-study/controllers"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", controllers.User.LoginAccount)
		userGroup.GET("/ping", controllers.User.UserPing)
		userGroup.POST("/get", controllers.User.GetUserInfo)

	}
}

package routers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	InitUserRoutes(r)
	//InitOrderRoutes(r)
	//InitProductRoutes(r)
	return r
}

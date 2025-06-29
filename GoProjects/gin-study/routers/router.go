package routers

import (
	"fmt"
	"gin-study/dtm_demo"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		fmt.Println("==> 请求路径:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	InitUserRoutes(r)
	//InitOrderRoutes(r)
	//InitProductRoutes(r)
	dtm_demo.InitOrderRoutes(r)
	return r
}

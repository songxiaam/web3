package routes

import (
	"github.com/gin-gonic/gin"
	"pledgev2-backend/config"
	"pledgev2-backend/pkg/controllers"
)

func InitRoute(e *gin.Engine) *gin.Engine {
	v2Group := e.Group("/pkg/v" + config.Config.Env.Version)

	poolBaseController := controllers.PoolBaseController{}
	v2Group.GET("/poolBaseInfo", poolBaseController.GetPoolBaseInfo)
	return e
}

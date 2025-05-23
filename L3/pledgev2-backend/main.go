package main

import (
	"github.com/gin-gonic/gin"
	"pledgev2-backend/config"
	"pledgev2-backend/middlewares"
	"pledgev2-backend/pkg/db"
	"pledgev2-backend/pkg/models/kucoin"
	"pledgev2-backend/pkg/models/ws"
	"pledgev2-backend/pkg/routes"
	"pledgev2-backend/pkg/static"
	"pledgev2-backend/pkg/validate"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	db.InitMysql()
	db.InitRedis()

	validate.BindingValidator()

	go ws.StartServer()

	go kucoin.GetExchangePrice()

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	staticPath := static.GetCurrentAbPathByCaller()
	app.Static("/storage", staticPath)
	app.Use(middlewares.Cors())
	routes.InitRoute(app)
	_ = app.Run(":" + config.Config.Env.Port)
}

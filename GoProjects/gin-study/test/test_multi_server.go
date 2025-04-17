package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "ok",
			"msg":   "welcome to server-01",
		})
	})
	return e
}

func Router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "ok",
			"msg":   "welcome to server-02",
		})
	})
	return e
}

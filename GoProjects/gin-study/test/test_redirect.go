package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestRedirect1(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com")
}

func TestRedirect2(c *gin.Context) {
	c.Redirect(http.StatusFound, "/test4")
}

func TestRedirect3(c *gin.Context, router *gin.Engine) {
	gin.Default()
	c.Request.URL.Path = "/test4"
	router.HandleContext(c)
}

func TestRedirect4(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

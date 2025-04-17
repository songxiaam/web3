package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestUserName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func TestUserNameAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

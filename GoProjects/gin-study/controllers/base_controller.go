package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (b *BaseController) Success(c *gin.Context, data map[string]any) {

	baseResponse := gin.H{
		"code": 1000,
		"msg":  "success",
	}
	for k, v := range data {
		baseResponse[k] = v
	}

	c.JSON(http.StatusOK, baseResponse)
}

func (b *BaseController) Fail(c *gin.Context, msg string) {
	baseResponse := gin.H{
		"code": 1001,
		"msg":  msg,
	}
	c.JSON(http.StatusBadRequest, baseResponse)
}

func (b *BaseController) get(c *gin.Context, msg string) {}

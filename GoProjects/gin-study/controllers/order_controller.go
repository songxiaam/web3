package controllers

import (
	"gin-study/models"
	"gin-study/protocol"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	BaseController
}

var Order = &OrderController{}

func (b *OrderController) OrderDetail(c *gin.Context) {
	var req protocol.OrderDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		b.Fail(c, err.Error())
		return
	}

	user, err := models.GetOrderById(req.ID)
	if err != nil {
		b.Fail(c, err.Error())
		return
	}

	b.Success(c, gin.H{"user": user})
}

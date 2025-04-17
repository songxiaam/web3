package controllers

import (
	"gin-study/protocol"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	BaseController
}

var Product = &ProductController{}

func (u *ProductController) ProductList(c *gin.Context) {
	var req protocol.ProductListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.Fail(c, err.Error())
		return
	}

	//user, err := models.FindUserByLoginNameAndPassword(req.LoginName, req.Password)
	//if err != nil {
	//	u.Fail(c, err.Error())
	//	return
	//}

	//u.Success(c, gin.H{"list": []})

}

package controllers

import (
	"fmt"
	"gin-study/models"
	"gin-study/protocol"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

var User = &UserController{}

func (u *UserController) LoginAccount(c *gin.Context) {
	fmt.Println("LoginAccount")
	var req protocol.LoginRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		u.Fail(c, err.Error())
		fmt.Println(err.Error())
		return
	}

	fmt.Printf(req.LoginName, req.Password)

	user, err := models.FindUserByLoginNameAndPassword(req.LoginName, req.Password)
	if err != nil {
		u.Fail(c, err.Error())
		return
	}

	u.Success(c, gin.H{"user": user})
}

func (u *UserController) UserPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	c.JSON(http.StatusOK, user)
}

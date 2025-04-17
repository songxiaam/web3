package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthParams struct {
	Token string `cookie:"token" binding:"-"`
}

func TestCookie(c *gin.Context) {
	fmt.Println("---------")
	cookie, err := c.Cookie("token")

	fmt.Printf("cookie=%s\n", cookie)

	var params AuthParams
	if err := c.ShouldBind(&params); err != nil {
		//fmt.Println(params.Token)
		fmt.Println(err.Error())
	} else {
		fmt.Printf("token=%s\n", params.Token)
	}

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("token", "test", 3600, "/", "localhost", false, true)
	}
	fmt.Printf("cookie value: %s \n", cookie)
}

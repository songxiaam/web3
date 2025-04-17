package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟私人数据
var secrets = gin.H{
	"mikew": gin.H{"email": "mike@gmail.com", "name": "mike", "age": 24, "gender": "male", "phone": "13221123122"},
	"john":  gin.H{"email": "john@sina.com", "name": "john", "age": 22, "gender": "female", "phone": "13921123122"},
	"hana":  gin.H{"email": "hana@163.com", "name": "hana", "age": 20, "gender": "female", "phone": "18588882222"},
}

func GetSecretUser(router *gin.Engine) {
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"mike": "zhao",
		"john": "wang",
		"hana": "li",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
}

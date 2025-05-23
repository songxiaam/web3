package middlewares

import (
	"github.com/gin-gonic/gin"
	"pledgev2-backend/config"
	"pledgev2-backend/pkg/common/statecode"
	"pledgev2-backend/pkg/db"
	"pledgev2-backend/pkg/models/response"
	"pledgev2-backend/utils"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := response.Gin{Res: c}
		token := c.Request.Header.Get("token")

		username, err := utils.ParseToken(token, config.Config.Jwt.SecretKey)
		if err != nil {
			res.Response(c, statecode.TokenErr, nil)
			c.Abort()
			return
		}

		if username != config.Config.DefaultAdmin.UserName {
			res.Response(c, statecode.TokenErr, nil)
			c.Abort()
			return
		}

		resByteArr, err := db.RedisGet(username)
		if string(resByteArr) != `"login_ok"` {
			res.Response(c, statecode.TokenErr, nil)
			c.Abort()
			return
		}
		c.Set("user_name", username)
		c.Next()
	}
}

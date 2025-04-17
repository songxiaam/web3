package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestSecureJSON(c *gin.Context) {
	names := []string{"a", "b", "c", "d"}
	c.SecureJSON(http.StatusOK, names)
}

package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UseHeader(c *gin.Context) {
	expectedHost := "localhost:8080"
	if c.Request.Host != expectedHost {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invail host header",
		})
		return
	}
	c.Header("X-Frame-Options", "DENY")
	c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
	c.Header("X-XSS-Protection", "1; mode=block")
	c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	c.Header("Referrer-Policy", "strict-origin")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("Permissions-Policy", "geolocation=(), midi=(), sync-xhr=(), microphone=(), camera=(), magnetometer=(), gyroscope=(), fullscreen=(self), payment=()")
	c.Next()
}

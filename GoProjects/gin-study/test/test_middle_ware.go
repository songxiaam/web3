package test

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("start log")
		t := time.Now()

		c.Set("example", "12345")

		// before request
		// do....
		log.Println("before")
		c.Next()
		log.Println("after")

		// after request
		latency := time.Since(t)
		log.Println(latency)

		status := c.Writer.Status()
		log.Println(status)

	}
}

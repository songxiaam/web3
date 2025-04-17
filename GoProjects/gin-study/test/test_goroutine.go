package test

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LongAsync(c *gin.Context) {
	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}

func LongSync(c *gin.Context) {
	time.Sleep(5 * time.Second)

	log.Println("Done! in path " + c.Request.URL.Path)
}

package test

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func RenderJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
		"status":  http.StatusOK,
	})
}

func RenderMoreJSON(c *gin.Context) {
	var data = struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Number int    `json:"number"`
	}{
		Number: 108,
		Name:   "Mike",
		Age:    18,
	}
	c.JSON(http.StatusOK, data)
}

func RenderXML(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"message": "hello world",
		"status":  http.StatusOK,
	})
}

func RenderYAML(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{
		"message": "hello world",
		"status":  http.StatusOK,
	})
}

func RenderProtoBuf(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	label := "test"
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(http.StatusOK, data)

}

package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type fromA struct {
	Foo string `json:"foo" xml:"foo" form:"foo" binding:"required"`
}

type fromB struct {
	Bar string `json:"bar" xml:"bar" form:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := fromA{}
	objB := fromB{}

	if err := c.ShouldBind(&objA); err == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {

	}
}

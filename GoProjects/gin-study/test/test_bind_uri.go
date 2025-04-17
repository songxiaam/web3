package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PersonUri struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func BindUri(c *gin.Context) {
	var person PersonUri
	if err := c.ShouldBindUri(&person); err != nil {
		fmt.Println(person.ID)
		fmt.Println(person.Name)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})

}

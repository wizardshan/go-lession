package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter1_5/request"
	"net/http"
)

type Column struct{}

func NewColumn() *Column {
	ctr := new(Column)
	return ctr
}

func (ctr *Column) Create(c *gin.Context) {
	request := new(request.ColumnCreate)
	if err := c.ShouldB(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"data": request})
}

func (ctr *Column) Detail(c *gin.Context) {
	request := new(request.ColumnDetail)
	if err := c.ShouldB(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"data": request})
}

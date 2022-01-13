package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter4_1/repository"
	"go-web/lesson/chapter4_1/request"
	"net/http"
)

type Order struct{
	repo *repository.Order
}

func NewOrder(repo *repository.Order) *Order {
	ctr := new(Order)
	ctr.repo = repo
	return ctr
}

func (ctr *Order) Get(c *gin.Context) {

	request := new(request.OrderGet)
	if err := c.ShouldB(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := ctr.repo.Get(c.Request.Context(), request.ID)
	c.JSON(http.StatusOK, gin.H{"data": order.MappingGet()})
}

func (ctr *Order) My(c *gin.Context) {

	orders := ctr.repo.My(c.Request.Context())

	c.JSON(http.StatusOK, gin.H{"data": orders.MappingMy()})
}

func (ctr *Order) List(c *gin.Context) {

	request := new(request.OrderList)
	if err := c.ShouldB(request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders := ctr.repo.List(c.Request.Context(), request)
	c.JSON(http.StatusOK, gin.H{"data": orders.MappingList()})
}
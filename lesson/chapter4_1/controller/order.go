package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter4_1/repository"
	"go-web/lesson/chapter4_1/request"
	"go-web/lesson/chapter4_1/response"
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

	resp := new(response.OrderGet)
	c.JSON(http.StatusOK, gin.H{"data": resp.Map(order)})
}

func (ctr *Order) My(c *gin.Context) {

	orders := ctr.repo.My(c.Request.Context())

	resp := response.OrderMys{}
	c.JSON(http.StatusOK, gin.H{"data": resp.Map(orders)})
}

func (ctr *Order) List(c *gin.Context) {

	request := new(request.OrderList)
	if err := c.ShouldB(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders := ctr.repo.List(c.Request.Context(), request)

	resp := response.OrderLists{}
	c.JSON(http.StatusOK, gin.H{"data": resp.Map(orders)})
}
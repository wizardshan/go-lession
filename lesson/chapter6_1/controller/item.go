package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter6_1/domain"
	"go-web/lesson/chapter6_1/pkg/mapper"
	"go-web/lesson/chapter6_1/repository"
	"go-web/lesson/chapter6_1/request"
	"go-web/lesson/chapter6_1/response"
	"net/http"
)

type Item struct{
	repo *repository.Item
}

func NewGoods(repo *repository.Item) *Item {
	ctr := new(Item)
	ctr.repo = repo
	return ctr
}

func (ctr *Item) Category(c *gin.Context) {

	var request request.GoodsCategory
	if err := c.ShouldB(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := ctr.repo.Category(c.Request.Context(), request.ID)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (ctr *Item) V1List(c *gin.Context) {
	data := ctr.repo.V1List(c.Request.Context())
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (ctr *Item) V2List(c *gin.Context) {

	data := ctr.repo.V2List(c.Request.Context())

	response := []*response.Goods{}
	mapper.Map(&response, data)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *Item) Get(c *gin.Context) {
	data := ctr.repo.Get(c.Request.Context(), 1)

	response := new(response.Goods)
	mapper.Map(&response, data)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *Item) List(c *gin.Context) {

	data := ctr.repo.List(c.Request.Context())

	response := []*response.Goods{}
	mapper.Map(&response, data)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *Item) Create(c *gin.Context) {

	goods := &domain.Item{Name: "商品", CategoryID: 1}
	data := ctr.repo.Create(c.Request.Context(), goods)

	response := new(response.Goods)
	mapper.Map(&response, data)
	c.JSON(http.StatusOK, gin.H{"data": response})
}




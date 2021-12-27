package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_1/domain"
	"go-web/lesson/chapter5_1/pkg/mapper"
	"go-web/lesson/chapter5_1/repository"
	"go-web/lesson/chapter5_1/request"
	"go-web/lesson/chapter5_1/response"
	"net/http"
)

type Goods struct{
	repo *repository.Goods
}

func NewGoods(repo *repository.Goods) *Goods {
	ctr := new(Goods)
	ctr.repo = repo
	return ctr
}

func (ctr *Goods) Category(c *gin.Context) {

	var request request.GoodsCategory
	if err := c.ShouldB(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := ctr.repo.Category(c.Request.Context(), request.ID)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (ctr *Goods) V1List(c *gin.Context) {
	data := ctr.repo.V1List(c.Request.Context())
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (ctr *Goods) V2List(c *gin.Context) {

	data := ctr.repo.V2List(c.Request.Context())
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (ctr *Goods) V3List(c *gin.Context) {

	data := ctr.repo.V2List(c.Request.Context())

	response := []*response.Goods{}
	mapper.Map(&response, data)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *Goods) Get(c *gin.Context) {
	data := ctr.repo.Get(c.Request.Context(), 1)

	response := new(response.Goods)
	mapper.Map(&response, data)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *Goods) List(c *gin.Context) {

	data := ctr.repo.List(c.Request.Context())

	response := []*response.Goods{}
	mapper.Map(&response, data)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *Goods) Create(c *gin.Context) {

	goods := &domain.Goods{Name: "商品", CategoryID: 1}
	data := ctr.repo.Create(c.Request.Context(), goods)

	response := new(response.Goods)
	mapper.Map(&response, data)
	c.JSON(http.StatusOK, gin.H{"data": response})
}




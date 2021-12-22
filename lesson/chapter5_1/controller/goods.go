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

func (ctr *Goods) V1List(c *gin.Context) {
	data, err := ctr.repo.V1List(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (ctr *Goods) V1Category(c *gin.Context) {

	var request request.GoodsCategory
	if err := c.ShouldB(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctr.repo.Category(c.Request.Context(), request.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (ctr *Goods) V2List(c *gin.Context) {

	var goods []*domain.Goods
	err := ctr.repo.V2List(c.Request.Context(), &goods)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": goods})
}

func (ctr *Goods) V3List(c *gin.Context) {

	var goods []*domain.Goods
	err := ctr.repo.V2List(c.Request.Context(), &goods)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []*response.Goods
	mapper.Map(&response, goods)
	c.JSON(http.StatusOK, gin.H{"data": response})
}


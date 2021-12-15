package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go-web/lesson/chapter2_1/domain"
	"go-web/lesson/chapter2_1/response"
	"net/http"
)

type Article struct {}

func (ctr *Article) List(c *gin.Context) {


	articles := []*domain.Article {
		{ID:1, Title: "标题1", Content: "内容1", ImageUrl: "1.jpg", Author: &domain.Author{ID: 1, Name: "作者1", Desc: "介绍1"}},
		{ID:2, Title: "标题2", Content: "内容2", ImageUrl: "2.jpg", Author: &domain.Author{ID: 2, Name: "作者2", Desc: "介绍2"}},
	}

	var response []*response.ArticleList
	copier.Copy(&response, articles)
	c.JSON(http.StatusOK, gin.H{"data": response})

}

func (ctr *Article) Detail(c *gin.Context) {
	article := &domain.Article{ID: 1, Title: "标题1", Content: "内容1", ImageUrl: "1.jpg"}

	var response response.Article
	copier.Copy(&response, article)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

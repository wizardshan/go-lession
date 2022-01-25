package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter2_1/domain"
	"go-web/lesson/chapter2_1/response"
	"net/http"
)

type Article struct{}

func NewArticle() *Article {
	ctr := new(Article)
	return ctr
}

func (ctr *Article) List(c *gin.Context) {

	articles := []*domain.Article{
		{ID: 1, Title: "标题1", Content: "内容1", ImageURL: "1.jpg", Author: &domain.Author{ID: 1, Name: "作者1", Desc: "介绍1"}},
		{ID: 2, Title: "标题2", Content: "内容2", ImageURL: "2.jpg", Author: &domain.Author{ID: 2, Name: "作者2", Desc: "介绍2"}},
	}

	//var response []*response.ArticleList
	//copier.Copy(&response, articles)
	//var response = new(response.ArticleLists)
	response := response.ArticleLists{}
	c.JSON(http.StatusOK, gin.H{"data": response.Map(articles)})
}

func (ctr *Article) Detail(c *gin.Context) {
	article := &domain.Article{ID: 1, Title: "标题1", Content: "内容1", ImageURL: "1.jpg"}

	//var response response.ArticleDetail
	//copier.Copy(&response, article)
	response := new(response.ArticleDetail)
	c.JSON(http.StatusOK, gin.H{"data": response.Map(article)})
}

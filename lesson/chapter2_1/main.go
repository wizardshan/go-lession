package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter2_1/controller"
)

func main() {
	engine := gin.New()

	ctrArticle := new(controller.Article)
	engine.GET("/articles", ctrArticle.List)
	engine.GET("/articles/:id", ctrArticle.Detail)
	engine.Run()
}
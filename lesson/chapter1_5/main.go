package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter1_5/controller"
)

func main() {
	engine := gin.New()

	ctrColumn := new(controller.Column)
	engine.POST("/column/create", ctrColumn.Create)
	engine.GET("/column", ctrColumn.Detail)
	engine.Run()
}


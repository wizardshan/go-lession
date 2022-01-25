package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter1_5/controller"
)

func main() {
	engine := gin.New()

	ctrColumn := controller.NewColumn()
	engine.POST("/column/create", ctrColumn.Create)
	engine.GET("/column", ctrColumn.Detail)
	engine.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter3_1/controller"
)

func main() {
	engine := gin.New()

	ctrUser := new(controller.User)
	engine.POST("/user/cash", ctrUser.Cash)
	engine.GET("/users/:id", ctrUser.Info)
	engine.POST("/user/update", ctrUser.Update)
	engine.Run()
}

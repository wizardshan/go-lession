package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter1_4/controller"
)

func main() {
	engine := gin.New()

	ctrUser := new(controller.User)
	engine.POST("/user/login", ctrUser.Login)
	engine.POST("/user/register", ctrUser.Register)
	engine.Run()
}
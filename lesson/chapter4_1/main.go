package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter4_1/controller"
	"go-web/lesson/chapter4_1/repository"
)

func main() {
	engine := gin.New()

	repoOrder := repository.NewOrder()
	ctrOrder := controller.NewOrder(repoOrder)
	engine.GET("/order", ctrOrder.Get)
	engine.GET("/orders/my", ctrOrder.My)
	engine.GET("/orders", ctrOrder.List)


	repoUser := repository.NewUser()
	ctrUser := controller.NewUser(repoUser)
	engine.POST("/user/register", ctrUser.Register)

	engine.Run()
}

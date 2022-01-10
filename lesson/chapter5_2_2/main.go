package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_2_2/controller"
	"go-web/lesson/chapter5_2_2/repository"
)

func main() {
	engine := gin.New()

	repoActivity := repository.NewActivity()
	ctrActivity := controller.NewActivity(repoActivity)
	engine.GET("/activity/lottery", ctrActivity.Lottery)

	engine.Run()
}

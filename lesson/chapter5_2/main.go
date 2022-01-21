package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_2/controller"
	"go-web/lesson/chapter5_2/repository"
)

func main() {
	engine := gin.New()

	repoUser := repository.NewUser()
	ctrUser := controller.NewUser(repoUser)
	engine.GET("/user/signs", ctrUser.Signs)
	engine.GET("/user/signs/v1", ctrUser.SignsV1)
	engine.GET("/user/signs/v2", ctrUser.SignsV2)
	engine.GET("/user/signs/v3", ctrUser.SignsV3)

	engine.Run()
}

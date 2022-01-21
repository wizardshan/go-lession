package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_1_3/controller"
	"go-web/lesson/chapter5_1_3/repository"
)

func main() {
	engine := gin.New()

	repoCaptcha := repository.NewCaptcha()
	ctrCaptcha := controller.NewCaptcha(repoCaptcha)
	engine.POST("/captcha/send", ctrCaptcha.Send)

	engine.Run()
}

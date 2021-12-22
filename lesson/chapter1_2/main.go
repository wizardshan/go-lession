package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter1_2/controller"
)

func main() {
	engine := gin.New()

	ctrUser := new(controller.User)
	ctrCaptcha := new(controller.Captcha)
	engine.POST("/captcha/send", ctrCaptcha.Send)
	engine.POST("/user/login", ctrUser.Login)
	engine.Run()
}

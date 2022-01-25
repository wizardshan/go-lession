package main

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter1_1/controller"
)

func main() {
	engine := gin.New()

	ctrUser := controller.NewUser()
	engine.POST("/user/login", ctrUser.Login)

	ctrCaptcha := controller.NewCaptcha()
	engine.POST("/captcha/send", ctrCaptcha.Send)

	engine.Run()
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type User struct {}

func (ctr *User) Login(c *gin.Context) {
	mobile := c.PostForm("mobile")
	captcha := c.PostForm("captcha")

	// 校验手机号逻辑
	if mobile == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "手机号不能为空"})
		return
	}

	matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile)
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "手机号格式不正确"})
		return
	}

	// 校验手机号逻辑
	if captcha == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证码不能为空"})
		return
	}

	if len(captcha) != 4 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证码为4位"})
		return
	}

	fmt.Println(mobile, captcha)
}

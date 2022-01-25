package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type Captcha struct{}

func NewCaptcha() *Captcha {
	ctr := new(Captcha)
	return ctr
}

func (ctr *Captcha) Send(c *gin.Context) {
	mobile := c.PostForm("mobile")

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

	c.JSON(http.StatusBadRequest, gin.H{"mobile": mobile})
}

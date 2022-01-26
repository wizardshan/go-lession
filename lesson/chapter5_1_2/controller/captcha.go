package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_1_2/domain"
	"go-web/lesson/chapter5_1_2/pkg/sms"
	"go-web/lesson/chapter5_1_2/repository"
	"go-web/lesson/chapter5_1_2/request"
	"net/http"
)

type Captcha struct{
	repo *repository.Captcha
}

func NewCaptcha(repo *repository.Captcha) *Captcha {
	ctr := new(Captcha)
	ctr.repo = repo
	return ctr
}

func (ctr *Captcha) Send(c *gin.Context) {
	request := new(request.CaptchaSend)
	if err := c.ShouldB(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	captcha := domain.NewCaptcha(request.Mobile)
	captcha.Generate(domain.CaptchaCategoryLogin)

	ctr.repo.Save(captcha)
	sms.Send(captcha.Mobile, captcha.Content)

	c.JSON(http.StatusBadRequest, gin.H{"data": captcha})
}

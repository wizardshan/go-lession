package controller

import (
	validator "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter1_2/request"
	"net/http"
)

type Captcha struct{}

func NewCaptcha() *Captcha {
	ctr := new(Captcha)
	return ctr
}

func (ctr *Captcha) Send(c *gin.Context) {
	request := new(request.Captcha)
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := validator.ValidateStruct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"data": request})
}

package repository

import (
	"go-web/lesson/chapter5_1_1/domain"
)

type Captcha struct {
}

func NewCaptcha() *Captcha {
	repo := new(Captcha)
	return repo
}

func (repo *Captcha) Save(c *domain.Captcha) {

}
package domain

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	CaptchaCategoryLogin = "login"
	CaptchaCategoryResetPassword = "resetPassword"
)

type Captcha struct {
	Code string
	Mobile string
	Content string
	ExpireTime time.Time

	Category string
}

func NewCaptcha(category string) *Captcha {
	c := new(Captcha)
	c.Category = category
	return c
}

// 面向过程或者是面向功能点
func (dom *Captcha) Generate(mobile string) {
	dom.Mobile = mobile
	dom.generateCode()
	dom.generateContent()
	dom.generateExpireTime()
}

func (dom *Captcha) generateCode() {
	dom.Code = fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

func (dom *Captcha) generateContent() {
	dom.Content = fmt.Sprintf(dom.getContent(dom.Category), dom.Code)
}

func (dom *Captcha) generateExpireTime() {
	dom.ExpireTime = time.Now().Add(time.Minute * time.Duration(dom.expireMinutes()))
}

func (dom *Captcha) expireMinutes() int {
	return 5
}

func (dom *Captcha) getContent(category string) string {
	return dom.contentsByCategory()[category]
}

func (dom *Captcha) contentsByCategory() map[string]string {
	return map[string]string {
		CaptchaCategoryLogin : "%s（登录验证码，请完成验证），如非本人操作，请忽略本短信。",
		CaptchaCategoryResetPassword : "%s （重置密码验证码，请完成重置），如非本人操作，请忽略本短信。",
	}
}



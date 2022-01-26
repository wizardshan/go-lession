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

func NewCaptcha(mobile string, category string) *Captcha {
	dom := new(Captcha)
	dom.Mobile = mobile
	dom.Category = category
	return dom
}

// 面向过程或者是面向功能点
func (dom *Captcha) Generate() {
	dom.generateCode()
	dom.generateContent()
	dom.generateExpireTime()
}

func (dom *Captcha) generateCode() {
	dom.Code = fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

func (dom *Captcha) generateContent() {
	dom.Content = fmt.Sprintf(dom.contentTemplate(dom.Category), dom.Code)
}

func (dom *Captcha) generateExpireTime() {
	dom.ExpireTime = time.Now().Add(time.Minute * time.Duration(dom.expireMinutes()))
}

func (dom *Captcha) expireMinutes() int {
	return 5
}

func (dom *Captcha) contentTemplate(category string) string {
	switch category {
	case CaptchaCategoryLogin:
		return "%s（登录验证码，请完成验证），如非本人操作，请忽略本短信。"
	case CaptchaCategoryResetPassword:
		return "%s （重置密码验证码，请完成重置），如非本人操作，请忽略本短信。"
	default:
		return ""
	}
}

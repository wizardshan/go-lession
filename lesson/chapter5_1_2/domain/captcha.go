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
}

func NewCaptcha(mobile string) *Captcha {
	dom := new(Captcha)
	dom.Mobile = mobile
	return dom
}

// 面向过程或者是面向功能点
func (dom *Captcha) Generate(category string) {
	// 功能点1：生成随机四位数字
	dom.generateCode()
	// 功能点2：生成短信内容
	dom.generateContent(category)
	// 功能点3：生成过期时间
	dom.generateExpireTime()
}

func (dom *Captcha) generateCode() {
	dom.Code = fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

func (dom *Captcha) generateContent(category string) {
	dom.Content = fmt.Sprintf(dom.contentTemplate(category), dom.Code)
}

func (dom *Captcha) generateExpireTime() {
	dom.ExpireTime = time.Now().Add(time.Minute * time.Duration(dom.expireMinutes()))
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

func (dom *Captcha) expireMinutes() int {
	return 5
}



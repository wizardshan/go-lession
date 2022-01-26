package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type contentTemplate interface {
	tpl() string
}

type LoginContentTemplate struct {
}

func (*LoginContentTemplate) tpl() string {
	return "%s（登录验证码，请完成验证），如非本人操作，请忽略本短信。"
}

type ResetPasswordContentTemplate struct {
}

func (*ResetPasswordContentTemplate) tpl() string {
	return "%s （重置密码验证码，请完成重置），如非本人操作，请忽略本短信。"
}

type Captcha struct {
	Code string
	Mobile string
	Content string
	ExpireTime time.Time

	contentTPL contentTemplate
}


type captchaAttributeSetter func(c *Captcha)

func CaptchaLoginContentTPL() captchaAttributeSetter {
	return func(c *Captcha) {
		c.contentTPL = new(LoginContentTemplate)
	}
}

func CaptchaResetPasswordContentTPL() captchaAttributeSetter {
	return func(c *Captcha) {
		c.contentTPL = new(ResetPasswordContentTemplate)
	}
}

func NewCaptcha(mobile string, setters ...captchaAttributeSetter) *Captcha {
	dom := new(Captcha)
	dom.Mobile = mobile

	for _, f := range setters {
		f(dom)
	}

	if dom.contentTPL == nil {
		CaptchaLoginContentTPL()(dom)
	}

	return dom
}

//func NewCaptcha(mobile string, tpl contentTemplate) *Captcha {
//	dom := new(Captcha)
//	dom.Mobile = mobile
//	dom.contentTPL = tpl
//	return dom
//}

func (dom *Captcha) Generate() {
	dom.generateCode()
	dom.generateContent()
	dom.generateExpireTime()
}

func (dom *Captcha) generateCode() {
	dom.Code = fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

func (dom *Captcha) generateContent() {
	dom.Content = fmt.Sprintf(dom.contentTPL.tpl(), dom.Code)
}

func (dom *Captcha) generateExpireTime() {
	dom.ExpireTime = time.Now().Add(time.Minute * time.Duration(dom.expireMinutes()))
}

func (dom *Captcha) expireMinutes() int {
	return 5
}



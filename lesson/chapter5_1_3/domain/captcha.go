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

func NewCaptcha(tpl contentTemplate) *Captcha {
	c := new(Captcha)
	c.contentTPL = tpl
	return c
}

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
	dom.Content = fmt.Sprintf(dom.contentTPL.tpl(), dom.Code)
}

func (dom *Captcha) generateExpireTime() {
	dom.ExpireTime = time.Now().Add(time.Minute * time.Duration(dom.expireMinutes()))
}

func (dom *Captcha) expireMinutes() int {
	return 5
}



package request

import "go-web/lesson/chapter1_3/request/captcha"

type UserLogin struct {
	Mobile
	captcha.Code
}
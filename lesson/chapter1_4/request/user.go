package request

import (
	"errors"
	"go-web/lesson/chapter1_4/request/captcha"
	"go-web/lesson/chapter1_4/request/user"
)

type UserLogin struct {
	MobileS
	captcha.CodeS
}

type UserRegister struct {
	MobileS
	captcha.CodeS
	user.PasswordS
	user.RePasswordS
	user.NicknameS
	user.BirthdayS
}

func (req *UserRegister) Validate() error {
	if req.Password != req.RePassword {
		return errors.New("两次密码不一致")
	}

	return nil
}

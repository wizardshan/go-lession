package request

import (
	"errors"
	"go-web/lesson/chapter4_1/request/captcha"
	"go-web/lesson/chapter4_1/request/user"
)


type UserRegister struct {
	MobileS
	captcha.CodeS
	user.PasswordS
	user.RePasswordS
}

func (req *UserRegister) Validate() error {
	if req.Password != req.RePassword {
		return errors.New("两次密码不一致")
	}

	return nil
}

package request

import "regexp"

type ID struct {
	Val  int `form:"id" valid:"required~id不能为空,int~id格式不正确"`
}

func (req *ID) ID() int {
	return req.Val
}

type Mobile struct {
	Val string `form:"mobile" valid:"required~手机号不能为空,numeric~手机号码应该为数字型,IsMobile~手机号码格式错误"`
}

func (req *Mobile) Mobile() string {
	return req.Val
}

func (req *Mobile) Valid(value string) bool {
	b, _ := regexp.MatchString(`^(1[1-9][0-9]\d{8})$`, value)
	return b
}





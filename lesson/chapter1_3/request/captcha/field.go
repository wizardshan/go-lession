package captcha

type Code struct {
	Val string `form:"code" valid:"required~验证码不能为空,numeric~验证码应该为数字型"`
}

func (req *Code) Code() string {
	return req.Val
}

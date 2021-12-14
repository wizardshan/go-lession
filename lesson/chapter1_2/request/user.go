package request

type User struct {
	Mobile string `form:"mobile" valid:"required~手机号不能为空,numeric~手机号码应该为数字型,IsMobile~手机号码格式错误"`
	Code string `form:"code" valid:"required~验证码不能为空,numeric~验证码应该为数字型"`
}
package request

type User struct {
	Mobile string `form:"mobile" valid:"required~手机号不能为空,numeric~手机号码应该为数字型,IsMobile~手机号码格式错误"`
	Code string `form:"code" valid:"required~验证码不能为空,numeric~验证码应该为数字型"`
}

//type User struct {
//	Mobile string `form:"mobile" valid:"required~手机号不能为空,numeric~手机号码应该为数字型,IsMobile~手机号码格式错误"`
//	Code string `form:"code" valid:"required~验证码不能为空,numeric~验证码应该为数字型"`
//	Password string `form:"password" valid:"required~密码不能为空,stringlength(6|18)~密码6-18个字符"`
//	Nickname string `form:"nickname" valid:"required~昵称不能为空,stringlength(2|10)~昵称2-10个字符"`
//	Birthday time.Time `form:"birthday" valid:"required~生日不能为空" time_format:"2006-01-02"`
//}
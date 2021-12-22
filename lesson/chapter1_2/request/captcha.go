package request

import (
	validator "github.com/asaskevich/govalidator"
	"regexp"
)

type Captcha struct {
	Mobile string `form:"mobile" valid:"required~手机号不能为空,numeric~手机号码应该为数字型,IsMobile~手机号码格式错误"`
}

func init() {
	validator.TagMap["IsMobile"] = func(value string) bool {
		return IsMobile(value)
	}
}

func IsMobile(value string) bool {
	b, _ := regexp.MatchString(`^(1[1-9][0-9]\d{8})$`, value)
	return b
}

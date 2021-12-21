package request

import (
	validator "github.com/asaskevich/govalidator"
	"regexp"
)

func init() {
	m := new(MobileS)
	validator.TagMap["IsMobile"] = func(value string) bool {
		return m.Valid(value)
	}
}

type MobileS struct {
	Mobile string `form:"mobile" valid:"required~手机号不能为空,numeric~手机号码应该为数字型,IsMobile~手机号码格式错误"`
}

func (req *MobileS) Valid(value string) bool {
	b, _ := regexp.MatchString(`^(1[1-9][0-9]\d{8})$`, value)
	return b
}


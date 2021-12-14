package request

import validator "github.com/asaskevich/govalidator"

func init() {
	mobile := new(Mobile)
	validator.TagMap["IsMobile"] = func(value string) bool {
		return mobile.Valid(value)
	}
}
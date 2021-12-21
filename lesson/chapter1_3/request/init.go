package request

import (
	validator "github.com/asaskevich/govalidator"
	"regexp"
)

func init() {
	validator.TagMap["IsMobile"] = func(value string) bool {
		return IsMobile(value)
	}
}

func IsMobile(value string) bool {
	b, _ := regexp.MatchString(`^(1[1-9][0-9]\d{8})$`, value)
	return b
}

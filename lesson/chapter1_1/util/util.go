package util

import (
	"regexp"
)

func IsMobile(mobile string) bool {
	matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile)
	return matched
}

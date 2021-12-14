package util

import (
	"errors"
	"regexp"
)

func IsMobile(mobile string) error {
	if mobile == "" {
		return errors.New("手机号不能为空")
	}

	matched, _ := regexp.MatchString(`^(1[3-9][0-9]\d{8})$`, mobile)
	if !matched {
		return errors.New("手机号格式不正确")
	}
	return nil
}

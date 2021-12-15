package request

import "strings"

type ImageURL string

func (req *ImageURL) Name() string {
	return strings.Replace(string(*req), "http://img.test.com/", "", -1)
}

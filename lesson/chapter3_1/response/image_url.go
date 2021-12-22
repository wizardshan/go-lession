package response

import (
	"fmt"
)

type ImageURL string

func (resp ImageURL) MarshalJSON() ([]byte, error) {
	url := fmt.Sprintf(`"%s"`, "http://img.test.com/"+resp)
	return []byte(url), nil
}

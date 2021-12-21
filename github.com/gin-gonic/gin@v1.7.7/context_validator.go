package gin

import (
	"github.com/asaskevich/govalidator"
)


type Validator interface {
	Validate() error
}

func (c *Context) ShouldB(data interface{}) error {
	if err := c.ShouldBind(data); err != nil {
		return err
	}

	if _, err := govalidator.ValidateStruct(data); err != nil {
		return err
	}

	var v Validator
	var ok bool
	if v, ok = data.(Validator); !ok {
		return nil
	}

	return v.Validate()
}
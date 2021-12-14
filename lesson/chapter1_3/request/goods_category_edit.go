package request

import (
	"go-web/lesson/chapter1_3/request/goodscategory"
)

type GoodsCategoryEdit struct {
	ID
	goodscategory.Name
	goodscategory.Desc
}
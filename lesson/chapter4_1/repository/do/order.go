package do

import (
	"go-web/lesson/chapter4_1/domain"
	"go-web/lesson/chapter4_1/pkg/mapper"
	"time"
)

type Orders []*Order

type Order struct {
	ID         int
	UserID int
	SN       string
	CreateTime time.Time
}

func (do *Order) Mapping() (dom *domain.Order) {
	dom = new(domain.Order)
	mapper.Map(&dom, do)
	return
}

func (do *Orders) Mapping() (dom domain.Orders) {
	dom = domain.Orders{}
	mapper.Map(&dom, do)
	return
}

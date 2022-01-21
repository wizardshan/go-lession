package entity

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

func (do *Order) ToDomain() (dom *domain.Order) {
	dom = new(domain.Order)
	mapper.Map(&dom, do)
	return
}

func (do *Orders) ToDomain() (dom domain.Orders) {
	dom = domain.Orders{}
	mapper.Map(&dom, do)
	return
}

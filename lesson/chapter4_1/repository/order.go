package repository

import (
	"context"
	"fmt"
	"go-web/lesson/chapter4_1/domain"
	"go-web/lesson/chapter4_1/pkg/mapper"
	"go-web/lesson/chapter4_1/repository/do"
	"go-web/lesson/chapter4_1/request"
	"time"
)

type Order struct {}

func NewOrder() *Order {
	repo := new(Order)
	return repo
}

func (repo *Order) Get(ctx context.Context, id int) *domain.Order {
	order := &do.Order{ID: 1, UserID: 1, SN: "123456", CreateTime: time.Now()}

	domain := new(domain.Order)
	mapper.Map(&domain, order)
	return domain
}

func (repo *Order) My(ctx context.Context) []*domain.Order {
	orders := []*do.Order{
		{ID: 1, UserID: 1, SN: "123456", CreateTime: time.Now()},
		{ID: 2, UserID: 1, SN: "987654", CreateTime: time.Now()},
	}

	domains := []*domain.Order{}
	mapper.Map(&domains, orders)
	return domains
}

func (repo *Order) List(ctx context.Context, request *request.OrderList) []*domain.Order {
	fmt.Println(request)

	orders := []*do.Order{
		{ID: 1, UserID: 1, SN: "123456", CreateTime: time.Now()},
		{ID: 2, UserID: 1, SN: "987654", CreateTime: time.Now()},
	}

	domains := []*domain.Order{}
	mapper.Map(&domains, orders)
	return domains
}
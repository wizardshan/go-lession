package repository

import (
	"context"
	"go-web/lesson/chapter4_1/domain"
	"go-web/lesson/chapter4_1/repository/entity"
	"go-web/lesson/chapter4_1/request"
	"time"
)

type Order struct {}

func NewOrder() *Order {
	repo := new(Order)
	return repo
}

func (repo *Order) Get(ctx context.Context, id int) *domain.Order {
	order := &entity.Order{ID: 1, UserID: 1, SN: "123456", CreateTime: time.Now()}
	return order.ToDomain()
}

func (repo *Order) My(ctx context.Context) domain.Orders {
	orders := entity.Orders{
		{ID: 1, UserID: 1, SN: "123456", CreateTime: time.Now()},
		{ID: 2, UserID: 1, SN: "987654", CreateTime: time.Now()},
	}

	return orders.ToDomain()
}

func (repo *Order) List(ctx context.Context, request *request.OrderList) domain.Orders {
	orders := entity.Orders{
		{ID: 1, UserID: 1, SN: "123456", CreateTime: time.Now()},
		{ID: 2, UserID: 1, SN: "987654", CreateTime: time.Now()},
	}

	return orders.ToDomain()
}
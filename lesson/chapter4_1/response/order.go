package response

import (
	"go-web/lesson/chapter4_1/domain"
	"go-web/lesson/chapter4_1/pkg/mapper"
	"time"
)

type Order struct {
	ID         int `json:"id"`
	UserID int `json:"userID"`
	SN       string `json:"sn"`
	CreateTime time.Time `json:"createTime"`
}

type OrderGet struct {
	Order
}

func (resp *OrderGet) Map(order *domain.Order) *OrderGet  {
	mapper.Map(&resp, order)
	return resp
}

type OrderMy struct {
	Order
}

type OrderMys []*OrderMy

func (resp *OrderMys) Map(orders *domain.Orders) *OrderMys {
	mapper.Map(&resp, orders)
	return resp
}

type OrderList struct {
	Order
}

type OrderLists []*OrderList

func (resp *OrderLists) Map(orders *domain.Orders) *OrderLists {
	mapper.Map(&resp, orders)
	return resp
}


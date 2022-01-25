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

func (resp *OrderGet) Mapping(order *domain.Order) *OrderGet  {
	mapper.Map(&resp, order)
	return resp
}

type OrderMy struct {
	Order
}

func (resp *OrderMy) Mapping() {

}

type OrderList struct {
	Order
}

func (resp *OrderList) Mapping() {

}


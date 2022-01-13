package domain

import (
	"go-web/lesson/chapter4_1/pkg/mapper"
	"go-web/lesson/chapter4_1/response"
)

func (dom *Order) MappingGet() (resp *response.OrderGet) {
	resp = new(response.OrderGet)
	mapper.Map(&resp, dom)
	return
}

func (dom *Orders) MappingMy() (resp []*response.OrderMy) {
	resp = []*response.OrderMy{}
	mapper.Map(&resp, dom)
	return
}

func (dom *Orders) MappingList() (resp []*response.OrderList) {
	resp = []*response.OrderList{}
	mapper.Map(&resp, dom)
	return
}
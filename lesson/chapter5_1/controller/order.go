package controller

import "go-web/lesson/chapter5_1/repository"

type Order struct{
	repo *repository.Order
}

func NewOrder(repo *repository.Order) *Order {
	ctr := new(Order)
	ctr.repo = repo
	return ctr
}


package repository

import (
	"context"
	"go-web/lesson/chapter6_1/repository/ent"
)

type Order struct {
	db *ent.Client
}

func NewOrder(db *ent.Client) *Order {
	repo := new(Order)
	repo.db = db
	return repo
}

func (repo *Order) All(ctx context.Context) error {

	return nil
}
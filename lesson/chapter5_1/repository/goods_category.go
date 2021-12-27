package repository

import (
	"context"
	"go-web/lesson/chapter5_1/repository/ent"
	"go-web/lesson/chapter5_1/repository/ent/goodscategory"
)

type GoodsCategory struct {
	db *ent.Client
}

func NewGoodsCategory(db *ent.Client) *GoodsCategory {
	repo := new(GoodsCategory)
	repo.db = db
	return repo
}

func (repo *GoodsCategory) List(ctx context.Context, id int) *ent.GoodsCategory {
	return repo.db.GoodsCategory.Query().WithGoods().Where(goodscategory.ID(id)).FirstX(ctx)
}


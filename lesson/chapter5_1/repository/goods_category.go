package repository

import (
	"context"
	"go-web/lesson/chapter5_1/repository/ent"
	"go-web/lesson/chapter5_1/repository/ent/goodscategory"
)

func (repo *Goods) Category(ctx context.Context, id int) *ent.GoodsCategory {
	return repo.db.GoodsCategory.Query().WithGoods().Where(goodscategory.ID(id)).FirstX(ctx)
}


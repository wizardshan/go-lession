package repository

import (
	"context"
	"go-web/lesson/chapter5_1/domain"
	"go-web/lesson/chapter5_1/pkg/mapper"
	"go-web/lesson/chapter5_1/repository/ent"
	"go-web/lesson/chapter5_1/repository/ent/goodscategory"
)

type Goods struct {
	db *ent.Client
}

func NewGoods(db *ent.Client) *Goods {
	repo := new(Goods)
	repo.db = db
	return repo
}

func (repo *Goods) V1List(ctx context.Context) ([]*ent.Goods, error) {
	return repo.db.Goods.Query().WithCategory().All(ctx)
}

func (repo *Goods) Category(ctx context.Context, id int) (*ent.GoodsCategory, error) {
	return repo.db.GoodsCategory.Query().WithGoods().Where(goodscategory.ID(id)).First(ctx)
}

func (repo *Goods) V2List(ctx context.Context, domainGoods *[]*domain.Goods) error {
	dos, err := repo.db.Goods.Query().WithCategory().All(ctx)
	if err != nil {
		return err
	}

	mapper.Map(domainGoods, dos)
	return nil
}
package repository

import (
	"context"
	"go-web/lesson/chapter5_1/domain"
	"go-web/lesson/chapter5_1/pkg/mapper"
	"go-web/lesson/chapter5_1/repository/ent"
	"go-web/lesson/chapter5_1/repository/ent/goods"
)

type Goods struct {
	db *ent.Client
}

func NewGoods(db *ent.Client) *Goods {
	repo := new(Goods)
	repo.db = db
	return repo
}

func (repo *Goods) V1List(ctx context.Context) []*ent.Goods {
	return repo.db.Goods.Query().WithCategory().AllX(ctx)
}

func (repo *Goods) V2List(ctx context.Context) []*domain.Goods {
	goods := repo.db.Goods.Query().WithCategory().AllX(ctx)

	domains := []*domain.Goods{}
	mapper.Map(&domains, goods)
	return domains
}

func (repo *Goods) Get(ctx context.Context, id int) *domain.Goods {
	query := repo.db.Goods.Query().WithCategory().Where(goods.ID(id))
	return goodsSelect(ctx, query)
}

func (repo *Goods) List(ctx context.Context) []*domain.Goods {
	query := repo.db.Goods.Query().WithCategory()
	return goodsSelectMany(ctx, query)
}

func (repo *Goods) Create(ctx context.Context, domain *domain.Goods) *domain.Goods {
	create := repo.db.Goods.Create().SetName(domain.Name).SetCategoryID(domain.CategoryID)
	return goodsSave(ctx, create)
}

func goodsSelect(ctx context.Context, query *ent.GoodsQuery) *domain.Goods {
	goods := query.FirstX(ctx)
	if goods == nil {
		return nil
	}

	domain := new(domain.Goods)
	mapper.Map(&domain, goods)
	return domain
}

func goodsSelectMany(ctx context.Context, query *ent.GoodsQuery) []*domain.Goods {
	goods := query.AllX(ctx)

	domains := []*domain.Goods{}
	mapper.Map(&domains, goods)
	return domains
}

func goodsSave(ctx context.Context, create *ent.GoodsCreate) *domain.Goods {
	goods := create.SaveX(ctx)
	if goods == nil {
		return nil
	}

	domain := new(domain.Goods)
	mapper.Map(&domain, goods)
	return domain
}

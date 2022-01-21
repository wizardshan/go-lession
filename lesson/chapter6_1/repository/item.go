package repository

import (
	"context"
	"go-web/lesson/chapter6_1/domain"
	"go-web/lesson/chapter6_1/pkg/mapper"
	"go-web/lesson/chapter6_1/repository/ent"
	"go-web/lesson/chapter6_1/repository/ent/item"
)

type Item struct {
	db *ent.Client
}

func NewItem(db *ent.Client) *Item {
	repo := new(Item)
	repo.db = db
	return repo
}

func (repo *Item) V1List(ctx context.Context) ent.Items {
	return repo.db.Item.Query().WithCategory().AllX(ctx)
}

func (repo *Item) V2List(ctx context.Context) []*domain.Item {
	goods := repo.db.Item.Query().WithCategory().AllX(ctx)

	domains := []*domain.Item{}
	mapper.Map(&domains, goods)
	return domains
}

func (repo *Item) Get(ctx context.Context, id int) *domain.Item {
	query := repo.db.Item.Query().WithCategory().Where(item.ID(id))
	return repo.fetch(ctx, query)
}

func (repo *Item) List(ctx context.Context) []*domain.Item {
	query := repo.db.Item.Query().WithCategory()
	return repo.fetchAll(ctx, query)
}

func (repo *Item) Create(ctx context.Context, domain *domain.Item) *domain.Item {
	create := repo.db.Item.Create().SetName(domain.Name).SetCategoryID(domain.CategoryID)
	return repo.save(ctx, create)
}

func (repo *Item) fetch(ctx context.Context, query *ent.ItemQuery) *domain.Item {
	goods := query.FirstX(ctx)
	if goods == nil {
		return nil
	}

	domain := new(domain.Item)
	mapper.Map(&domain, goods)
	return domain
}

func (repo *Item) fetchAll(ctx context.Context, query *ent.ItemQuery) []*domain.Item {
	goods := query.AllX(ctx)

	domains := []*domain.Item{}
	mapper.Map(&domains, goods)
	return domains
}

func (repo *Item) save(ctx context.Context, create *ent.ItemCreate) *domain.Item {
	goods := create.SaveX(ctx)
	if goods == nil {
		return nil
	}

	domain := new(domain.Item)
	mapper.Map(&domain, goods)
	return domain
}

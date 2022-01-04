package ent

import (
	"go-web/lesson/chapter6_1/domain"
	"go-web/lesson/chapter6_1/pkg/mapper"
)

func (_go *Goods) Category() *domain.GoodsCategory {
	if _go.Edges.Category == nil {
		return nil
	}

	domain := new(domain.GoodsCategory)
	mapper.Map(domain, _go.Edges.Category)
	return domain
}
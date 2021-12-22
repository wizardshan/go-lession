package ent

import (
	"go-web/lesson/chapter5_1/domain"
	"go-web/lesson/chapter5_1/pkg/mapper"
)

func (do *Goods) Category() *domain.GoodsCategory {
	if do.Edges.Category == nil {
		return nil
	}

	domain := new(domain.GoodsCategory)
	mapper.Map(domain, do.Edges.Category)
	return domain
}
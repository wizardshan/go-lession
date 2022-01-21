package ent

import (
	"go-web/lesson/chapter6_1/domain"
	"go-web/lesson/chapter6_1/pkg/mapper"
)

func (i *Item) Category() *domain.ItemCategory {
	if i.Edges.Category == nil {
		return nil
	}

	domain := new(domain.ItemCategory)
	mapper.Map(domain, i.Edges.Category)
	return domain
}
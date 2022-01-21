package repository

import (
	"context"
	"fmt"
	"go-web/lesson/chapter6_1/domain"
	"go-web/lesson/chapter6_1/repository/ent"
	"go-web/lesson/chapter6_1/repository/ent/itemcategory"
)

func (repo *Item) Category(ctx context.Context, id int) *ent.ItemCategory {


	var goods []domain.Item
	//执行单条查询
	rows, _ := repo.db.DB().Query("select `id`, `name` from items")
	for rows.Next() {
		var g domain.Item
		rows.Scan(&g.ID, &g.Name)

		fmt.Println(g)
		goods = append(goods, g)
	}

	return repo.db.ItemCategory.Query().WithItems().Where(itemcategory.ID(id)).FirstX(ctx)
}


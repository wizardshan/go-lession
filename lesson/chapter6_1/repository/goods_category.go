package repository

import (
	"context"
	"fmt"
	"go-web/lesson/chapter6_1/domain"
	"go-web/lesson/chapter6_1/repository/ent"
	"go-web/lesson/chapter6_1/repository/ent/goodscategory"
)

func (repo *Goods) Category(ctx context.Context, id int) *ent.GoodsCategory {


	var goods []domain.Goods
	//执行单条查询
	rows, _ := repo.db.DB().Query("select `id`, `name` from goods")
	for rows.Next() {
		var g domain.Goods
		rows.Scan(&g.ID, &g.Name)

		fmt.Println(g)
		goods = append(goods, g)
	}



	return repo.db.GoodsCategory.Query().WithGoods().Where(goodscategory.ID(id)).FirstX(ctx)
}


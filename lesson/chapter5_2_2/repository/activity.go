package repository

import (
	"context"
	"go-web/lesson/chapter5_2_2/domain"
)

type Activity struct {
}

func NewActivity() *Activity {
	repo := new(Activity)
	return repo
}

func (repo *Activity) PrizePool(ctx context.Context) *domain.PrizePool {
	// 模拟数据库查询
	gold := new(domain.Gold)
	gold.Name = "金币"

	point := new(domain.Point)
	point.Name = "积分"

	physical := new(domain.Physical)
	physical.Name = "手机"

	thanks := new(domain.Thanks)
	thanks.Name = "谢谢参与"

	return &domain.PrizePool {
		Prizes: []domain.IPrize {gold, point, physical, thanks},
	}
}

func (repo *Activity) Save(ctx context.Context, prize domain.IPrize) {
	// 模拟数据库保存
}
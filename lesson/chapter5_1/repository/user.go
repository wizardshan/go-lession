package repository

import (
	"context"
	"go-web/lesson/chapter5_1/domain"
	"time"
)

type User struct {
}

func NewUser() *User {
	repo := new(User)
	return repo
}

func (repo *User) Signs(ctx context.Context, days []string) []*domain.UserSign {
	// 模拟数据库查询
	loc, _ := time.LoadLocation("Local")
	signDate, _ := time.ParseInLocation("2006-01-02", days[0], loc)

	return []*domain.UserSign {
		{Id: 1, UserId: 1, Date: signDate},
	}
}
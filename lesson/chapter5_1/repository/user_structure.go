package repository

import (
	"context"
	"go-web/lesson/chapter5_1/domain"
	"go-web/lesson/chapter5_1/pkg/mapper"
	"go-web/lesson/chapter5_1/repository/ent"
)


func userSelect(ctx context.Context, db *ent.Client, query *ent.UserQuery) *domain.User {
	var domain domain.User
	mapper.Map(&domain, db.User.LoadQuery(query).FirstX(ctx))
	return &domain
}

func userSelectMany(ctx context.Context, db *ent.Client, query *ent.UserQuery) []*domain.User {
	var domains []*domain.User
	mapper.Map(&domains, db.User.LoadQuery(query).AllX(ctx))
	return domains
}

func userExist(ctx context.Context, db *ent.Client, query *ent.UserQuery) bool {
	return db.User.LoadQuery(query).ExistX(ctx)
}

func userCount(ctx context.Context, db *ent.Client, query *ent.UserQuery) int {
	return db.User.LoadQuery(query).CountX(ctx)
}


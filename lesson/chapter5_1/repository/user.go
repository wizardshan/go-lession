package repository

import (
	"context"
	"go-web/lesson/chapter5_1/domain"
	"go-web/lesson/chapter5_1/repository/ent"
	"go-web/lesson/chapter5_1/repository/ent/user"
)

type User struct {
	db *ent.Client
}

func NewUser(db *ent.Client) *User {
	repo := new(User)
	repo.db = db
	return repo
}

func (repo *User) Get(ctx context.Context, id int) *domain.User {
	return userSelect(ctx, repo.db, new(ent.UserQuery).Where(user.ID(id)))
}

func (repo *User) List(ctx context.Context) []*domain.User {
	return userSelectMany(ctx, repo.db, new(ent.UserQuery))
}


//func (repo *User) MobileExist(ctx context.Context, mobile string) (bool, error) {
//	return userExist(ctx, repo.db, user.Mobile(mobile))
//}
//
//func (repo *User) NicknameExist(ctx context.Context, nickname string) (bool, error) {
//	return repo.db.User.Query().Where(user.Nickname(nickname)).Exist(ctx)
//}


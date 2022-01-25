package repository

import (
	"context"
	"go-web/lesson/chapter4_1/domain"
	"go-web/lesson/chapter4_1/repository/entity"
)

type User struct {

}

func NewUser() *User {
	repo := new(User)
	return repo
}

func (repo *User) Register(ctx context.Context, userDomain *domain.User) *domain.User {

	userEntity := new(entity.User)
	userEntity.ID = 1
	userEntity.Mobile = userDomain.Mobile
	userEntity.Password = userDomain.Password
	userEntity.HeaderImage = userDomain.HeaderImage
	userEntity.CreateTime = userDomain.CreateTime
	// entity对象入库

	return userEntity.ToDomain()
}


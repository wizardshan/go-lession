package repository

import (
	"context"
	"go-web/lesson/chapter4_1/domain"
	"go-web/lesson/chapter4_1/pkg/mapper"
	"go-web/lesson/chapter4_1/repository/do"
	"go-web/lesson/chapter4_1/repository/entity"
	"time"
)

type User struct {

}

func NewUser() *User {
	repo := new(User)
	return repo
}

func (repo *User) V1List(ctx context.Context) []*domain.User {
	users := []*entity.User{
		{ID: 1, Nickname: "用户昵称", HeaderImage: "header.jpg", Money: 12345, CreateTime: time.Now()},
		{ID: 1, Nickname: "用户昵称", HeaderImage: "header.jpg", Money: 12345, CreateTime: time.Now()},
	}

	domains := []*domain.User{}
	mapper.Map(&domains, users)
	return domains
}

func (repo *User) V2List(ctx context.Context) []*domain.User {
	users := []*entity.User{
		{ID: 1, Nickname: "用户昵称", HeaderImage: "header.jpg", Money: 12345, CreateTime: time.Now()},
		{ID: 1, Nickname: "用户昵称", HeaderImage: "header.jpg", Money: 12345, CreateTime: time.Now()},
	}

	domains := []*domain.User{}
	mapper.Map(&domains, users)
	return domains
}

func (repo *User) Get(ctx context.Context) *domain.User {
	user := &entity.User{ID: 1, Nickname: "用户昵称", HeaderImage: "header.jpg", Money: 12345, CreateTime: time.Now()}

	domain := new(domain.User)
	mapper.Map(&domain, user)
	return domain
}

func (repo *User) Register(ctx context.Context, user *domain.User) *domain.User {

	do := new(entity.User)
	do.ID = 1
	do.Mobile = user.Mobile
	do.Password = user.Password
	do.HeaderImage = user.HeaderImage
	do.CreateTime = user.CreateTime
	// do对象入库

	domain := new(domain.User)
	mapper.Map(&domain, do)
	return domain
}


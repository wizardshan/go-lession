package repository

import (
	"context"
	"go-web/lesson/chapter4_1/domain"
	"go-web/lesson/chapter4_1/pkg/mapper"
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

func (repo *User) Register(ctx context.Context, userDomain *domain.User) *domain.User {

	userEntity := new(entity.User)
	userEntity.ID = 1
	userEntity.Mobile = userDomain.Mobile
	userEntity.Password = userDomain.Password
	userEntity.HeaderImage = userDomain.HeaderImage
	userEntity.CreateTime = userDomain.CreateTime
	// entity对象入库

	mapper.Map(&userEntity, userDomain)
	return userDomain
}


package entity

import (
	"go-web/lesson/chapter4_1/domain"
	"go-web/lesson/chapter4_1/pkg/mapper"
	"time"
)

type Users []*User

type User struct {
	ID          int
	Mobile      string
	Password    string
	Nickname    string
	HeaderImage string
	Money       int
	CreateTime time.Time
}

func (ent *User) ToDomain() *domain.User {
	dom := new(domain.User)
	mapper.Map(&dom, ent)
	return dom
}


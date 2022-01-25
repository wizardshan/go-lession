package response

import (
	"go-web/lesson/chapter4_1/domain"
	"go-web/lesson/chapter4_1/pkg/mapper"
	"time"
)

type User struct {
	ID          int      `json:"id"`
	Nickname    string   `json:"nickname"`
	Mobile      string `json:"mobile"`
	HeaderImage string `json:"headerImage"`
	Money       int    `json:"money"`
	CreateTime time.Time `json:"createTime"`
}

type UserRegister struct {
	User
}

func (resp *UserRegister) Map(user *domain.User) *UserRegister {
	mapper.Map(&resp, user)
	return resp
}


package response

import (
	"go-web/lesson/chapter3_1/domain"
	"go-web/lesson/chapter3_1/pkg/mapper"
)

type User struct {
	ID          int      `json:"id"`
	Nickname    string   `json:"nickname"`
	HeaderImage ImageURL `json:"headerImage"`
	Money       Money    `json:"money"`
	CreateTime DateTime `json:"createTime"`
}

type UserInfo struct {
	User
}

func (resp *UserInfo) Map(user *domain.User) *UserInfo {
	mapper.Map(resp, user)
	return resp
}

type UserUpdate struct {
	User
}

func (resp *UserUpdate) Map(user *domain.User) *UserUpdate {
	mapper.Map(resp, user)
	return resp
}


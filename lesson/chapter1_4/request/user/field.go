package user

import "time"

type PasswordS struct {
	Password string `form:"password" valid:"required~密码不能为空,stringlength(6|18)~密码6-18个字符"`
}

type RePasswordS struct {
	RePassword string `form:"rePassword" valid:"required~重复密码不能为空,stringlength(6|18)~重复密码6-18个字符"`
}

type NicknameS struct {
	Nickname string `form:"nickname" valid:"required~昵称不能为空,stringlength(2|10)~昵称2-10个字符"`
}


type BirthdayS struct {
	Birthday time.Time `form:"birthday" valid:"required~生日不能为空" time_format:"2006-01-02"`
}
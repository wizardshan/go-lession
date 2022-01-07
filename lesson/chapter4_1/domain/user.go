package domain

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type NicknameS struct {}

func (dom *NicknameS) generate(str string) string {
	return str[0:4] + "****" + str[8:11]
}

type PasswordS struct {}

func (dom *PasswordS) generate(str string) string {
	str += dom.salt()
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (dom *PasswordS) salt() string {
	return "~!@#abc123"
}

type User struct {
	ID          int
	Mobile      string
	Nickname    string
	NicknameS    *NicknameS
	Password    string
	PasswordS    *PasswordS
	HeaderImage string
	Money       int
	CreateTime time.Time
}

func (dom *User) Register(mobile string, password string) {
	dom.Mobile = mobile
	dom.Nickname = dom.NicknameS.generate(mobile)
	dom.Password = dom.PasswordS.generate(password)
	dom.CreateTime = time.Now()
	dom.HeaderImage = dom.DefaultHeaderImage()
}

func (dom *User) DefaultHeaderImage() string {
	return "header.jpg"
}





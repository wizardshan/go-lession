package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter3_1/domain"
	"go-web/lesson/chapter3_1/request"
	"go-web/lesson/chapter3_1/response"
	"net/http"
	"time"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Cash(c *gin.Context) {
	request := new(request.UserCash)
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": request})
}

func (ctr *User) Info(c *gin.Context) {
	user := &domain.User{ID: 1, Nickname: "用户昵称", HeaderImage: "header.jpg", Money: 12345, CreateTime: time.Now()}

	response := new(response.UserInfo)
	c.JSON(http.StatusOK, gin.H{"data": response.Map(user)})
}

func (ctr *User) Update(c *gin.Context) {
	request := new(request.UserUpdate)
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &domain.User{ID: 1, Nickname: "用户昵称", HeaderImage: request.HeaderImage.Name(), Money: 12345}

	response := new(response.UserUpdate)
	c.JSON(http.StatusOK, gin.H{"data": response.Map(user)})
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go-web/lesson/chapter3_1/domain"
	"go-web/lesson/chapter3_1/request"
	"go-web/lesson/chapter3_1/response"
	"net/http"
)

type User struct{}

func (ctr *User) Cash(c *gin.Context) {
	var request request.UserCash
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(request.Money.Value())
}

func (ctr *User) Info(c *gin.Context) {
	user := &domain.User{ID: 1, Nickname: "用户昵称", HeaderImage: "header.jpg", Money: 12345}

	var response response.User
	copier.Copy(&response, user)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *User) Update(c *gin.Context) {
	var request request.UserUpdate
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(request.HeaderImage.Name())
	user := &domain.User{ID: 1, Nickname: "用户昵称", HeaderImage: request.HeaderImage.Name(), Money: 12345}
	var response response.User
	copier.Copy(&response, user)
	c.JSON(http.StatusOK, gin.H{"data": response})
}

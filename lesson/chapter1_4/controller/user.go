package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter1_4/request"
	"net/http"
)

type User struct{}

func (ctr *User) Login(c *gin.Context) {
	var request request.UserLogin
	if err := c.ShouldB(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(request.Mobile, request.Code)
}

func (ctr *User) Register(c *gin.Context) {
	var request request.UserRegister
	if err := c.ShouldB(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(request.Mobile, request.Code, request.Password, request.RePassword, request.Nickname, request.Birthday)
}

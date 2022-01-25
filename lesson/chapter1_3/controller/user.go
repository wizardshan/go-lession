package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter1_3/request"
	"net/http"
)

type User struct{}

func NewUser() *User {
	ctr := new(User)
	return ctr
}

func (ctr *User) Register(c *gin.Context) {
	request := new(request.UserRegister)
	if err := c.ShouldB(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"data": request})
}

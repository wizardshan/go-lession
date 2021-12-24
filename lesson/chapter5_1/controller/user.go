package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_1/repository"
	"net/http"
)

type User struct{
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": nil})
}

func (ctr *User) Get(c *gin.Context) {
	user := ctr.repo.Get(c.Request.Context(), 1)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctr *User) List(c *gin.Context) {
	users := ctr.repo.List(c.Request.Context())
	c.JSON(http.StatusOK, gin.H{"data": users})
}


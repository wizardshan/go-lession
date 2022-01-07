package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter4_1/domain"
	"go-web/lesson/chapter4_1/pkg/mapper"
	"go-web/lesson/chapter4_1/repository"
	"go-web/lesson/chapter4_1/request"
	"go-web/lesson/chapter4_1/response"
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
	var request request.UserRegister
	if err := c.ShouldB(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := new(domain.User)
	user.Register(request.Mobile, request.Password)

	data := ctr.repo.Register(c.Request.Context(), user)

	response := new(response.User)
	mapper.Map(&response, data)
	c.JSON(http.StatusOK, gin.H{"data": response})
}
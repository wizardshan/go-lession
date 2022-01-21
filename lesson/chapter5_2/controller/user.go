package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_2/repository"
	"net/http"
	"time"
)

type User struct{
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	ctr := new(User)
	ctr.repo = repo
	return ctr
}

func (ctr *User) Signs(c *gin.Context) {

	now := time.Now()
	index := int(now.Weekday())
	days := make([]string, 7)
	for i := 0; i < 7; i++ {
		day := now.AddDate(0, 0, i - index).Format("2006-01-02")
		days[i] = day
	}
	// dates [2022-01-02 2022-01-03 2022-01-04 2022-01-05 2022-01-06 2022-01-07 2022-01-08]

	signs := ctr.repo.Signs(c.Request.Context(), days)
	// signs [&{Id:1 UserId:1 Date:2022-01-02 00:00:00 +0800 CST}]

	response := make(map[string]bool)
	for _, s := range signs {
		k := s.Date.Format("2006-01-02")
		response[k] = true
	}
	// response [2022-01-02:true]

	for i := 0; i < 7; i++ {
		k := days[i]
		if _, ok := response[k]; ok {
			continue
		}
		response[k] = false
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
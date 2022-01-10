package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_2_1/repository"
	"net/http"
)

type Activity struct{
	repo *repository.Activity
}

func NewActivity(repo *repository.Activity) *Activity {
	ctr := new(Activity)
	ctr.repo = repo
	return ctr
}

func (ctr *Activity) Lottery(c *gin.Context) {

	pool := ctr.repo.PrizePool(c.Request.Context())
	prize, err := pool.Lottery()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctr.repo.Save(c.Request.Context(), prize)

	c.JSON(http.StatusOK, gin.H{"data": prize})
}

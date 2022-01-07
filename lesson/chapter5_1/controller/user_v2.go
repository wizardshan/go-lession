package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_1/domain"
	"go-web/lesson/chapter5_1/pkg/date"
	"net/http"
)


func (ctr *User) SignsV2(c *gin.Context) {

	var week date.Week
	dates := week.CurrentDates()

	var signs domain.UserSigns = ctr.repo.Signs(c.Request.Context(), dates)

	response := ctr.transformSignsV2(signs.TransformMap(), dates)

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *User) transformSignsV2(signs map[string]*domain.UserSign, dates []string) map[string]bool {

	signed := true //消除隐性知识
	unSign := false

	signsMap := make(map[string]bool)
	len := len(dates)
	for i := 0; i < len; i++ {
		k := dates[i]
		if _, ok := signs[k]; ok {
			signsMap[k] = signed
		}
		signsMap[k] = unSign
	}
	return signsMap
}
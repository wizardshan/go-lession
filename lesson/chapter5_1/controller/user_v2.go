package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_1/domain"
	"go-web/lesson/chapter5_1/pkg/date"
	"net/http"
)


func (ctr *User) SignsV2(c *gin.Context) {

	var week date.Week
	days := week.AllDays()

	var signs domain.UserSigns = ctr.repo.Signs(c.Request.Context(), days)

	response := ctr.transformSignsV2(signs.TransformMap(), days)

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *User) transformSignsV2(signs map[string]*domain.UserSign, days []string) map[string]bool {

	signed := true //消除隐性知识
	unSign := false

	signsMap := make(map[string]bool)
	len := len(days)
	for i := 0; i < len; i++ {
		k := days[i]
		if _, ok := signs[k]; ok {
			signsMap[k] = signed
			continue
		}
		signsMap[k] = unSign
	}
	return signsMap
}
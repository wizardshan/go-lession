package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_2/domain"
	"go-web/lesson/chapter5_2/pkg/date"
	"net/http"
)


func (ctr *User) SignsV3(c *gin.Context) {

	periodType, _ := c.GetQuery("periodType")
	period := date.NewPeriod(periodType)
	days := period.AllDays()

	var signs domain.UserSigns = ctr.repo.Signs(c.Request.Context(), days)

	response := ctr.transformSignsV3(signs.TransformMap(), days)

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *User) transformSignsV3(signs map[string]*domain.UserSign, dates []string) map[string]bool {

	signed := true //消除隐性知识
	unSign := false

	signsMap := make(map[string]bool)
	len := len(dates)
	for i := 0; i < len; i++ {
		k := dates[i]
		if _, ok := signs[k]; ok {
			signsMap[k] = signed
			continue
		}
		signsMap[k] = unSign
	}
	return signsMap
}
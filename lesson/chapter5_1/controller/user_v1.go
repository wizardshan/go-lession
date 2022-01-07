package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/lesson/chapter5_1/domain"
	"net/http"
	"time"
)


func (ctr *User) SignsV1(c *gin.Context) {

	dates := ctr.getCurrentWeekDates()

	signs := ctr.repo.Signs(c.Request.Context(), dates)

	response := ctr.transformSigns(signs, dates)

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (ctr *User) getCurrentWeekDates() []string {
	now := time.Now()
	index := int(now.Weekday())

	weekdays := 7 // 每周七天 用变量名自解释，消除隐性知识
	dates := make([]string, weekdays)
	for i := 0; i < weekdays; i++ {
		dates[i] = now.AddDate(0, 0, i - index).Format(ctr.getDateFormatLayout())
	}

	return dates
}

func (ctr *User) transformSigns(signs []*domain.UserSign, dates []string) map[string]bool {

	signed := true //消除隐性知识
	unSign := false

	signsMap := make(map[string]bool)
	for _, s := range signs {
		k := s.Date.Format(ctr.getDateFormatLayout())
		signsMap[k] = signed
	}

	len := len(dates)
	for i := 0; i < len; i++ {
		if _, ok := signsMap[dates[i]]; ok {
			continue
		}
		signsMap[dates[i]] = unSign
	}
	return signsMap
}

// 函数名自解释，消除隐性知识
func (ctr *User) getDateFormatLayout() string {
	return "2006-01-02"
}
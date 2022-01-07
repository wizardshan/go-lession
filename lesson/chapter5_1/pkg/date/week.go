package date

import (
	"go-web/lesson/chapter5_1/pkg/timelayout"
	"time"
)

// 用变量名自解释，消除隐性知识
var Weekdays = 7

type Week struct {
}

func (w *Week) CurrentDates() []string {
	now := time.Now()
	index := int(now.Weekday())

	dates := make([]string, Weekdays)
	for i := 0; i < Weekdays; i++ {
		dates[i] = now.AddDate(0, 0, i - index).Format(timelayout.Ymd)
	}

	return dates
}



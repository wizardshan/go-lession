package date

import (
	"time"
)

// 用变量名自解释，消除隐性知识
var Weekdays = 7

type Week struct {
}

func (d *Week) AllDays() []string {
	now := time.Now()
	index := int(now.Weekday())

	days := make([]string, Weekdays)
	for i := 0; i < Weekdays; i++ {
		days[i] = now.AddDate(0, 0, i - index).Format(LayoutYmd)
	}

	return days
}



package date

import (
	"time"
)

type Month struct {
}

func (d *Month) AllDays() []string {
	now := time.Now()
	first := d.FirstDate(now)
	lastDay := d.LastDate(now).Format(LayoutYmd)

	var days []string
	var i int
	for {
		day := first.AddDate(0, 0, i).Format(LayoutYmd)
		days = append(days, day)
		i++

		if day == lastDay {
			break
		}
	}

	return days
}


func (d *Month) FirstDate(t time.Time) time.Time {
	t = t.AddDate(0, 0, -t.Day() + 1)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func (d *Month) LastDate(t time.Time) time.Time {
	return d.FirstDate(t).AddDate(0, 1, -1)
}
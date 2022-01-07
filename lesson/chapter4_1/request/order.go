package request

import "time"

type OrderGet struct {
	IDS
}

type OrderList struct {
	SN string `form:"sn"`
	CreateTimeStart time.Time `form:"createTimeStart" time_format:"2006-01-02 15:04:05"`
	CreateTimeEnd time.Time `form:"createTimeEnd" time_format:"2006-01-02 15:04:05"`
}


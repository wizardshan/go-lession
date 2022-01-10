package domain

import (
	"go-web/lesson/chapter5_1/pkg/date"
	"time"
)

type UserSign struct {
	Id int
	UserId int
	Date time.Time
}

type UserSigns []*UserSign

func (dom UserSigns) TransformMap() map[string]*UserSign {
	signsMap := make(map[string]*UserSign)
	for _, s := range dom {
		k := s.Date.Format(date.LayoutYmd)
		signsMap[k] = s
	}

	return signsMap
}


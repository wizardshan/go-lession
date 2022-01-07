package domain

import (
	"go-web/lesson/chapter5_1/pkg/timelayout"
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
		k := s.Date.Format(timelayout.Ymd)
		signsMap[k] = s
	}

	return signsMap
}


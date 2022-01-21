package domain

import (
	"errors"
	"time"
)

type IPrizeLimiter interface {
	CheckLimit() error
}

type IPrize interface {
	Award()
}

type Prize struct {
	Id int
	Number int
	Percent time.Time
	Name string
	TotalLimit int
}

type PrizeLimiter struct {
	Prize
}

func (dom *PrizeLimiter) CheckLimit() error {
	if dom.TotalLimit == 0 {
		return errors.New("奖品总数限制已达到")
	}

	return nil
}

type Gold struct {
	PrizeLimiter
}

func (dom *Gold) Award() {
	// 调用金币系统发放金币
}

type Physical struct {
	PrizeLimiter
}

func (dom *Physical) Award() {
	// 调用奖品系统发放实物奖品
}

type Point struct {
	PrizeLimiter
}

func (dom *Point) Award() {
	// 调用积分系统发放实物奖品
}

type Thanks struct {
	Prize
}

func (dom *Thanks) Award() {
	// 调用奖品系统发放谢谢参与
}
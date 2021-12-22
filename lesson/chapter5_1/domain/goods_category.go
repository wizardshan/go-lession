package domain

import "time"

type GoodsCategory struct {
	ID         int
	Name       string
	CreateTime time.Time
}

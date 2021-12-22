package domain

import "time"

type Goods struct {
	ID         int
	CategoryID int
	Category *GoodsCategory
	Name       string
	CreateTime time.Time
}

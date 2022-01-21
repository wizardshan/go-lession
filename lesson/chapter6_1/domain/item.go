package domain

import "time"

type Item struct {
	ID         int
	CategoryID int
	Category *ItemCategory
	Name       string
	CreateTime time.Time
}


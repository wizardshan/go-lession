package domain

import "time"

type Order struct {
	ID         int
	CategoryID int
	Name       string
	CreateTime time.Time
}

package domain

import (
	"time"
)

type Orders []*Order

type Order struct {
	ID         int
	UserID int
	SN       string
	CreateTime time.Time
}

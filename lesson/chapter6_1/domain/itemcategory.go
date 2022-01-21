package domain

import (
	"time"
)

type ItemCategory struct {
	ID         int
	Name       string
	CreateTime time.Time
}

package domain

import "time"

type User struct {
	ID         int
	Mobile string
	Nickname       string
	CreateTime time.Time
}

package domain

import "time"

type User struct {
	ID          int
	Nickname    string
	HeaderImage string
	Money       int
	CreateTime time.Time
}

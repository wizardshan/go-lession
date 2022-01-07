package do

import "time"

type Order struct {
	ID         int
	UserID int
	SN       string
	CreateTime time.Time
}

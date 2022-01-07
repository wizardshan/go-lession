package do

import "time"

type User struct {
	ID          int
	Mobile      string
	Password    string
	Nickname    string
	HeaderImage string
	Money       int
	CreateTime time.Time
}


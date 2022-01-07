package response

import "time"

type User struct {
	ID          int      `json:"id"`
	Nickname    string   `json:"nickname"`
	Mobile      string `json:"mobile"`
	HeaderImage string `json:"headerImage"`
	Money       int    `json:"money"`
	CreateTime time.Time `json:"createTime"`
}

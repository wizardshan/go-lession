package response

import "time"

type Order struct {
	ID         int `json:"id"`
	UserID int `json:"userID"`
	SN       string `json:"sn"`
	CreateTime time.Time `json:"createTime"`
}
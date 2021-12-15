package response

type User struct {
	ID int `json:"id"`
	Nickname string `json:"nickname"`
	HeaderImage ImageURL `json:"headerImage"`
	Money Money `json:"money"`
}
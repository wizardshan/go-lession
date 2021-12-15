package request

type UserCash struct {
	Money Money `form:"money"`
}

type UserUpdate struct {
	HeaderImage ImageURL `form:"headerImage"`
}
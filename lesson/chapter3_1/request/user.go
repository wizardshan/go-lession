package request

type UserCash struct {
	Money MoneyS `form:"money"`
}

type UserUpdate struct {
	HeaderImage ImageURL `form:"headerImage"`
}

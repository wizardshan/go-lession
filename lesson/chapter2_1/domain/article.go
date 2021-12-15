package domain

type Article struct {
	ID int
	Title string
	ImageUrl string
	Content string

	Author *Author
}
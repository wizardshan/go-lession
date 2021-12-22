package domain

type Article struct {
	ID       int
	Title    string
	ImageURL string
	Content  string

	Author *Author
}

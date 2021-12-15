package response

type Author struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type ArticleListAuthor struct {
	Author
	Desc string `json:"desc,omitempty" copier:"-"`
}
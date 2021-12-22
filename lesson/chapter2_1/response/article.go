package response

type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ImageUrl string `json:"imageUrl"`
	Content  string `json:"content"`
}

type ArticleListAuthor struct {
	Author
	Desc string `json:"desc,omitempty" copier:"-"`
}

type ArticleList struct {
	Article
	Content           string `json:"content,omitempty" copier:"-"`
	ArticleListAuthor `json:"author"`
}

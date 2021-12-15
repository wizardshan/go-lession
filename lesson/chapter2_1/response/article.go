package response

type Article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	ImageUrl string `json:"imageUrl"`
	Content string `json:"content"`

}

type ArticleList struct {
	Article
	Content string `json:"content,omitempty" copier:"-"`
	ArticleListAuthor `json:"author"`
}

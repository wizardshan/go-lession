package response

import (
	"go-web/lesson/chapter2_1/domain"
	"go-web/lesson/chapter2_1/pkg/mapper"
)

type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"imageURL"`
	Content  string `json:"content"`
}

type ArticleListAuthor struct {
	Author
	Desc string `json:"desc,omitempty" copier:"-"`
}

type ArticleLists []*ArticleList

func (resp *ArticleLists) Map(articles []*domain.Article) *ArticleLists {
	mapper.Map(resp, articles)
	return resp
}

type ArticleList struct {
	Article
	Content           string `json:"content,omitempty" copier:"-"`
	ArticleListAuthor `json:"author"`
}

type ArticleDetail struct {
	Article
}

func (resp *ArticleDetail) Map(article *domain.Article) *ArticleDetail {
	mapper.Map(resp, article)
	return resp
}
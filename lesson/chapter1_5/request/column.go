package request

import (
	"go-web/lesson/chapter1_5/request/article"
	"go-web/lesson/chapter1_5/request/column"
)

type ColumnCreateArticle struct {
	IDS
	article.TitleS
}

type ColumnCreate struct {
	column.TitleS
	Article  *ColumnCreateArticle  `form:"article"`
	Articles []ColumnCreateArticle `form:"articles"`
}

type ColumnDetail struct {
	IDS
	ArticleIsImportant *bool `form:"articleIsImportant"`
}

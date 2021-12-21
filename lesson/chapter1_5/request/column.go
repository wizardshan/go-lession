package request

import (
	"go-web/lesson/chapter1_5/request/article"
	"go-web/lesson/chapter1_5/request/column"
)

type ColumnArticleCreate struct {
	IDS
	article.TitleS
}

type ColumnCreate struct {
	column.TitleS
	Article *ColumnArticleCreate `form:"article"`
	Articles []ColumnArticleCreate `form:"articles"`
}

type ColumnDetail struct {
	IDS
}
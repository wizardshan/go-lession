package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Goods struct {
	ent.Schema
}

func (Goods) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id").Optional(),
		field.String("name"),
	}
}

func (Goods) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Goods) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", GoodsCategory.Type).
			Ref("goods").Field("category_id").
			Unique(),
	}
}

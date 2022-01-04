package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type GoodsCategory struct {
	ent.Schema
}

func (GoodsCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (GoodsCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (GoodsCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("goods", Goods.Type),
	}
}

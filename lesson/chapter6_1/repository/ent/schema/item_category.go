package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type ItemCategory struct {
	ent.Schema
}

func (ItemCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (ItemCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (ItemCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", Item.Type),
	}
}

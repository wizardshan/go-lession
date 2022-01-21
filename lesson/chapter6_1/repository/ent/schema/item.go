package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Item struct {
	ent.Schema
}

func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id").Optional(),
		field.String("name"),
	}
}

func (Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", ItemCategory.Type).
			Ref("items").Field("category_id").
			Unique(),
	}
}

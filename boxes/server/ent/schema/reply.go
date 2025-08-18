package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Reply struct {
	ent.Schema
}

func (Reply) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content").NotEmpty(),
	}
}

func (Reply) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("replies").
			Unique().
			Required(),
		edge.From("post", Post.Type).
			Ref("replies").
			Unique().
			Required(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DiscussionBoard struct {
	ent.Schema
}

func (DiscussionBoard) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

func (DiscussionBoard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.From("course", Course.Type).
			Ref("discussion_board").
			Unique().
			Required(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type DiscussionBoard struct {
	ent.Schema
}

func (DiscussionBoard) Fields() []ent.Field {
	return nil
}

func (DiscussionBoard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.From("course", Course.Type).Ref("board").Unique().Required(),
	}
}

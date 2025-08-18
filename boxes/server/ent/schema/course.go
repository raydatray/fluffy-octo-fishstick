package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Course struct {
	ent.Schema
}

func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique(),
	}
}

func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("discussion_board", DiscussionBoard.Type).Unique(),
		edge.To("sections", CourseSection.Type),
	}
}

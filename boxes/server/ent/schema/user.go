package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("password"),
		field.Enum("role").
			Values("PROFESSOR", "TA", "STUDENT").
			Default("STUDENT"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("replies", Reply.Type),
		edge.To("teaching_sections", CourseSection.Type),
		edge.To("assisting_sections", CourseSection.Type),
		edge.To("enrolled_sections", CourseSection.Type),
	}
}

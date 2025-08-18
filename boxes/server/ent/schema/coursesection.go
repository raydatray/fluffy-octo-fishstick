package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CourseSection struct {
	ent.Schema
}

func (CourseSection) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number"),
	}
}

func (CourseSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).
			Ref("sections").
			Unique().
			Required(),
		edge.From("professors", User.Type).
			Ref("teaching_sections").
			Required(),
		edge.From("teaching_assistants", User.Type).
			Ref("assisting_sections"),
		edge.From("students", User.Type).
			Ref("enrolled_sections"),
	}
}

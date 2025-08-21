package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").NotEmpty(),
		field.String("middle_name").Optional(),
		field.String("last_name").NotEmpty(),
		field.String("email").Unique(),
		field.String("password"),
		field.Enum("role").
			Values("PROFESSOR", "TA", "CA", "STUDENT").
			Default("STUDENT"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("replies", Reply.Type),
		edge.To("teaching_sections", CourseSection.Type),
		edge.To("teaching_assistant_sections", CourseSection.Type),
		edge.To("course_assistant_sections", CourseSection.Type),
		edge.To("enrolled_sections", CourseSection.Type),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

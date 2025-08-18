package schema

import "entgo.io/ent"

// CourseSection holds the schema definition for the CourseSection entity.
type CourseSection struct {
	ent.Schema
}

// Fields of the CourseSection.
func (CourseSection) Fields() []ent.Field {
	return nil
}

// Edges of the CourseSection.
func (CourseSection) Edges() []ent.Edge {
	return nil
}

package schema

import "entgo.io/ent"

// DiscussionBoard holds the schema definition for the DiscussionBoard entity.
type DiscussionBoard struct {
	ent.Schema
}

// Fields of the DiscussionBoard.
func (DiscussionBoard) Fields() []ent.Field {
	return nil
}

// Edges of the DiscussionBoard.
func (DiscussionBoard) Edges() []ent.Edge {
	return nil
}

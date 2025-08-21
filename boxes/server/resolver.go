package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/raydatray/fluffy-octo-fishstick/boxes/server/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	entClient *ent.Client
}

func NewSchema(entClient *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{entClient},
	})
}

package main

import (
	"context"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/raydatray/fluffy-octo-fishstick/boxes/server"
	"github.com/raydatray/fluffy-octo-fishstick/boxes/server/ent"
	"github.com/raydatray/fluffy-octo-fishstick/boxes/server/ent/migrate"
)

func main() {
	entClient, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	if err := entClient.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running ent client migration", err)
	}

	srv := handler.NewDefaultServer(server.NewSchema(entClient))
	http.Handle("/", playground.Handler("Fluffy Octo Fishstick", "/query"))
	http.Handle("/query", srv)

	log.Println("listening on :8081")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}

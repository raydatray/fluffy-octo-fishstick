package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/raydatray/fluffy-octo-fishstick/boxes/server/ent"
	"github.com/raydatray/fluffy-octo-fishstick/boxes/server/handlers"
)

type Server struct {
	client      *ent.Client
	userHandler *handlers.UserHandler
}

func NewServer() *Server {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &Server{
		client:      client,
		userHandler: handlers.NewUserHandler(client),
	}
}

func (s *Server) Close() error {
	return s.client.Close()
}

// Health check endpoint
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func main() {
	server := NewServer()
	defer server.Close()

	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", server.handleHealth)

	// User endpoints
	mux.HandleFunc("/api/users", server.userHandler.HandleUsers)
	mux.HandleFunc("/api/users/", server.userHandler.HandleUserByID)

	// Course endpoints will be added later

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

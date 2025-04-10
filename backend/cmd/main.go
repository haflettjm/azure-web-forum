package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/haflettjm/azure-web-forum/backend/internal/db"
	"github.com/haflettjm/azure-web-forum/backend/internal/handlers"
)

func server(route string) {
	switch route {
	case "/users":
		http.HandleFunc("/users", handlers.UserHandler)
	case "/posts":
		http.HandleFunc("/posts", handlers.PostHandler)
	default:
		http.HandleFunc("/", handlers.HomeHandler)
	}
}

func main() {
	// Load environment variable
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Initialize the DB connection
	if err := db.InitDB(connStr); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Set up routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/users", handlers.UserHandler)
	http.HandleFunc("/posts", handlers.PostHandler)

	// Start the server
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

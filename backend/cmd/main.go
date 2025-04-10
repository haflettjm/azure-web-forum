package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yourusername/azure-web-forum/backend/internal/handlers"
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
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Fatal("DATABASE_URL enviornment variable not set")
	}
	// Define Handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/users", handlers.UserHandler)
	http.HandleFunc("/posts", handlers.PostHandler)

	// Start Server
	fmt.Println("Starting Server on :8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

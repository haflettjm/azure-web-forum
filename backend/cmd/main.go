package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/yourusername/azure-web-forum/backend/internal/handlers"

)


func main() {
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

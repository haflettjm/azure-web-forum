package main

import (
  "fmt"
  "log"
  "net/http"
)

func homeHandler (w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "Welcome to azure web forum backend!")
} 


func main() {
  http.HandleFunc("/", homeHandler)

  fmt.Println("Starting server on :8080")

  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatalf("Server failed to start: %v", err)
  }
}

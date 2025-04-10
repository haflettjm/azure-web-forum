package handlers

import (
    "fmt"
    db "github.com/haflettjm/azure-web-forum/backend/internal/db"
    "net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintln(w, "get request recieved")
        users := getUsers
        
    case http.MethodPost:
        fmt.Fprintln(w, "Create a new user")
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
  }


  func getUsers(w http.ResponseWriter, r *http.Request)  {
      switch
  }
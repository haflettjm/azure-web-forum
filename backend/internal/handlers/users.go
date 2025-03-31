package handlers

import (
    "fmt"
    "net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintln(w, "List of users")
    case http.MethodPost:
        fmt.Fprintln(w, "Create a new user")
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

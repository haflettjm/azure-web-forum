package handlers

import (
    "fmt"
    "net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintln(w, "List of posts")
    case http.MethodPost:
        fmt.Fprintln(w, "Create a new post")
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

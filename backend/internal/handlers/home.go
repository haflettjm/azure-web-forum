package handlers

import(
  "fmt"
  "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "welcome to AAA Web Forum Backend!")
}


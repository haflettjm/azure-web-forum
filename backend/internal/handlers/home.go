package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/haflettjm/azure-web-forum/backend/internal/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to AAA Web Forum Backend!")
	response := models.APIResponse{
		Status: "success",
		Data: map[string]string{
			"message": "Welcome to AAA forum Backend!",
		},
	}
	json.NewEncoder(w).Encode(response)
}

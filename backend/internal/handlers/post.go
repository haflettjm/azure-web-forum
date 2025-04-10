package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/haflettjm/azure-web-forum/backend/internal/db"
	"github.com/haflettjm/azure-web-forum/backend/internal/models"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "List of posts")
	case http.MethodPost:
		var newPost models.Post
		err := json.NewDecoder(r.Body).Decode(&newPost)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
		}
		ctx := r.Context()
		err = db.DB.QueryRowContext(
			ctx,
			`INSERT INTO posts (userid, title, content, image_url, parent_id)
          	VALUES ($1, $2, $3, $4, $5)
           	RETURNING id`,
			newPost.UserID,
			newPost.Title,
			newPost.Content,
			newPost.Image_url,
			newPost.ParentID,
		).Scan(&newPost.ID)
		if err != nil {
			http.Error(w, "Database Error", http.StatusBadRequest)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

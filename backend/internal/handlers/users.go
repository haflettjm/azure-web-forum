package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/haflettjm/azure-web-forum/backend/internal/db"
	"github.com/haflettjm/azure-web-forum/backend/internal/models"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ctx := r.Context()

		// Query all users from the database
		rows, err := db.DB.QueryContext(ctx, "SELECT id, username, email FROM users")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []models.User

		// Loop through the rows
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
				http.Error(w, "Error scanning row", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		// Check for query iteration errors
		if err := rows.Err(); err != nil {
			http.Error(w, "Error reading result rows", http.StatusInternalServerError)
			return
		}

		// Return the result as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var newUser models.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if newUser.Username == "" || newUser.Email == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
		}
		ctx := r.Context()
		err = db.DB.QueryRowContext(
			ctx,
			"INSERT INTO users (username, email) VALUES ($1,$2) RETURNING id",
			newUser.Username,
			newUser.Email,
		).Scan(&newUser.ID)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
}

package models

type Post struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Image_url string `json:"image_url"`
	ParentID  *int   `json:"parent_id"`
}

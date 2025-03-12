package handlers

import (
	"net/http"
	"html/template"
	"ribbit/internal/database"
	"ribbit/internal/templates"
	"fmt"
	"time"
)

type PondPageData struct {
	Pond  *database.Pond
	Posts []PostTemplate
	User  *templates.UserTemplate
}

// PostTemplate adds display-specific fields to database.Post
type PostTemplate struct {
	ID          int64
	Title       string
	Description string
	Comments    int
	Likes       int
	PondName    string
	Author      string
	CreatedAt   time.Time
	TimeAgo     string
}

// Convert database.Post to PostTemplate
func convertToPostTemplate(post database.Post) PostTemplate {
	return PostTemplate{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		Comments:    post.Comments,
		Likes:       post.Likes,
		PondName:    post.PondName,
		Author:      post.Author,
		CreatedAt:   post.CreatedAt,
		TimeAgo:     formatTimeAgo(post.CreatedAt),
	}
}

// Helper function to format the time ago
func formatTimeAgo(t time.Time) string {
	duration := time.Since(t)
	hours := duration.Hours()

	// Less than 24 hours
	if hours < 24 {
		if hours < 1 {
			return "just now"
		}
		return fmt.Sprintf("%d hours ago", int(hours))
	}

	days := int(hours / 24)
	
	// Less than 30 days
	if days < 30 {
		if days == 1 {
			return "yesterday"
		}
		return fmt.Sprintf("%d days ago", days)
	}
	
	// Less than 365 days
	if days < 365 {
		months := days / 30
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	}
	
	// Years
	years := days / 365
	if years == 1 {
		return "1 year ago"
	}
	return fmt.Sprintf("%d years ago", years)
}

func HandlePondPage(w http.ResponseWriter, r *http.Request) {
	// Get pond name from URL query parameter
	pondName := r.URL.Query().Get("name")
	if pondName == "" {
		http.Error(w, "Pond name is required", http.StatusBadRequest)
		return
	}

	// Get database connection
	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get pond details
	pond, err := database.GetPondByName(db, pondName)
	if err != nil {
		http.Error(w, "Pond not found", http.StatusNotFound)
		return
	}

	// Get pond posts using existing GetPostsByPond function
	dbPosts, err := database.GetPostsByPond(db, pondName)
	if err != nil {
		dbPosts = []database.Post{}
	}

	// Convert database posts to template posts
	posts := make([]PostTemplate, len(dbPosts))
	for i, post := range dbPosts {
		posts[i] = convertToPostTemplate(post)
	}

	// Get user data if logged in
	var user *templates.UserTemplate
	if cookie, err := r.Cookie("user"); err == nil {
		user = templates.GetUserTemplate(cookie.Value)
	}

	// Render template
	tmpl, err := template.ParseFiles("templates/pond.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PondPageData{
		Pond:  pond,
		Posts: posts,
		User:  user,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
package handlers

import (
	"net/http"
	"html/template"
	"ribbit/internal/database"
	"ribbit/internal/templates"
	"fmt"
	"time"
	"strconv"
	"github.com/go-chi/chi/v5"
	"database/sql"
)

type PondPageData struct {
	Pond      *database.Pond
	Posts     []PostTemplate
	User      *templates.UserTemplate
	IsMember  bool
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

func GetPondByName(w http.ResponseWriter, r *http.Request) {
	pondName := chi.URLParam(r, "name")
	
	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get pond data
	var pond database.Pond
	err = db.QueryRow(`
		SELECT id, name, description, member_count, created_at 
		FROM ponds 
		WHERE name = ?`, pondName).Scan(
			&pond.ID,
			&pond.Name,
			&pond.Description,
			&pond.MemberCount,
			&pond.CreatedAt,
	)
	if err != nil {
		http.Error(w, "Pond not found", http.StatusNotFound)
		return
	}

	// Get posts for this pond
	posts, err := database.GetPondPosts(db, pondName)
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	// Convert posts to template format
	var postTemplates []PostTemplate
	for _, post := range posts {
		postTemplates = append(postTemplates, convertToPostTemplate(post))
	}

	// Get user from cookie and check membership
	var isMember bool
	cookie, err := r.Cookie("user")
	if err == nil {
		// Get user ID
		var userID int64
		err = db.QueryRow("SELECT id FROM users WHERE username = ?", cookie.Value).Scan(&userID)
		if err == nil {
			// Check if user is a member of this pond
			var exists bool
			err = db.QueryRow(`
				SELECT EXISTS(
					SELECT 1 
					FROM user_ponds 
					WHERE user_id = ? AND pond_id = ?
				)`, userID, pond.ID).Scan(&exists)
			if err == nil {
				isMember = exists
			}
		}
	}

	// Get user template if logged in
	var user *templates.UserTemplate
	if cookie != nil {
		user = templates.GetUserTemplate(cookie.Value)
	}

	// Prepare page data
	data := PondPageData{
		Pond:     &pond,
		Posts:    postTemplates,
		User:     user,
		IsMember: isMember,
	}

	// Render template
	tmpl, err := template.ParseFiles("templates/pond.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// JoinPond handles POST requests to join a pond
func JoinPond(w http.ResponseWriter, r *http.Request) {
	pondID := r.URL.Query().Get("pondID")
	if pondID == "" {
		http.Error(w, "Invalid pond ID", http.StatusBadRequest)
		return
	}

	// Convert pondID to int64
	pondIDInt, err := strconv.ParseInt(pondID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid pond ID format", http.StatusBadRequest)
		return
	}

	// Get user from cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get user ID
	var userID int64
	err = db.QueryRow("SELECT id FROM users WHERE username = ?", cookie.Value).Scan(&userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Insert user-pond relationship
	_, err = tx.Exec(`
		INSERT INTO user_ponds (user_id, pond_id, joined_at) 
		VALUES (?, ?, CURRENT_TIMESTAMP)
	`, userID, pondIDInt)
	if err != nil {
		http.Error(w, "Failed to join pond", http.StatusInternalServerError)
		return
	}

	// Update member count
	_, err = tx.Exec(`
		UPDATE ponds 
		SET member_count = member_count + 1 
		WHERE id = ?
	`, pondIDInt)
	if err != nil {
		http.Error(w, "Failed to update member count", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// LeavePond handles POST requests to leave a pond
func LeavePond(w http.ResponseWriter, r *http.Request) {
	pondID := r.URL.Query().Get("pondID")
	if pondID == "" {
		http.Error(w, "Invalid pond ID", http.StatusBadRequest)
		return
	}

	// Convert pondID to int64
	pondIDInt, err := strconv.ParseInt(pondID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid pond ID format", http.StatusBadRequest)
		return
	}

	// Get user from cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get user ID
	var userID int64
	err = db.QueryRow("SELECT id FROM users WHERE username = ?", cookie.Value).Scan(&userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Delete user-pond relationship
	result, err := tx.Exec(`
		DELETE FROM user_ponds 
		WHERE user_id = ? AND pond_id = ?
	`, userID, pondIDInt)
	if err != nil {
		http.Error(w, "Failed to leave pond", http.StatusInternalServerError)
		return
	}

	// Check if the user was actually a member
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "User was not a member of this pond", http.StatusBadRequest)
		return
	}

	// Update member count
	_, err = tx.Exec(`
		UPDATE ponds 
		SET member_count = member_count - 1 
		WHERE id = ?
	`, pondIDInt)
	if err != nil {
		http.Error(w, "Failed to update member count", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Add this function to handle the pond page route
func HandlePondPage(w http.ResponseWriter, r *http.Request) {
	pondName := r.URL.Query().Get("name")
	if pondName == "" {
		http.Error(w, "Pond name is required", http.StatusBadRequest)
		return
	}
	
	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get pond data
	var pond database.Pond
	query := `
		SELECT id, name, description, member_count, created_at 
		FROM ponds 
		WHERE name = ?`
	
	err = db.QueryRow(query, pondName).Scan(
		&pond.ID,
		&pond.Name,
		&pond.Description,
		&pond.MemberCount,
		&pond.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Pond not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Get posts for this pond
	posts, err := database.GetPondPosts(db, pondName)
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	// Convert posts to template format
	var postTemplates []PostTemplate
	for _, post := range posts {
		postTemplates = append(postTemplates, convertToPostTemplate(post))
	}

	// Get user from cookie and check membership
	var isMember bool
	cookie, err := r.Cookie("user")
	if err == nil {
		// Get user ID
		var userID int64
		err = db.QueryRow("SELECT id FROM users WHERE username = ?", cookie.Value).Scan(&userID)
		if err == nil {
			// Check if user is a member of this pond
			var exists bool
			err = db.QueryRow(`
				SELECT EXISTS(
					SELECT 1 
					FROM user_ponds 
					WHERE user_id = ? AND pond_id = ?
				)`, userID, pond.ID).Scan(&exists)
			if err == nil {
				isMember = exists
			}
		}
	}

	// Get user template if logged in
	var user *templates.UserTemplate
	if cookie != nil {
		user = templates.GetUserTemplate(cookie.Value)
	}

	// Prepare page data
	data := PondPageData{
		Pond:     &pond,
		Posts:    postTemplates,
		User:     user,
		IsMember: isMember,
	}

	// Render template
	tmpl, err := template.ParseFiles("templates/pond.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Helper function to calculate time ago
func timeAgo(t time.Time) string {
	duration := time.Since(t)
	
	switch {
	case duration.Hours() < 24:
		if hours := int(duration.Hours()); hours > 0 {
			if hours == 1 {
				return "1 hour ago"
			}
			return fmt.Sprintf("%d hours ago", hours)
		}
		if minutes := int(duration.Minutes()); minutes > 0 {
			if minutes == 1 {
				return "1 minute ago"
			}
			return fmt.Sprintf("%d minutes ago", minutes)
		}
		return "just now"
	case duration.Hours() < 48:
		return "yesterday"
	case duration.Hours() < 168:
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	default:
		return t.Format("Jan 2, 2006")
	}
}

// HandleDiscoverPonds handles the discover ponds page
func HandleDiscoverPonds(w http.ResponseWriter, r *http.Request) {
	// Get database connection
	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get all ponds sorted by member count
	ponds, err := database.GetAllPondsSortedByMembers(db)
	if err != nil {
		http.Error(w, "Failed to get ponds", http.StatusInternalServerError)
		return
	}

	// Get user from cookie if logged in
	var user *templates.UserTemplate
	if cookie, err := r.Cookie("user"); err == nil {
		user = templates.GetUserTemplate(cookie.Value)
	}

	// Create page data struct
	data := PageData{
		User:  user,
		Ponds: ponds,
	}

	// Parse and execute template
	tmpl, err := template.ParseFiles("templates/discover_ponds.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// PageData represents the data we'll pass to our template
type PageData struct {
	User  *templates.UserTemplate
	Ponds []database.Pond
}
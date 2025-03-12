package handlers

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
	"ribbit/internal/database"
)

// EngagementMetrics defines the ranges for different engagement levels
type EngagementMetrics struct {
	Likes    int
	Comments int
}

// getEngagementMetrics returns like and comment counts based on engagement level
func getEngagementMetrics(level string) EngagementMetrics {
	rand.Seed(time.Now().UnixNano())
	
	metrics := EngagementMetrics{}
	
	switch level {
	case "low":
		metrics.Likes = rand.Intn(200) + 200     // 200-400 likes
		metrics.Comments = rand.Intn(80) + 70    // 70-150 comments
	case "medium":
		metrics.Likes = rand.Intn(300) + 400     // 400-700 likes
		metrics.Comments = rand.Intn(100) + 150  // 150-250 comments
	case "high":
		metrics.Likes = rand.Intn(400) + 800     // 800-1200 likes
		metrics.Comments = rand.Intn(150) + 300  // 300-450 comments
	case "controversial":
		metrics.Likes = rand.Intn(400) + 400     // 400-800 likes
		metrics.Comments = rand.Intn(200) + 400  // 400-600 comments
	default:
		metrics.Likes = rand.Intn(200) + 200     // Default to low engagement
		metrics.Comments = rand.Intn(80) + 70
	}
	
	return metrics
}

// getRandomEngagementLevel returns an engagement level based on probability distribution
func getRandomEngagementLevel() string {
	rand.Seed(time.Now().UnixNano())
	roll := rand.Float64() * 100 // Random number between 0-100
	
	switch {
	case roll < 40: // 40% chance
		return "medium"
	case roll < 70: // 30% chance
		return "low"
	case roll < 90: // 20% chance
		return "controversial"
	default: // 10% chance
		return "high"
	}
}

func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user from cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Parse form data
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	pondName := r.FormValue("pond")
	
	// Automatically determine engagement level
	engagement := getRandomEngagementLevel()

	// Get engagement metrics
	metrics := getEngagementMetrics(engagement)

	// Get database connection
	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Insert new post
	_, err = db.Exec(`
		INSERT INTO ripples (
			title, content, like_count, comment_count,
			author_username, pond_name, created_at
		) VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`,
		title, content, metrics.Likes, metrics.Comments,
		cookie.Value, pondName,
	)

	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	// Redirect to home page after successful post creation
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// DeletePost handles the deletion of a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Only allow DELETE method
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get post ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	postID := parts[len(parts)-1]

	// Get user from cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get database connection
	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Verify post belongs to user before deleting
	var authorUsername string
	err = db.QueryRow("SELECT author_username FROM ripples WHERE id = ?", postID).Scan(&authorUsername)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	if authorUsername != cookie.Value {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Delete the post
	_, err = db.Exec("DELETE FROM ripples WHERE id = ?", postID)
	if err != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}

	// Return success
	w.WriteHeader(http.StatusOK)
}
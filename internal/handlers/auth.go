package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"ribbit/internal/database"
)

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	db, err := database.GetDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"taken": exists})
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("templates/signup.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	// Handle POST request
	username := r.FormValue("username")
	password := r.FormValue("password")
	description := r.FormValue("description")

	// Add logging
	log.Printf("Signup attempt - Username: %s, Description length: %d", username, len(description))

	if username == "" || password == "" || description == "" {
		renderSignupError(w, "All fields are required")
		return
	}

	// Validate username length and characters
	if len(username) < 3 || len(username) > 30 {
		renderSignupError(w, "Username must be between 3 and 30 characters")
		return
	}

	// Validate password length
	if len(password) < 8 {
		renderSignupError(w, "Password must be at least 8 characters long")
		return
	}

	db, err := database.GetDB()
	if err != nil {
		renderSignupError(w, "Database error")
		return
	}
	defer db.Close()

	// Check if username exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&exists)
	if err != nil {
		renderSignupError(w, "Database error")
		return
	}
	if exists {
		renderSignupError(w, "Username is already taken")
		return
	}

	// Insert new user with logging
	result, err := db.Exec(`
		INSERT INTO users (username, password_hash, description, join_date)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
	`, username, password, description)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		renderSignupError(w, "Failed to create account")
		return
	}

	// Verify the insert
	id, _ := result.LastInsertId()
	log.Printf("Created new user with ID: %d", id)

	// Set cookie and redirect to home
	http.SetCookie(w, &http.Cookie{
		Name:    "user",
		Value:   username,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func renderSignupError(w http.ResponseWriter, message string) {
	tmpl, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, struct{ ErrorMessage string }{message})
}
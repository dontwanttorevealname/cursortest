package main

import (
    "html/template"
    "log"
    "net/http"
    "ribbit/internal/templates"
    "strings"
    "github.com/joho/godotenv"
    "ribbit/internal/handlers"
    "ribbit/internal/database"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

// PageData represents the data we'll pass to our template
type PageData struct {
    User          *templates.UserTemplate
    ErrorMessage  string
    TrendingPosts []templates.Post
    AllPosts      []templates.Post
    SearchResults []templates.Post
    SearchPonds   []templates.Pond
    Query         string
}

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Create a new router
    r := chi.NewRouter()

    // Basic middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Serve static files
    fileServer := http.FileServer(http.Dir("static"))
    r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

    // Handle routes
    r.HandleFunc("/", handleHome)
    r.HandleFunc("/login", handleLoginSubmit)
    r.HandleFunc("/logout", handleLogout)
    r.HandleFunc("/trending", handleTrending)
    r.HandleFunc("/profile", handleProfile)
    r.HandleFunc("/search", handleSearch)
    r.HandleFunc("/create-post", handleCreatePost)
    r.HandleFunc("/submit-post", handlers.HandleCreatePost)
    r.HandleFunc("/api/posts/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodDelete {
            handlers.DeletePost(w, r)
            return
        }
        http.NotFound(w, r)
    })
    r.Get("/pond", handlers.HandlePondPage)
    r.Post("/api/ponds/join", handlers.JoinPond)
    r.Post("/api/ponds/leave", handlers.LeavePond)
    r.Get("/signup", handlers.HandleSignup)
    r.Post("/signup", handlers.HandleSignup)
    r.Get("/check-username", handlers.CheckUsername)
    r.Get("/discover", handlers.HandleDiscoverPonds)

    // Add template handler route
    r.Get("/templates/{template}", func(w http.ResponseWriter, r *http.Request) {
        templateName := chi.URLParam(r, "template")
        
        // Security check - only allow specific templates
        allowedTemplates := map[string]bool{
            "navbar.html": true,
            // Add other shared templates here as needed
        }
        
        if !allowedTemplates[templateName] {
            http.NotFound(w, r)
            return
        }

        // Get current user for navbar template
        user := getCurrentUser(r)
        
        // Create template data
        data := struct {
            User *templates.UserTemplate
        }{
            User: user,
        }

        // Parse and execute template
        tmpl := template.Must(template.ParseFiles("templates/" + templateName))
        err := tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    })

    // Start server
    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    // Check if user is logged in
    cookie, err := r.Cookie("user")
    if err != nil {
        // If not logged in, show login page
        tmpl, err := template.ParseFiles("templates/login.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
        return
    }

    // Get user template
    userTemplate := templates.GetUserTemplate(cookie.Value)
    if userTemplate == nil {
        // Invalid user, clear cookie and redirect to login
        http.SetCookie(w, &http.Cookie{
            Name:   "user",
            Value:  "",
            Path:   "/",
            MaxAge: -1,
        })
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    // Prepare data for the home page
    data := PageData{
        User: userTemplate,
    }

    tmpl, err := template.ParseFiles("templates/home.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleLoginSubmit(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    email := r.FormValue("username")
    password := r.FormValue("password")

    // Get user template and verify password
    userTemplate := templates.GetUserTemplate(email)
    if userTemplate != nil && password == userTemplate.Password {
        http.SetCookie(w, &http.Cookie{
            Name:  "user",
            Value: email,
            Path:  "/",
        })
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    // If login fails, show login page with error
    tmpl, err := template.ParseFiles("templates/login.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := PageData{
        ErrorMessage: "Invalid email or password",
    }
    tmpl.Execute(w, data)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
    // Clear the user cookie
    http.SetCookie(w, &http.Cookie{
        Name:   "user",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    })
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleTrending(w http.ResponseWriter, r *http.Request) {
    // Get database connection
    db, err := database.GetDB()
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Get trending posts with db connection (limit to 8)
    dbPosts, err := database.GetTrendingPosts(db, 8)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Convert database posts to template posts
    trendingPosts := make([]templates.Post, len(dbPosts))
    for i, post := range dbPosts {
        trendingPosts[i] = templates.Post{
            ID:          post.ID,
            Title:       post.Title,
            Description: post.Description,
            Comments:    post.Comments,
            Likes:       post.Likes,
            PondName:    post.PondName,
            Author:      post.Author,
            CreatedAt:   post.CreatedAt,
        }
    }

    // Get user data if logged in
    var user *templates.UserTemplate
    if cookie, err := r.Cookie("user"); err == nil {
        user = templates.GetUserTemplate(cookie.Value)
    }

    data := struct {
        User          *templates.UserTemplate
        TrendingPosts []templates.Post
    }{
        User:          user,
        TrendingPosts: trendingPosts,
    }

    // Create template with functions
    tmpl, err := template.New("trending.html").Funcs(template.FuncMap{
        "add": func(a, b int) int {
            return a + b
        },
    }).ParseFiles("templates/trending.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
    // Check if user is logged in
    cookie, err := r.Cookie("user")
    if err != nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    // Get user template
    userTemplate := templates.GetUserTemplate(cookie.Value)
    if userTemplate == nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    // Prepare data for the profile page
    allPosts, err := templates.GetAllPosts()
    if err != nil {
        log.Printf("Error getting all posts: %v", err)
        allPosts = []templates.Post{}
    }

    data := PageData{
        User:     userTemplate,
        AllPosts: allPosts,
    }

    tmpl, err := template.New("profile.html").Funcs(template.FuncMap{
        "add": func(a, b int) int {
            return a + b
        },
    }).ParseFiles("templates/profile.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
    query := strings.ToLower(r.URL.Query().Get("q"))
    
    // Get user template if logged in
    var userTemplate *templates.UserTemplate
    if cookie, err := r.Cookie("user"); err == nil {
        userTemplate = templates.GetUserTemplate(cookie.Value)
    }

    // If no query, redirect to home
    if query == "" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    // Search through all posts and ponds
    var results []templates.Post
    var pondResults []templates.Pond
    
    db, err := database.GetDB()
    if err != nil {
        log.Printf("Error getting database: %v", err)
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Search posts
    allPosts, err := templates.GetAllPosts()
    if err != nil {
        log.Printf("Error getting posts: %v", err)
        allPosts = []templates.Post{}
    }
    
    for _, post := range allPosts {
        if strings.Contains(strings.ToLower(post.Title), query) ||
           strings.Contains(strings.ToLower(post.Description), query) {
            results = append(results, post)
        }
    }

    // Search ponds
    allPonds, err := database.GetAllPonds(db)
    if err != nil {
        log.Printf("Error getting ponds: %v", err)
        allPonds = []database.Pond{}
    }

    for _, pond := range allPonds {
        if strings.Contains(strings.ToLower(pond.Name), query) ||
           strings.Contains(strings.ToLower(pond.Description), query) {
            pondResults = append(pondResults, templates.ConvertDatabasePond(pond))
        }
    }

    // Prepare data for the template
    data := PageData{
        User:          userTemplate,
        SearchResults: results,
        SearchPonds:   pondResults,
        Query:         r.URL.Query().Get("q"),
    }

    // Parse and execute template
    tmpl, err := template.ParseFiles("templates/search.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleCreatePost(w http.ResponseWriter, r *http.Request) {
    // Check if user is logged in
    cookie, err := r.Cookie("user")
    if err != nil {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Get user template
    userTemplate := templates.GetUserTemplate(cookie.Value)
    if userTemplate == nil {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Prepare data for the template
    data := PageData{
        User: userTemplate,
    }

    // Parse and execute template
    tmpl, err := template.ParseFiles("templates/create_post.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Update getCurrentUser function to use cookie-based auth
func getCurrentUser(r *http.Request) *templates.UserTemplate {
    // Get user from cookie
    cookie, err := r.Cookie("user")
    if err != nil {
        return nil
    }
    
    // Get user template using the email from cookie
    userTemplate := templates.GetUserTemplate(cookie.Value)
    return userTemplate
} 
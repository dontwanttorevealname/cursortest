package main

import (
    "html/template"
    "log"
    "net/http"
    "ribbit/internal/templates"
)

// PageData represents the data we'll pass to our template
type PageData struct {
    User          *templates.UserTemplate
    ErrorMessage  string
    TrendingPosts []templates.Post
    AllPosts      []templates.Post
}

func main() {
    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Handle routes
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/login", handleLoginSubmit)
    http.HandleFunc("/logout", handleLogout)
    http.HandleFunc("/trending", handleTrending)
    http.HandleFunc("/profile", handleProfile)

    // Start server
    log.Println("Server starting on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
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

    // Get trending posts
    trendingPosts := templates.GetTrendingPosts()

    // Prepare data for the trending page
    data := PageData{
        User:          userTemplate,
        TrendingPosts: trendingPosts,
    }

    tmpl, err := template.New("trending.html").Funcs(template.FuncMap{
        "add": func(a, b int) int {
            return a + b
        },
        "div": func(a, b int) float64 {
            return float64(a) / float64(b)
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
    data := PageData{
        User:     userTemplate,
        AllPosts: templates.GetAllPosts(),
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
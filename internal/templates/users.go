package templates

import (
    _ "github.com/tursodatabase/libsql-client-go/libsql"
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "time"
    "ribbit/internal/database"
    "database/sql"
    "strings"
    "html/template"
)

type UserTemplate struct {
    ID            int64
    Email         string    // Using username from SQL as email
    Password      string    // Using password_hash from SQL
    Description   string    // Make sure this field exists
    JoinDate      time.Time
    OfficialPosts []Post    // Add this field for official posts
    Posts         []Post    // This will be pond-specific posts
    Ponds         []Pond
}

type Post struct {
    ID          int64
    Title       string
    Description string
    Comments    int
    Likes       int
    PondName    string
    Author      string
    CreatedAt   time.Time
    TimeAgo     string    // Add this field
}

type Pond struct {
    Name        string
    Description string
    Members     string
}

// Helper function to get random posts from a pond
func getRandomPostsFromPond(db *sql.DB, pondName string, minCount int, excludeAuthor string) ([]Post, error) {
    rows, err := db.Query(`
        SELECT id, title, content, comment_count, like_count, pond_name, author_username, created_at
        FROM ripples
        WHERE pond_name = ? AND author_username != ?
        ORDER BY RANDOM()
        LIMIT ?`, 
        pondName, excludeAuthor, minCount+3)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var dbPost database.Post
        err := rows.Scan(
            &dbPost.ID,
            &dbPost.Title,
            &dbPost.Description,
            &dbPost.Comments,
            &dbPost.Likes,
            &dbPost.PondName,
            &dbPost.Author,
            &dbPost.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        posts = append(posts, ConvertDatabasePost(dbPost))
    }

    return posts, nil
}

// Get random posts for a user based on their pond memberships
func getRandomPostsForUser(db *sql.DB, userPonds []Pond, excludeAuthor string) ([]Post, error) {
    if len(userPonds) == 0 {
        return []Post{}, nil
    }

    // Build list of pond names
    pondNames := make([]string, len(userPonds))
    for i, pond := range userPonds {
        pondNames[i] = pond.Name
    }

    // Create the placeholder string for the IN clause
    placeholders := make([]string, len(pondNames))
    for i := range pondNames {
        placeholders[i] = "?"
    }
    placeholderString := strings.Join(placeholders, ",")

    // Create the query arguments (pond names + exclude author)
    args := make([]interface{}, len(pondNames)+1)
    for i, name := range pondNames {
        args[i] = name
    }
    args[len(args)-1] = excludeAuthor

    // Query random posts from user's ponds
    query := fmt.Sprintf(`
        SELECT id, title, content, comment_count, like_count, pond_name, author_username, created_at
        FROM ripples
        WHERE pond_name IN (%s)
        AND author_username != ?
        ORDER BY RANDOM()
        LIMIT 20`, placeholderString)

    rows, err := db.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var dbPost database.Post
        err := rows.Scan(
            &dbPost.ID,
            &dbPost.Title,
            &dbPost.Description,
            &dbPost.Comments,
            &dbPost.Likes,
            &dbPost.PondName,
            &dbPost.Author,
            &dbPost.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        posts = append(posts, ConvertDatabasePost(dbPost))
    }

    return posts, nil
}

// Convert database.Post to templates.Post
func ConvertDatabasePost(dbPost database.Post) Post {
    return Post{
        ID:          dbPost.ID,
        Title:       dbPost.Title,
        Description: dbPost.Description,
        Comments:    dbPost.Comments,
        Likes:       dbPost.Likes,
        PondName:    dbPost.PondName,
        Author:      dbPost.Author,
        CreatedAt:   dbPost.CreatedAt,
        TimeAgo:     dbPost.TimeAgo,
    }
}

// ConvertDatabasePond converts a database.Pond to a templates.Pond
func ConvertDatabasePond(dbPond database.Pond) Pond {
    return Pond{
        Name:        dbPond.Name,
        Description: dbPond.Description,
        Members:     formatMemberCount(dbPond.MemberCount),
    }
}

// Convert slice of database posts to template posts
func ConvertDatabasePosts(dbPosts []database.Post) []Post {
    posts := make([]Post, len(dbPosts))
    for i, dbPost := range dbPosts {
        posts[i] = ConvertDatabasePost(dbPost)
    }
    return posts
}

// Convert slice of database ponds to template ponds
func convertDatabasePonds(dbPonds []database.Pond) []Pond {
    ponds := make([]Pond, len(dbPonds))
    for i, dbPond := range dbPonds {
        ponds[i] = ConvertDatabasePond(dbPond)
    }
    return ponds
}

// Add a method to get official posts
func (u *UserTemplate) GetOfficialPosts() []Post {
    db, err := database.GetDB()
    if err != nil {
        return nil
    }
    defer db.Close()

    dbPosts, err := database.GetOfficialPosts(db, 10)
    if err != nil {
        return nil
    }

    return ConvertDatabasePosts(dbPosts)
}

// Add these new functions
var templateFuncs = template.FuncMap{
    "add": func(a, b int) int {
        return a + b
    },
}

// Update GetUserTemplate to use the function map
func GetUserTemplate(username string) *UserTemplate {
    db, err := database.GetDB()
    if err != nil {
        return nil
    }
    defer db.Close()

    var user UserTemplate
    err = db.QueryRow(`
        SELECT id, username, password_hash, description, join_date 
        FROM users 
        WHERE username = ?`, username).Scan(
        &user.ID,
        &user.Email,
        &user.Password,
        &user.Description,
        &user.JoinDate,
    )
    if err != nil {
        return nil
    }

    // Get official posts
    officialPosts, err := database.GetOfficialPosts(db, 10)
    if err != nil {
        officialPosts = []database.Post{}
    }

    // Get user's pond posts
    pondPosts, err := database.GetUserFeed(db, user.ID, 0, 20)
    if err != nil {
        pondPosts = []database.Post{}
    }

    // Get user's ponds
    ponds, err := database.GetUserPonds(db, username)
    if err != nil {
        ponds = []database.Pond{}
    }

    return &UserTemplate{
        ID:            user.ID,
        Email:         user.Email,
        Password:      user.Password,
        Description:   user.Description,
        JoinDate:      user.JoinDate,
        OfficialPosts: ConvertDatabasePosts(officialPosts),
        Posts:         ConvertDatabasePosts(pondPosts),
        Ponds:         convertDatabasePonds(ponds),
    }
}

// Helper function to format member count
func formatMemberCount(count int) string {
    if count >= 1000 {
        return fmt.Sprintf("%.1fK", float64(count)/1000)
    }
    return fmt.Sprintf("%d", count)
}

// GetTrendingPosts returns the top 8 posts by engagement (likes + comments)
func GetTrendingPosts(db *sql.DB) ([]Post, error) {
    if db == nil {
        return nil, fmt.Errorf("database connection is nil")
    }

    rows, err := db.Query(`
        SELECT id, title, content, comment_count, like_count, pond_name, author_username, created_at
        FROM ripples
        ORDER BY (comment_count + like_count) DESC
        LIMIT 8`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var dbPost database.Post
        err := rows.Scan(
            &dbPost.ID,
            &dbPost.Title,
            &dbPost.Description,
            &dbPost.Comments,
            &dbPost.Likes,
            &dbPost.PondName,
            &dbPost.Author,
            &dbPost.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        // Set TimeAgo before converting to template post
        dbPost.TimeAgo = formatTimeAgo(dbPost.CreatedAt)
        posts = append(posts, ConvertDatabasePost(dbPost))
    }

    return posts, nil
}

// Add the formatTimeAgo function
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

// GetAllPosts retrieves all posts from the database
func GetAllPosts() ([]Post, error) {
    db, err := database.GetDB()
    if err != nil {
        return nil, err
    }
    defer db.Close()

    dbPosts, err := database.GetAllPosts(db)
    if err != nil {
        return nil, err
    }

    // Make sure TimeAgo is set for each post
    for i := range dbPosts {
        if dbPosts[i].TimeAgo == "" {
            dbPosts[i].TimeAgo = formatTimeAgo(dbPosts[i].CreatedAt)
        }
    }

    return ConvertDatabasePosts(dbPosts), nil
}

// Update GetPaginatedPosts to use conversions
func (u *UserTemplate) GetPaginatedPosts(start, count int) []Post {
    db, err := database.GetDB()
    if err != nil {
        return nil
    }
    defer db.Close()

    dbPosts, err := database.GetUserFeed(db, u.ID, start, count)
    if err != nil {
        return nil
    }

    return ConvertDatabasePosts(dbPosts)
}

// Update HandleGetPosts to handle both official and pond posts
func HandleGetPosts(w http.ResponseWriter, r *http.Request) {
    // Parse query parameters
    startStr := r.URL.Query().Get("start")
    countStr := r.URL.Query().Get("count")
    postType := r.URL.Query().Get("type")
    
    start, err := strconv.Atoi(startStr)
    if err != nil {
        start = 0
    }
    
    count, err := strconv.Atoi(countStr)
    if err != nil {
        count = 3
    }
    
    // Get user from session
    user, ok := r.Context().Value("user").(*UserTemplate)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    var posts []Post
    if postType == "official" {
        posts = user.GetOfficialPosts()
    } else {
        posts = user.GetPaginatedPosts(start, count)
    }
    
    w.Header().Set("Content-Type", "application/json")
    
    if err := json.NewEncoder(w).Encode(posts); err != nil {
        http.Error(w, "Failed to encode posts", http.StatusInternalServerError)
        return
    }
}

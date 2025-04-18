package database

import (
    "database/sql"
    "fmt"
    "time"
)

type Post struct {
    ID           int64
    Title        string
    Description  string
    Comments     int
    Likes        int
    PondName     string
    Author       string
    CreatedAt    time.Time
    TimeAgo       string
}

// GetAllPosts retrieves all posts from the database, sorted by creation time
func GetAllPosts(db *sql.DB) ([]Post, error) {
    rows, err := db.Query(`
        SELECT id, title, content, comment_count, like_count, pond_name, author_username, created_at
        FROM ripples
        ORDER BY created_at DESC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        err := rows.Scan(
            &post.ID,
            &post.Title,
            &post.Description,
            &post.Comments,
            &post.Likes,
            &post.PondName,
            &post.Author,
            &post.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        post.TimeAgo = FormatTimeAgo(post.CreatedAt)
        posts = append(posts, post)
    }
    return posts, nil
}

// GetPostsByPond retrieves all posts from a specific pond
func GetPostsByPond(db *sql.DB, pondName string) ([]Post, error) {
    rows, err := db.Query(`
        SELECT id, title, content, comment_count, like_count, pond_name, author_username, created_at
        FROM ripples
        WHERE pond_name = ?
        ORDER BY created_at DESC`, pondName)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        err := rows.Scan(
            &post.ID,
            &post.Title,
            &post.Description,
            &post.Comments,
            &post.Likes,
            &post.PondName,
            &post.Author,
            &post.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        post.TimeAgo = FormatTimeAgo(post.CreatedAt)
        posts = append(posts, post)
    }
    return posts, nil
}

// GetPaginatedPosts retrieves a subset of posts for pagination
func GetPaginatedPosts(db *sql.DB, start, count int) ([]Post, error) {
    rows, err := db.Query(`
        SELECT id, title, content, comment_count, like_count, pond_name, author_username, created_at
        FROM ripples
        ORDER BY created_at DESC
        LIMIT ? OFFSET ?`, count, start)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        err := rows.Scan(
            &post.ID,
            &post.Title,
            &post.Description,
            &post.Comments,
            &post.Likes,
            &post.PondName,
            &post.Author,
            &post.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        post.TimeAgo = FormatTimeAgo(post.CreatedAt)
        posts = append(posts, post)
    }
    return posts, nil
}

// GetOfficialPosts retrieves official posts from the database
func GetOfficialPosts(db *sql.DB, count int) ([]Post, error) {
    rows, err := db.Query(`
        SELECT id, title, content, comment_count, like_count, 
               pond_name, author_username, created_at
        FROM ripples
        WHERE pond_name = 'Official'
        ORDER BY created_at DESC
        LIMIT ?`, count)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        err := rows.Scan(
            &post.ID,
            &post.Title,
            &post.Description,
            &post.Comments,
            &post.Likes,
            &post.PondName,
            &post.Author,
            &post.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        post.TimeAgo = FormatTimeAgo(post.CreatedAt)
        posts = append(posts, post)
    }
    return posts, nil
}

// GetUserFeed retrieves posts from user's ponds (excluding official posts)
func GetUserFeed(db *sql.DB, userID int64, start, count int) ([]Post, error) {
    rows, err := db.Query(`
        WITH UserPonds AS (
            SELECT DISTINCT p.name 
            FROM ponds p 
            JOIN user_ponds up ON p.id = up.pond_id 
            WHERE up.user_id = ?
        )
        SELECT r.id, r.title, r.content, r.comment_count, r.like_count, 
               r.pond_name, r.author_username, r.created_at
        FROM ripples r
        WHERE r.pond_name IN (SELECT name FROM UserPonds)
        AND r.pond_name != 'Official'  -- Explicitly exclude official posts
        ORDER BY RANDOM()
        LIMIT ?`, userID, count)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        err := rows.Scan(
            &post.ID,
            &post.Title,
            &post.Description,
            &post.Comments,
            &post.Likes,
            &post.PondName,
            &post.Author,
            &post.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        post.TimeAgo = FormatTimeAgo(post.CreatedAt)
        posts = append(posts, post)
    }

    return posts, nil
}

// GetRandomPostsFromPond gets random posts from a specific pond
func GetRandomPostsFromPond(db *sql.DB, pondName string, count int) ([]Post, error) {
    rows, err := db.Query(`
        SELECT id, title, content, comment_count, like_count, 
               pond_name, author_username, created_at
        FROM ripples
        WHERE pond_name = ?
        ORDER BY RANDOM()
        LIMIT ?`, pondName, count)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        err := rows.Scan(
            &post.ID,
            &post.Title,
            &post.Description,
            &post.Comments,
            &post.Likes,
            &post.PondName,
            &post.Author,
            &post.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        post.TimeAgo = FormatTimeAgo(post.CreatedAt)
        posts = append(posts, post)
    }
    return posts, nil
}

// GetRandomPostsFromUserPonds gets random posts from all ponds a user is member of
func GetRandomPostsFromUserPonds(db *sql.DB, userID int64, postsPerPond int) ([]Post, error) {
    // First get all user's ponds
    rows, err := db.Query(`
        SELECT DISTINCT p.name 
        FROM ponds p 
        JOIN user_ponds up ON p.id = up.pond_id 
        WHERE up.user_id = ?`, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var pondNames []string
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            return nil, err
        }
        pondNames = append(pondNames, name)
    }

    // Get random posts from each pond
    var allPosts []Post
    for _, pondName := range pondNames {
        posts, err := GetRandomPostsFromPond(db, pondName, postsPerPond)
        if err != nil {
            return nil, err
        }
        allPosts = append(allPosts, posts...)
    }

    return allPosts, nil
}

// GetTrendingPosts returns the top N trending posts, excluding official posts and Ribbit Admin posts
func GetTrendingPosts(db *sql.DB, limit int) ([]Post, error) {
    query := `
        SELECT id, title, content, like_count, comment_count, author_username, pond_name, created_at
        FROM ripples 
        WHERE pond_name != 'OFFICIAL' 
        AND author_username != 'Ribbit Admin'
        ORDER BY (like_count + comment_count) DESC, created_at DESC
        LIMIT ?
    `
    
    rows, err := db.Query(query, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        err := rows.Scan(
            &post.ID,
            &post.Title,
            &post.Description,
            &post.Likes,
            &post.Comments,
            &post.Author,
            &post.PondName,
            &post.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        post.TimeAgo = FormatTimeAgo(post.CreatedAt)
        posts = append(posts, post)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return posts, nil
}

// GetPondPosts retrieves all posts for a specific pond
func GetPondPosts(db *sql.DB, pondName string) ([]Post, error) {
    query := `
        SELECT id, title, content, comment_count, like_count, 
               pond_name, author_username, created_at
        FROM ripples
        WHERE pond_name = ?
        ORDER BY created_at DESC
    `
    
    rows, err := db.Query(query, pondName)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        err := rows.Scan(
            &post.ID,
            &post.Title,
            &post.Description,  // Note: This maps to 'content' in the database
            &post.Comments,
            &post.Likes,
            &post.PondName,
            &post.Author,
            &post.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        post.TimeAgo = FormatTimeAgo(post.CreatedAt)
        posts = append(posts, post)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return posts, nil
}

// FormatTimeAgo formats a time.Time into a human-readable "time ago" string
func FormatTimeAgo(t time.Time) string {
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
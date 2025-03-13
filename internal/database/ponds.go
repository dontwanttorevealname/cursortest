package database

import (
    "database/sql"
)

type Pond struct {
    ID          int64
    Name        string
    Description string
    MemberCount int
    CreatedAt   string
}

// GetAllPonds retrieves all ponds from the database
func GetAllPonds(db *sql.DB) ([]Pond, error) {
    rows, err := db.Query(`
        SELECT id, name, description, member_count, created_at
        FROM ponds
        ORDER BY member_count DESC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ponds []Pond
    for rows.Next() {
        var pond Pond
        err := rows.Scan(
            &pond.ID,
            &pond.Name,
            &pond.Description,
            &pond.MemberCount,
            &pond.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        ponds = append(ponds, pond)
    }
    return ponds, nil
}

// GetPondByName retrieves a specific pond by its name
func GetPondByName(db *sql.DB, name string) (*Pond, error) {
    var pond Pond
    err := db.QueryRow(`
        SELECT id, name, description, member_count, created_at
        FROM ponds
        WHERE name = ?`, name).Scan(
        &pond.ID,
        &pond.Name,
        &pond.Description,
        &pond.MemberCount,
        &pond.CreatedAt,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &pond, nil
}

// GetTrendingPonds retrieves the top N ponds by member count
func GetTrendingPonds(db *sql.DB, limit int) ([]Pond, error) {
    rows, err := db.Query(`
        SELECT id, name, description, member_count, created_at
        FROM ponds
        ORDER BY member_count DESC
        LIMIT ?`, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ponds []Pond
    for rows.Next() {
        var pond Pond
        err := rows.Scan(
            &pond.ID,
            &pond.Name,
            &pond.Description,
            &pond.MemberCount,
            &pond.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        ponds = append(ponds, pond)
    }
    return ponds, nil
}

// GetUserPonds gets all ponds a user is a member of
func GetUserPonds(db *sql.DB, username string) ([]Pond, error) {
    query := `
        SELECT p.id, p.name, p.description, p.member_count
        FROM ponds p
        JOIN user_ponds up ON p.id = up.pond_id
        JOIN users u ON up.user_id = u.id
        WHERE u.username = ?
    `
    
    rows, err := db.Query(query, username)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ponds []Pond
    for rows.Next() {
        var pond Pond
        err := rows.Scan(
            &pond.ID,
            &pond.Name,
            &pond.Description,
            &pond.MemberCount,
        )
        if err != nil {
            return nil, err
        }
        ponds = append(ponds, pond)
    }

    return ponds, nil
}

// GetPosts gets all posts ordered by creation date
func GetPosts(db *sql.DB) ([]Post, error) {
    query := `
        SELECT id, title, content, comment_count, like_count, 
               pond_name, author_username, created_at
        FROM ripples
        ORDER BY created_at DESC
    `
    
    rows, err := db.Query(query)
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
        posts = append(posts, post)
    }

    return posts, nil
}

// GetAllPondsSortedByMembers returns all ponds sorted by member count
func GetAllPondsSortedByMembers(db *sql.DB) ([]Pond, error) {
    rows, err := db.Query(`
        SELECT id, name, description, member_count, created_at 
        FROM ponds 
        ORDER BY member_count DESC
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ponds []Pond
    for rows.Next() {
        var pond Pond
        err := rows.Scan(
            &pond.ID,
            &pond.Name,
            &pond.Description,
            &pond.MemberCount,
            &pond.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        ponds = append(ponds, pond)
    }

    return ponds, nil
} 
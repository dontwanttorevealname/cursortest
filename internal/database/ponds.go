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
package database

import (
    "database/sql"
    "time"
    _ "github.com/tursodatabase/libsql-client-go/libsql"
    "log"
    "os"
    "strings"
)

type User struct {
    ID           int64
    Username     string    // Using username from SQL
    PasswordHash string    // Using password_hash from SQL
    Description  string
    JoinDate     time.Time
}

// GetDB returns a database connection using environment variables
func GetDB() (*sql.DB, error) {
    dbURL := os.Getenv("TURSO_DATABASE_URL")
    authToken := os.Getenv("TURSO_AUTH_TOKEN")
    
    if !strings.HasPrefix(dbURL, "libsql://") {
        dbURL = "libsql://" + dbURL
    }
    
    fullURL := dbURL + "?authToken=" + authToken
    
    return sql.Open("libsql", fullURL)
}

// GetUser retrieves a user from the database by username
func GetUser(db *sql.DB, username string) (*User, error) {
    var user User
    err := db.QueryRow(`
        SELECT id, username, password_hash, description, join_date 
        FROM users 
        WHERE username = ?`, username).Scan(
        &user.ID,
        &user.Username,
        &user.PasswordHash,
        &user.Description,
        &user.JoinDate,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}

// CreateUser adds a new user to the database
func CreateUser(db *sql.DB, username, password, description string) error {
    _, err := db.Exec(`
        INSERT INTO users (username, password_hash, description, join_date)
        VALUES (?, ?, ?, ?)`,
        username, password, description, time.Now())
    return err
}

// UpdateUser updates an existing user's information
func UpdateUser(db *sql.DB, user *User) error {
    _, err := db.Exec(`
        UPDATE users 
        SET username = ?, password_hash = ?, description = ?
        WHERE id = ?`,
        user.Username, user.PasswordHash, user.Description, user.ID)
    return err
}

// DeleteUser removes a user and their associated data
func DeleteUser(db *sql.DB, id int64) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    // Delete user's pond memberships
    _, err = tx.Exec(`DELETE FROM user_ponds WHERE user_id = ?`, id)
    if err != nil {
        tx.Rollback()
        return err
    }

    // Delete the user
    _, err = tx.Exec(`DELETE FROM users WHERE id = ?`, id)
    if err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit()
}

// ValidateUserCredentials checks if the username/password combination is valid
func ValidateUserCredentials(db *sql.DB, username, password string) bool {
    var storedPassword string
    err := db.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&storedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("No user found with username: %s", username)
        } else {
            log.Printf("Error querying user: %v", err)
        }
        return false
    }

    // TODO: Implement proper password hashing
    return password == storedPassword
}

// GetUserPonds retrieves all ponds that a user is a member of
func GetUserPonds(db *sql.DB, userID int64) ([]Pond, error) {
    rows, err := db.Query(`
        SELECT p.name, p.description, p.member_count 
        FROM ponds p 
        JOIN user_ponds up ON p.id = up.pond_id 
        WHERE up.user_id = ?`, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ponds []Pond
    for rows.Next() {
        var pond Pond
        var memberCount int
        err := rows.Scan(&pond.Name, &pond.Description, &memberCount)
        if err != nil {
            return nil, err
        }
        pond.MemberCount = memberCount
        ponds = append(ponds, pond)
    }
    return ponds, nil
}


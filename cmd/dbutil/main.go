package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/pressly/goose/v3"
)

type User struct {
	Username    string
	Description string
	JoinDate    string
	Ponds       []Pond
	Posts       []Post
}

type Pond struct {
	Name        string
	Description string
	MemberCount int
}

type Post struct {
	Title        string
	Content      string
	LikeCount    int
	CommentCount int
	PondName     string
	CreatedAt    string
}

func resetMigrations(db *sql.DB, dir string) error {
	// Drop goose_db_version table to force a clean state
	_, err := db.Exec("DROP TABLE IF EXISTS goose_db_version;")
	if err != nil {
		return fmt.Errorf("failed to drop version table: %v", err)
	}
	return nil
}

func main() {
	// Create context
	ctx := context.Background()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	// Create database connection string
	connStr := fmt.Sprintf("%s?authToken=%s", dbURL, authToken)

	// Open database connection
	db, err := sql.Open("libsql", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("✅ Connected to database successfully")

	// Set up goose
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal("Error setting dialect:", err)
	}

	fmt.Println("�� Resetting database state...")
	if err := resetMigrations(db, "migrations"); err != nil {
		log.Printf("Warning during reset: %v", err)
	}

	// Drop all tables to ensure clean state
	tables := []string{"user_ponds", "ripples", "ponds", "users"}
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", table))
		if err != nil {
			log.Printf("Warning dropping table %s: %v", table, err)
		}
	}

	fmt.Println("🔄 Running schema migrations...")
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal("Error running schema migrations:", err)
	}

	fmt.Println("🔄 Running seed migrations...")
	if err := resetMigrations(db, "seeds"); err != nil {
		log.Printf("Warning during seed reset: %v", err)
	}
	if err := goose.Up(db, "seeds"); err != nil {
		log.Fatal("Error running seed migrations:", err)
	}

	fmt.Println("✅ Database setup complete")

	// Debug: Check user_ponds table
	var rows *sql.Rows
	rows, err = db.QueryContext(ctx, `
		SELECT u.username, p.name 
		FROM user_ponds up
		JOIN users u ON u.id = up.user_id
		JOIN ponds p ON p.id = up.pond_id
		ORDER BY u.username, p.name
	`)
	if err != nil {
		log.Printf("Debug: Error querying user_ponds: %v", err)
	} else {
		fmt.Println("\n🔍 Debug: Current user_pond memberships:")
		for rows.Next() {
			var username, pondName string
			if err := rows.Scan(&username, &pondName); err != nil {
				log.Printf("Debug: Error scanning row: %v", err)
				continue
			}
			fmt.Printf("  %s -> %s\n", username, pondName)
		}
		rows.Close()
	}

	// Print the current tables and their schemas
	fmt.Println("\nDebug: Checking database schema...")
	for _, table := range tables {
		var count int
		err = db.QueryRowContext(ctx, fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&count)
		if err != nil {
			fmt.Printf("Error checking %s table: %v\n", table, err)
		} else {
			fmt.Printf("Table %s has %d rows\n", table, count)
		}
	}

	fmt.Println("\n🐸 Database State:")
	fmt.Println(strings.Repeat("=", 80))

	// Modified query to prevent duplicates
	query := `
		SELECT DISTINCT
			u.username,
			u.description,
			u.join_date,
			p.name,
			p.description,
			p.member_count,
			r.title,
			r.content,
			r.like_count,
			r.comment_count,
			r.pond_name,
			r.created_at
		FROM users u
		LEFT JOIN user_ponds up ON u.id = up.user_id
		LEFT JOIN ponds p ON up.pond_id = p.id
		LEFT JOIN ripples r ON (
			CASE 
				WHEN u.username = 'admin@ribbit.com' AND r.author_username = 'Ribbit Admin' THEN 1
				WHEN u.username = r.author_username THEN 1
				ELSE 0
			END
		) = 1
		ORDER BY u.username, p.name, r.created_at DESC
	`

	rows, err = db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal("Error querying database:", err)
	}
	defer rows.Close()

	users := make(map[string]*User)

	// Process results
	for rows.Next() {
		var (
			username     string
			userDesc    string
			joinDate    string
			pondName    sql.NullString
			pondDesc    sql.NullString
			memberCount sql.NullInt64
			title       sql.NullString
			content     sql.NullString
			likeCount   sql.NullInt64
			commentCount sql.NullInt64
			postPondName sql.NullString
			createdAt   sql.NullString
		)

		if err := rows.Scan(
			&username, &userDesc, &joinDate,
			&pondName, &pondDesc, &memberCount,
			&title, &content, &likeCount, &commentCount,
			&postPondName, &createdAt,
		); err != nil {
			log.Fatal("Error scanning row:", err)
		}

		// Get or create user
		user, exists := users[username]
		if !exists {
			user = &User{
				Username:    username,
				Description: userDesc,
				JoinDate:    joinDate,
				Ponds:       make([]Pond, 0),
				Posts:       make([]Post, 0),
			}
			users[username] = user
		}

		// Add pond if it exists and isn't already added
		if pondName.Valid {
			pondExists := false
			for _, p := range user.Ponds {
				if p.Name == pondName.String {
					pondExists = true
					break
				}
			}
			if !pondExists {
				user.Ponds = append(user.Ponds, Pond{
					Name:        pondName.String,
					Description: pondDesc.String,
					MemberCount: int(memberCount.Int64),
				})
			}
		}

		// Add post if it exists
		if title.Valid {
			// Check if post already exists
			postExists := false
			for _, p := range user.Posts {
				if p.Title == title.String && p.CreatedAt == createdAt.String {
					postExists = true
					break
				}
			}
			if !postExists {
				user.Posts = append(user.Posts, Post{
					Title:        title.String,
					Content:      content.String,
					LikeCount:    int(likeCount.Int64),
					CommentCount: int(commentCount.Int64),
					PondName:     postPondName.String,
					CreatedAt:    createdAt.String,
				})
			}
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
	}

	if len(users) == 0 {
		fmt.Println("❌ No users found in database")
	} else {
		for _, user := range users {
			fmt.Printf("\n👤 User: %s\n", user.Username)
			fmt.Printf("📝 Description: %s\n", user.Description)
			fmt.Printf("📅 Joined: %s\n", user.JoinDate)
			
			if len(user.Ponds) > 0 {
				fmt.Println("\n🌿 Member of Ponds:")
				for _, pond := range user.Ponds {
					fmt.Printf("  - %s (%d members)\n", pond.Name, pond.MemberCount)
					fmt.Printf("    %s\n", pond.Description)
				}
			} else {
				fmt.Println("\n🌿 No pond memberships")
			}

			if len(user.Posts) > 0 {
				fmt.Println("\n📝 Posts:")
				for _, post := range user.Posts {
					fmt.Printf("\n  📌 %s\n", post.Title)
					fmt.Printf("  🏷️ Posted in: %s\n", post.PondName)
					fmt.Printf("  📅 %s\n", post.CreatedAt)
					fmt.Printf("  👍 %d likes | 💬 %d comments\n", post.LikeCount, post.CommentCount)
					fmt.Printf("  %s\n", post.Content)
				}
			} else {
				fmt.Println("\n📝 No posts")
			}
			
			fmt.Println(strings.Repeat("-", 80))
		}
	}
} 
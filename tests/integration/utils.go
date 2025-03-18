package integration

import (
	"fmt"
	"math/rand"
	"time"
)

// init initializes the random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// generateRandomString generates a random string of the specified length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// generateUniqueUsername generates a unique username for testing
func generateUniqueUsername() string {
	return fmt.Sprintf("testuser_%s", generateRandomString(8))
}

// generateUniquePostTitle generates a unique post title for testing
func generateUniquePostTitle() string {
	return fmt.Sprintf("Test Post %s", generateRandomString(8))
}

// generateUniquePostContent generates unique post content for testing
func generateUniquePostContent() string {
	return fmt.Sprintf("This is a test post created at %s with random ID %s", 
		time.Now().Format(time.RFC3339), 
		generateRandomString(16))
} 
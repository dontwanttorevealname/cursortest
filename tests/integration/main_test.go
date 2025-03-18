package integration

import (
	"fmt"
	"log"
	"os"
	"testing"
	

	"github.com/playwright-community/playwright-go"
)

const (
	baseURL        = "http://localhost:8080"
	testUsername   = "testuser123"
	testPassword   = "Password123!"
	testPostTitle  = "Test Post Title"
	testPostContent = "This is a test post created by the integration test suite."
)

var (
	pw      *playwright.Playwright
	browser playwright.Browser
)

func TestMain(m *testing.M) {
	// Setup
	var err error
	pw, err = playwright.Run()
	if err != nil {
		log.Fatalf("Could not start Playwright: %v", err)
	}

	// Launch browser
	browser, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true), // Changed to true for headless mode
	})
	if err != nil {
		log.Fatalf("Could not launch browser: %v", err)
	}

	// Run tests
	code := m.Run()

	// Teardown
	if err = browser.Close(); err != nil {
		log.Fatalf("Could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("Could not stop Playwright: %v", err)
	}

	os.Exit(code)
}

// Helper function to create a new browser context and page
func newPage() (playwright.BrowserContext, playwright.Page, error) {
	// Create a new browser context for each test
	context, err := browser.NewContext()
	if err != nil {
		return nil, nil, fmt.Errorf("could not create browser context: %w", err)
	}

	// Create a new page
	page, err := context.NewPage()
	if err != nil {
		return nil, nil, fmt.Errorf("could not create page: %w", err)
	}

	// Set default timeout
	page.SetDefaultTimeout(30000)

	return context, page, nil
}

// Helper function to wait for navigation to complete
func waitForNavigation(page playwright.Page) error {
	return page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	})
}

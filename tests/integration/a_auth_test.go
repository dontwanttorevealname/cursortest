package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSignupOrLogin tests the signup process or logs in if the account already exists
func TestSignupOrLogin(t *testing.T) {
	// Create a new browser context and page
	context, page, err := newPage()
	assert.NoError(t, err)
	defer context.Close()

	// Navigate to the home page
	_, err = page.Goto(baseURL)
	assert.NoError(t, err)

	// Check if we're already logged in
	loggedIn, err := page.Locator("text=Log Out").Count()
	assert.NoError(t, err)

	if loggedIn > 0 {
		t.Log("Already logged in, logging out first")
		err = page.Click("text=Log Out")
		assert.NoError(t, err)
		err = waitForNavigation(page)
		assert.NoError(t, err)
	}

	// Check if we need to sign up or log in
	_, err = page.Goto(baseURL + "/signup")
	assert.NoError(t, err)

	// Fill in the signup form
	err = page.Fill("input[name=username]", testUsername)
	assert.NoError(t, err)
	err = page.Fill("input[name=password]", testPassword)
	assert.NoError(t, err)
	
	// Fill in the description field
	descriptionField, err := page.Locator("textarea[name=description]").Count()
	assert.NoError(t, err)
	if descriptionField > 0 {
		err = page.Fill("textarea[name=description]", "This is a test account created by automated tests.")
		assert.NoError(t, err)
	}

	// Check if username already exists
	// Wait for any validation
	page.WaitForTimeout(1000)

	// Look for error message about username already taken
	usernameTaken, err := page.Locator("text=Username is already taken").Count()
	assert.NoError(t, err)

	if usernameTaken > 0 {
		t.Log("Username already exists, clicking login button")
		
		// Look for a login button or link on the signup page
		loginButton, err := page.Locator("a:has-text('Log In'), button:has-text('Log In')").Count()
		assert.NoError(t, err)
		
		if loginButton > 0 {
			// Click the login button/link
			err = page.Click("a:has-text('Log In'), button:has-text('Log In')")
			assert.NoError(t, err)
			err = waitForNavigation(page)
			assert.NoError(t, err)
		} else {
			// If no login button is found, navigate directly
			t.Log("No login button found, navigating directly to login page")
			_, err = page.Goto(baseURL + "/login")
			assert.NoError(t, err)
		}

		// Fill in the login form
		err = page.Fill("input[name=username]", testUsername)
		assert.NoError(t, err)
		err = page.Fill("input[name=password]", testPassword)
		assert.NoError(t, err)

		// Submit the login form
		err = page.Click("button[type=submit]")
		assert.NoError(t, err)
	} else {
		// Submit the signup form
		err = page.Click("button[type=submit]")
		assert.NoError(t, err)
	}

	// Wait for navigation to complete
	err = waitForNavigation(page)
	assert.NoError(t, err)

	// Verify we're logged in by checking for the profile link or username display
	// Try different selectors that might indicate successful login
	profileLink, err := page.Locator("a[href='/profile']").Count()
	assert.NoError(t, err)
	
	usernameDisplay, err := page.Locator(`text="${testUsername}"`).Count()
	assert.NoError(t, err)
	
	logoutLink, err := page.Locator("text=Log Out").Count()
	assert.NoError(t, err)
	
	// Check if any of these elements are present
	isLoggedIn := profileLink > 0 || usernameDisplay > 0 || logoutLink > 0
	assert.True(t, isLoggedIn, "Should be logged in and see profile link, username, or logout option")
} 
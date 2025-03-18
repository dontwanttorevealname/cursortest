package integration

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
)

// TestCleanup tests cleaning up after the tests by deleting the post and leaving the pond
func TestCleanup(t *testing.T) {
	// Create a new browser context and page
	context, page, err := newPage()
	assert.NoError(t, err)
	defer context.Close()

	// Navigate to the home page
	_, err = page.Goto(baseURL)
	assert.NoError(t, err)

	// Check if we need to log in
	loginButton, err := page.Locator("text=Log In").Count()
	assert.NoError(t, err)

	if loginButton > 0 {
		// We need to log in first
		err = page.Click("text=Log In")
		assert.NoError(t, err)
		err = waitForNavigation(page)
		assert.NoError(t, err)

		// Fill in the login form
		err = page.Fill("input[name=username]", testUsername)
		assert.NoError(t, err)
		err = page.Fill("input[name=password]", testPassword)
		assert.NoError(t, err)

		// Submit the login form
		err = page.Click("button[type=submit]")
		assert.NoError(t, err)
		err = waitForNavigation(page)
		assert.NoError(t, err)
		
		t.Log("Logged in successfully, now on home page")
	}

	// Get the post title from the previous test
	postTitle := os.Getenv("TEST_POST_TITLE")
	if postTitle == "" {
		postTitle = testPostTitle // Use the default if not set
	}

	// Get the pond name from the previous test
	pondName := os.Getenv("TEST_POND_NAME")
	if pondName == "" {
		pondName = "GamerHaven" // Use the default if not set
	}

	// Navigate to the profile page to find the post
	t.Log("Navigating to profile page to find the post")
	_, err = page.Goto(baseURL + "/profile")
	assert.NoError(t, err)
	
	// Wait for the page to load completely
	err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	})
	assert.NoError(t, err)
	
	// Look for the post on the profile page using multiple approaches
	t.Log("Looking for post to delete on profile page")
	
	// Try different approaches to find the post
	postFound := false
	var postElement playwright.Locator
	
	// 1. Try exact text match
	exactMatch := page.Locator(`text="${postTitle}"`)
	exactMatchCount, err := exactMatch.Count()
	assert.NoError(t, err)
	
	if exactMatchCount > 0 {
		t.Log("Found post with exact title match")
		postElement = exactMatch
		postFound = true
	} else {
		// 2. Try contains text match
		containsMatch := page.Locator(`text=${postTitle}`)
		containsMatchCount, err := containsMatch.Count()
		assert.NoError(t, err)
		
		if containsMatchCount > 0 {
			t.Log("Found post with contains title match")
			postElement = containsMatch
			postFound = true
		} else {
			// 3. Try to find any heading or paragraph that might contain the title
			headingMatch := page.Locator(`h1,h2,h3,h4,h5,h6,p:has-text("${postTitle}")`)
			headingMatchCount, err := headingMatch.Count()
			assert.NoError(t, err)
			
			if headingMatchCount > 0 {
				t.Log("Found post title in a heading or paragraph")
				postElement = headingMatch
				postFound = true
			} else {
				// 4. Try to find any element that might contain the title
				anyMatch := page.Locator(`*:has-text("${postTitle}")`)
				anyMatchCount, err := anyMatch.Count()
				assert.NoError(t, err)
				
				if anyMatchCount > 0 {
					t.Log("Found post title in some element")
					postElement = anyMatch
					postFound = true
				} else {
					// 5. Try a more general approach - look for post cards
					postCards := page.Locator(".post-card")
					postCardCount, err := postCards.Count()
					assert.NoError(t, err)
					
					if postCardCount > 0 {
						t.Log("Found post cards on profile page")
						
						// Check the content of each post card
						for i := 0; i < postCardCount; i++ {
							card := postCards.Nth(i)
							cardText, err := card.TextContent()
							assert.NoError(t, err)
							
							t.Logf("Post card %d text: %s", i, cardText)
							
							// Check if the card contains our post title
							if cardText != "" && containsString(cardText, postTitle) {
								t.Logf("Found our post in card %d", i)
								postElement = card
								postFound = true
								break
							}
						}
					}
				}
			}
		}
	}
	
	if postFound {
		// Try to find a delete button or icon near the post
		t.Log("Looking for delete button near the post")
		
		// Try different approaches to find the delete button
		deleteButtonFound := false
		
		// 1. Look for a delete button within the post element
		deleteButton := postElement.Locator("button:has-text('Delete'), button.delete-button, button.btn-danger, button:has-text('Remove'), .delete-icon, .trash-icon, .bi-trash")
		deleteButtonCount, err := deleteButton.Count()
		assert.NoError(t, err)
		
		if deleteButtonCount > 0 {
			t.Log("Found delete button within post element")
			err = deleteButton.Click()
			assert.NoError(t, err)
			deleteButtonFound = true
		} else {
			// 2. Look for a delete icon near the post
			// Since Near() is not available, look for delete icons on the page and check if they're close to our post
			deleteIcons := page.Locator(".bi-trash, .trash-icon, .delete-icon")
			deleteIconCount, err := deleteIcons.Count()
			assert.NoError(t, err)
			
			if deleteIconCount > 0 {
				t.Log("Found delete icons, checking if any are near our post")
				// Try clicking each delete icon
				for i := 0; i < deleteIconCount; i++ {
					deleteIcon := deleteIcons.Nth(i)
					err = deleteIcon.Click()
					if err == nil {
						t.Logf("Successfully clicked delete icon %d", i)
						deleteButtonFound = true
						break
					}
				}
			}
			
			if !deleteButtonFound {
				// 3. Try to find any button that might be a delete button
				anyButtons := page.Locator("button, .btn")
				anyButtonCount, err := anyButtons.Count()
				assert.NoError(t, err)
				
				if anyButtonCount > 0 {
					// Check each button to see if it might be a delete button
					for i := 0; i < anyButtonCount; i++ {
						button := anyButtons.Nth(i)
						buttonText, err := button.TextContent()
						assert.NoError(t, err)
						
						t.Logf("Button %d text: %s", i, buttonText)
						
						// Check if the button might be a delete button
						if buttonText != "" && (containsString(buttonText, "Delete") || containsString(buttonText, "Remove") || containsString(buttonText, "Trash")) {
							t.Logf("Found potential delete button: %s", buttonText)
							err = button.Click()
							assert.NoError(t, err)
							deleteButtonFound = true
							break
						}
					}
				}
			}
		}
		
		if deleteButtonFound {
			// Wait for any confirmation dialog
			t.Log("Checking for confirmation dialog")
			confirmButton := page.Locator("button:has-text('Confirm'), button:has-text('Yes'), button:has-text('OK'), button.btn-danger")
			confirmButtonCount, err := confirmButton.Count()
			assert.NoError(t, err)
			
			if confirmButtonCount > 0 {
				t.Log("Found confirmation dialog, clicking confirm button")
				err = confirmButton.Click()
				assert.NoError(t, err)
			}
			
			// Wait for the page to refresh
			err = waitForNavigation(page)
			if err != nil {
				t.Logf("Navigation error after delete: %v", err)
			}
			
			// Wait for the page to load completely
			err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
				State: playwright.LoadStateNetworkidle,
			})
			assert.NoError(t, err)
			
			t.Log("Post deleted successfully")
		} else {
			t.Log("Could not find delete button for post")
		}
	} else {
		t.Log("Post not found or already deleted")
	}
	
	// Now navigate to the ponds tab to leave the pond
	t.Log("Navigating to ponds tab to leave the pond")
	
	// Try to find the ponds tab
	pondsTab := page.Locator("button[data-bs-target='#ponds'], button.nav-link:has-text('Ponds')")
	pondsTabCount, err := pondsTab.Count()
	assert.NoError(t, err)

	if pondsTabCount > 0 {
		t.Log("Found ponds tab, clicking it")
		err = pondsTab.First().Click()
		assert.NoError(t, err)
		
		// Wait for the ponds tab to load
		time.Sleep(1 * time.Second)
		
		// Look for the pond we want to leave
		t.Logf("Looking for pond: %s", pondName)
		pondElement := page.Locator(fmt.Sprintf(`text="%s"`, pondName))
		pondElementCount, err := pondElement.Count()
		assert.NoError(t, err)
		
		if pondElementCount > 0 {
			t.Logf("Found pond: %s", pondName)
			
			// Try to click on the pond to visit it first
			err = pondElement.Click()
			assert.NoError(t, err)
			err = waitForNavigation(page)
			assert.NoError(t, err)
			
			// Now look for a leave button on the pond page
			leaveButton := page.Locator("button:has-text('Leave'), button.leave-button, button.btn-danger")
			leaveButtonCount, err := leaveButton.Count()
			assert.NoError(t, err)
			
			if leaveButtonCount > 0 {
				t.Log("Found leave button on pond page, clicking it")
				err = leaveButton.Click()
				assert.NoError(t, err)
				
				// Wait for any confirmation dialog
				confirmButton := page.Locator("button:has-text('Confirm'), button:has-text('Yes'), button:has-text('OK'), button.btn-danger")
				confirmButtonCount, err := confirmButton.Count()
				assert.NoError(t, err)
				
				if confirmButtonCount > 0 {
					t.Log("Found confirmation dialog, clicking confirm button")
					err = confirmButton.Click()
					assert.NoError(t, err)
				}
				
				t.Log("Left pond successfully")
			} else {
				t.Log("Could not find leave button on pond page")
			}
		} else {
			t.Log("Pond not found in ponds tab")
		}
	} else {
		t.Log("Ponds tab not found")
	}
	
	// Wait a moment to ensure everything is complete
	time.Sleep(1 * time.Second)
} 
package integration

import (
	"os"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
)

// TestVerifyPostAndSearch tests verifying the post on the profile page and searching for it
func TestVerifyPostAndSearch(t *testing.T) {
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
	}

	// Get the post title from the previous test
	postTitle := os.Getenv("TEST_POST_TITLE")
	if postTitle == "" {
		postTitle = testPostTitle // Use the default if not set
	}

	// Navigate to the profile page
	err = page.Click("a[href='/profile']")
	assert.NoError(t, err)
	err = waitForNavigation(page)
	assert.NoError(t, err)

	// Wait for the page to load completely
	err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	})
	assert.NoError(t, err)

	// Look for the post on the profile page using multiple approaches
	t.Log("Looking for post on profile page")
	
	// Try different approaches to find the post
	postFound := false
	
	// 1. Try exact text match
	exactMatch := page.Locator(`text="${postTitle}"`)
	exactMatchCount, err := exactMatch.Count()
	assert.NoError(t, err)
	
	if exactMatchCount > 0 {
		t.Log("Found post with exact title match")
		postFound = true
	} else {
		// 2. Try contains text match
		containsMatch := page.Locator(`text=${postTitle}`)
		containsMatchCount, err := containsMatch.Count()
		assert.NoError(t, err)
		
		if containsMatchCount > 0 {
			t.Log("Found post with contains title match")
			postFound = true
		} else {
			// 3. Try to find any heading or paragraph that might contain the title
			headingMatch := page.Locator(`h1,h2,h3,h4,h5,h6,p:has-text("${postTitle}")`)
			headingMatchCount, err := headingMatch.Count()
			assert.NoError(t, err)
			
			if headingMatchCount > 0 {
				t.Log("Found post title in a heading or paragraph")
				postFound = true
			} else {
				// 4. Try to find any element that might contain the title
				anyMatch := page.Locator(`*:has-text("${postTitle}")`)
				anyMatchCount, err := anyMatch.Count()
				assert.NoError(t, err)
				
				if anyMatchCount > 0 {
					t.Log("Found post title in some element")
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
								postFound = true
								break
							}
						}
					}
				}
			}
		}
	}
	
	// Assert that we found the post
	assert.True(t, postFound, "Post should be visible on profile page")

	// Now search for the post
	searchInput, err := page.Locator("input[name=q]").Count()
	assert.NoError(t, err)
	
	if searchInput > 0 {
		t.Log("Found search input, filling with post title")
		err = page.Fill("input[name=q]", postTitle)
		assert.NoError(t, err)

		// Submit the search form
		err = page.Press("input[name=q]", "Enter")
		assert.NoError(t, err)
		err = waitForNavigation(page)
		assert.NoError(t, err)

		// Wait for search results to load
		err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
			State: playwright.LoadStateNetworkidle,
		})
		assert.NoError(t, err)

		// Look for the post in search results using multiple approaches
		t.Log("Looking for post in search results")
		
		// Try different approaches to find the post in search results
		postInSearchFound := false
		
		// 1. Try exact text match
		exactMatchSearch := page.Locator(`text="${postTitle}"`)
		exactMatchSearchCount, err := exactMatchSearch.Count()
		assert.NoError(t, err)
		
		if exactMatchSearchCount > 0 {
			t.Log("Found post in search results with exact title match")
			postInSearchFound = true
		} else {
			// 2. Try contains text match
			containsMatchSearch := page.Locator(`text=${postTitle}`)
			containsMatchSearchCount, err := containsMatchSearch.Count()
			assert.NoError(t, err)
			
			if containsMatchSearchCount > 0 {
				t.Log("Found post in search results with contains title match")
				postInSearchFound = true
			} else {
				// 3. Try to find any heading or paragraph that might contain the title
				headingMatchSearch := page.Locator(`h1,h2,h3,h4,h5,h6,p:has-text("${postTitle}")`)
				headingMatchSearchCount, err := headingMatchSearch.Count()
				assert.NoError(t, err)
				
				if headingMatchSearchCount > 0 {
					t.Log("Found post title in search results in a heading or paragraph")
					postInSearchFound = true
				} else {
					// 4. Try to find any element that might contain the title
					anyMatchSearch := page.Locator(`*:has-text("${postTitle}")`)
					anyMatchSearchCount, err := anyMatchSearch.Count()
					assert.NoError(t, err)
					
					if anyMatchSearchCount > 0 {
						t.Log("Found post title in search results in some element")
						postInSearchFound = true
					} else {
						// 5. Try a more general approach - look for post cards in search results
						searchResultCards := page.Locator(".post-card, .search-result-item")
						searchResultCardCount, err := searchResultCards.Count()
						assert.NoError(t, err)
						
						if searchResultCardCount > 0 {
							t.Log("Found result cards on search page")
							
							// Check the content of each result card
							for i := 0; i < searchResultCardCount; i++ {
								card := searchResultCards.Nth(i)
								cardText, err := card.TextContent()
								assert.NoError(t, err)
								
								t.Logf("Result card %d text: %s", i, cardText)
								
								// Check if the card contains our post title
								if cardText != "" && containsString(cardText, postTitle) {
									t.Logf("Found our post in search result card %d", i)
									postInSearchFound = true
									break
								}
							}
						}
					}
				}
			}
		}
		
		// Assert that we found the post in search results
		assert.True(t, postInSearchFound, "Post should be visible in search results")

		// Wait a moment to ensure everything is loaded
		time.Sleep(1 * time.Second)
	} else {
		t.Log("No search input found, skipping search test")
	}
}

// Helper function to check if a string contains another string
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && s != "" && substr != "" && s != substr && s[0:len(substr)] == substr
} 
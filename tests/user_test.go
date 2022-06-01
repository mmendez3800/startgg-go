// Package tests performs a review of the API wrapper methods implemented and confirms that their
// functionality can be utilized as needed.
package tests

import (
	"log"
	"os"
	"testing"

	"github.com/mmendez3800/startgg-go"
)

// TestUser checks the functionality of the methods and confirms that they provide a valid output.
func TestUser(t *testing.T) {
	// Sets the authentication token based on parameter passed for tests.
	startgg.SetAuthToken(os.Getenv("STARTGG_TOKEN"))

	// Implements a level of logging to indicate type.
	UserLogger := log.New(os.Stdout, "USER TESTS - ", log.LstdFlags)

	// Performs test for user object by ID.
	idResult, err := startgg.UserByID(41259)
	if err != nil {
		UserLogger.Fatalf("- ID Error - %s", err.Error())
	}
	UserLogger.Printf("- ID Results - %+v", idResult)

	// Performs test for user object by slug.
	slugResult, err := startgg.UserBySlug("user/fd58260e")
	if err != nil {
		UserLogger.Fatalf("- Slug Error - %s", err.Error())
	}
	UserLogger.Printf("- Slug Results - %+v\n\n", slugResult)
}

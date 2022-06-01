// Package tests performs a review of the API wrapper methods implemented and confirms that their
// functionality can be utilized as needed.
package tests

import (
	"log"
	"os"
	"testing"

	"github.com/mmendez3800/startgg-go"
)

// TestPlayer checks the functionality of the methods and confirms that they provide a valid output.
func TestPlayer(t *testing.T) {
	// Sets the authentication token based on parameter passed for tests.
	startgg.SetAuthToken(os.Getenv("STARTGG_TOKEN"))

	// Implements a level of logging to indicate type.
	PlayerLogger := log.New(os.Stdout, "PLAYER TESTS - ", log.LstdFlags)

	// Performs test for Player object by ID.
	idResult, err := startgg.PlayerByID(219999)
	if err != nil {
		PlayerLogger.Fatalf("- ID Error - %s", err.Error())
	}
	PlayerLogger.Printf("- ID Results - %+v\n\n", idResult)
}

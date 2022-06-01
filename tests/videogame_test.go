// Package tests performs a review of the API wrapper methods implemented and confirms that their
// functionality can be utilized as needed.
package tests

import (
	"log"
	"os"
	"testing"

	"github.com/mmendez3800/startgg-go"
)

// TestVideogame checks the functionality of the methods and confirms that they provide a valid output.
func TestVideogame(t *testing.T) {
	// Sets the authentication token based on parameter passed for tests.
	startgg.SetAuthToken(os.Getenv("STARTGG_TOKEN"))

	// Implements a level of logging to indicate type.
	VideogameLogger := log.New(os.Stdout, "VIDEOGAME TESTS - ", log.LstdFlags)

	// Performs test for Videogame object by ID.
	idResult, err := startgg.VideogameByID(1)
	if err != nil {
		VideogameLogger.Fatalf("- ID Error - %s", err.Error())
	}
	VideogameLogger.Printf("- ID Results - %+v", idResult)

	// Performs test for Videogame object by slug.
	slugResult, err := startgg.VideogameBySlug("game/ultimate-1")
	if err != nil {
		VideogameLogger.Fatalf("- Slug Error - %s", err.Error())
	}
	VideogameLogger.Printf("- Slug Results - %+v\n\n", slugResult)
}

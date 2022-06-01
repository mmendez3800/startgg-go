// Package tests performs a review of the API wrapper methods implemented and confirms that their
// functionality can be utilized as needed.
package tests

import (
	"log"
	"os"
	"testing"

	"github.com/mmendez3800/startgg-go"
)

// TestTournament checks the functionality of the methods and confirms that they provide a valid output.
func TestTournament(t *testing.T) {
	// Sets the authentication token based on parameter passed for tests.
	startgg.SetAuthToken(os.Getenv("STARTGG_TOKEN"))

	// Implements a level of logging to indicate type.
	TournamentLogger := log.New(os.Stdout, "TOURNAMENT TESTS - ", log.LstdFlags)

	// Performs test for Tournament object by ID.
	idResult, err := startgg.TournamentByID(326786)
	if err != nil {
		TournamentLogger.Fatalf("- ID Error - %s", err.Error())
	}
	TournamentLogger.Printf("- ID Results - %+v", idResult)

	// Performs test for Tournament object by slug.
	slugResult, err := startgg.TournamentBySlug("tournament/genesis-8")
	if err != nil {
		TournamentLogger.Fatalf("- Slug Error - %s", err.Error())
	}
	TournamentLogger.Printf("- Slug Results - %+v\n\n", slugResult)
}

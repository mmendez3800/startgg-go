package tests

import (
	"github.com/mmendez3800/startgg-go"
	"log"
	"os"
	"testing"
)

func TestTournament(t *testing.T) {
	startgg.SetAuthToken(os.Getenv("STARTGG_TOKEN"))

	TournamentLogger := log.New(os.Stdout, "TOURNAMENT TESTS - ", log.LstdFlags)

	idResult, err := startgg.TournamentByID(326786)
	if err != nil {
		TournamentLogger.Fatalf("- ID Error - %s", err.Error())
	}
	TournamentLogger.Printf("- ID Results - %+v", idResult)

	slugResult, err := startgg.TournamentBySlug("tournament/genesis-8")
	if err != nil {
		TournamentLogger.Fatalf("- Slug Error - %s", err.Error())
	}
	TournamentLogger.Printf("- Slug Results - %+v", slugResult)
}
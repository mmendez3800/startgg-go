package startgg

import (
	"github.com/mmendez3800/startgg-go/functions"
	"github.com/mmendez3800/startgg-go/objects"

)

func SetAuthToken(token string) {
	functions.AuthToken = token
}

func TournamentByID(id int64) (*objects.StartGGTournament, error) {
	return functions.GetTournament("id", id)
}

func TournamentBySlug(slug string) (*objects.StartGGTournament, error) {
	return functions.GetTournament("slug", slug)
}

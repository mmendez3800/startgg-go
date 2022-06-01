// Package startgg implements the core methods that can be used with the API wrapper.
package startgg

import (
	"github.com/mmendez3800/startgg-go/functions"
	"github.com/mmendez3800/startgg-go/objects"

)

// SetAuthToken allows the user to indicate the authentication token that is to be used
// when performing the calls.
func SetAuthToken(token string) {
	functions.AuthToken = token
}

// PlayerByID allows the user to pull data for a Player object based on the ID value indicated.
func PlayerByID(id int64) (*objects.StartGGPlayer, error) {
	return functions.GetPlayer("id", id)
}

// ShopByID allows the user to pull data for a Shop object based on the ID value indicated.
func ShopByID(id int64) (*objects.StartGGShop, error) {
	return functions.GetShop("id", id)
}

// ShopBySlug allows the user to pull data for a Shop object based on the slug value indicated.
func ShopBySlug(slug string) (*objects.StartGGShop, error) {
	return functions.GetShop("slug", slug)
}

// TournamentByID allows the user to pull data for a Tournament object based on the
// ID value indicated.
func TournamentByID(id int64) (*objects.StartGGTournament, error) {
	return functions.GetTournament("id", id)
}

// TournamentBySlug allows the user to pull data for a Tournament object based on the
// slug value indicated.
func TournamentBySlug(slug string) (*objects.StartGGTournament, error) {
	return functions.GetTournament("slug", slug)
}

// UserByID allows the user to pull data for a User object based on the ID value indicated.
func UserByID(id int64) (*objects.StartGGUser, error) {
	return functions.GetUser("id", id)
}

// UserBySlug allows the user to pull data for a User object based on the slug value indicated.
func UserBySlug(slug string) (*objects.StartGGUser, error) {
	return functions.GetUser("slug", slug)
}

// VideogameByID allows the user to pull data for a Videogame object based on the
// ID value indicated.
func VideogameByID(id int64) (*objects.StartGGVideogame, error) {
	return functions.GetVideogame("id", id)
}

// VideogameBySlug allows the user to pull data for a Videogame object based on the
// slug value indicated.
func VideogameBySlug(slug string) (*objects.StartGGVideogame, error) {
	return functions.GetVideogame("slug", slug)
}

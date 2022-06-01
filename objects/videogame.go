// Package objects defines the structures of the StartGG objects that can be called upon
// when performing the API calls.
package objects

type StartGGVideogame struct {
	DataVideogame DataVideogame `json:"data"`
	Errors []Errors `json:"errors"`
}

type DataVideogame struct {
	Videogame Videogame `json:"videogame"`
}

type Videogame struct {
	ID int64 `json:"id"`
	DisplayName string `json:"displayName"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

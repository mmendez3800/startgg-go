// Package objects defines the structures of the StartGG objects that can be called upon
// when performing the API calls.
package objects

type StartGGPlayer struct {
	DataPlayer DataPlayer `json:"data"`
	Errors []Errors `json:"errors"`
}

type DataPlayer struct {
	Player Player `json:"player"`
}

type Player struct {
	ID int64 `json:"id"`
	GamerTag string `json:"gamerTag"`
	Prefix string `json:"prefix"`
}

// Package objects defines the structures of the StartGG objects that can be called upon
// when performing the API calls.
package objects

type StartGGShop struct {
	DataShop DataShop `json:"data"`
	Errors []Errors `json:"errors"`
}

type DataShop struct {
	Shop Shop `json:"shop"`
}

type Shop struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	URL string `json:"url"`
}

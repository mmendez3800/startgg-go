// Package objects defines the structures of the StartGG objects that can be called upon
// when performing the API calls.
package objects

type StartGGUser struct {
	DataUser DataUser `json:"data"`
	Errors []Errors `json:"errors"`
}

type DataUser struct {
	User User `json:"user"`
}

type User struct {
	ID int64 `json:"id"`
	Bio string `json:"bio"`
	Birthday string `json:"birthday"`
	Discriminator string `json:"discriminator"`
	GenderPronoun string `json:"genderPronoun"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

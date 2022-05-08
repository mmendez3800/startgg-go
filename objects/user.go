package objects

type User struct {
	ID int `json:"id"`
	Bio string `json:"bio"`
	Birthday string `json:"birthday"`
	Discriminator string `json:"discriminator"`
	GenderPronoun string `json:"genderPronoun"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type UserAddress struct {
	Location Address `json:"location"`
}

type UserPlayser struct {
	Player Player `json:"player"`
}

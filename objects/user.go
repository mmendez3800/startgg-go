package objects

type User struct {
	id int
	bio string
	birthday string
	discriminator string
	genderPronoun string
	name string
	slug string
}

type UserAddress struct {
	location Address
}

type UserPlayser struct {
	player Player
}

package objects

type User struct {
	id int
	bio string
	birthday string
	discriminator string
	genderPronoun string
	location Address
	name string
	player Player
	slug string
}

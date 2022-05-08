package objects

type Player struct {
	ID int `json:"id"`
	GamerTag string `json:"gamerTag"`
	Prefix string `json:"prefix"`
}

type PlayerUser struct {
	User User `json:"user"`
}

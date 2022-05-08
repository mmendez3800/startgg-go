package objects

type Address struct {
	ID int `json:"id"`
	City string `json:"city"`
	Country string `json:"country"`
	CountryID int `json:"countryId"`
	State string `json:"state"`
	StateID int `json:"stateId"`
}

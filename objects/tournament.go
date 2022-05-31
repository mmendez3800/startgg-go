// Package objects defines the structures of the Start GG objects that can be called upon
// when performing the API calls.
package objects

type StartGGTournament struct {
	DataTournament DataTournament `json:"data"`
	Errors []Errors `json:"errors"`
}

type DataTournament struct {
	Tournament Tournament `json:"tournament"`
}

type Tournament struct {
	ID int64 `json:"id"`
	AddrState string `json:"addrState"`
	City string `json:"city"`
	CountryCode string `json:"countryCode"`
	CreatedAt int64 `json:"createdAt"`
	Currency string `json:"currency"`
	EndAt int64 `json:"endAt"`
	EventRegistrationClosesAt int64 `json:"eventRegistrationClosesAt"`
	HasOfflineEvents bool `json:"hasOfflineEvents"`
	HasOnlineEvents bool `json:"hasOnlineEvents"`
	Hashtag string `json:"hashtag"`
	IsOnline bool `json:"isOnline"`
	IsRegistrationOpen bool `json:"isRegistrationOpen"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	MapsPlaceID string `json:"mapsPlaceId"`
	Name string `json:"name"`
	NumAttendees int64 `json:"numAttendees"`
	PostalCode string `json:"postalCode"`
	PrimaryContact string `json:"primaryContact"`
	PrimaryContactType string `json:"primaryContactType"`
	Publishing map[string]interface{} `json:"publishing"`
	RegistrationClosesAt int64 `json:"registrationClosesAt"`
	Rules string `json:"rules"`
	ShortSlug string `json:"shortSlug"`
	Slug string `json:"slug"`
	StartAt int64 `json:"startAt"`
	State int64 `json:"state"`
	TeamCreationClosesAt int64 `json:"teamCreationClosesAt"`
	Timezone string `json:"timezone"`
	TournamentType int64 `json:"tournamentType"`
	UpdatedAt int64 `json:"updatedAt"`
	URL string `json:"url"`
	VenueAddress string `json:"venueAddress"`
	VenueName string `json:"venueName"`
}

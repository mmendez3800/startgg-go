package queries

const (
	TournamentQuery = `query TournamentQuery($id: ID, $slug: String) {
		tournament(id: $id, slug: $slug) {
			id
			addrState
			endAt
			eventRegistrationClosesAt
			hasOfflineEvents
			hasOnlineEvents
			hashtag
			isOnline
			isRegistrationOpen
			lat
			lng
			mapsPlaceId
			name
			numAttendees
			postalCode
			primaryContact
			primaryContactType
			publishing
			registrationClosesAt
			rules
			shortSlug
			slug
			startAt
			state
			teamCreationClosesAt
			timezone
			tournamentType
			updatedAt
			url
			venueAddress
			venueName
		}
	}`
)

// Package queries stores the list of GraphQL queries that can be run on the Start GG
// API and indicates the possible values that can be passed.
package queries

const (
	TournamentQuery = `query TournamentQuery($id: ID, $slug: String) {
		tournament(id: $id, slug: $slug) {
			id
			addrState
			city
			countryCode
			createdAt
			currency
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

package queries

const (
	UserQuery = `query UserQuery($id: ID) {
		user(id: $id) {
			id
			bio
			birthday
			discriminator
			genderPronoun
			name
			slug
		}
	}`
	UserAddressQuery = `query UserAddressQuery($id: ID) {
		user(id: $id) {
			address {
				id
				city
				country
				countryId
				state
				stateId
			}
		}
	}`
	UserPlayerQuery = `query UserPlayerQuery($id: ID) {
		user(id: $id) {
			player {
				id
				gamerTag
				prefix
			}
		}
	}`
)

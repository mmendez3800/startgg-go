package queries

const (
	PlayerQuery = `query PlayerQuery($id: ID) {
		player(id: $id) {
			id
			gamerTag
			prefix
		}
	}`
	PlayerUserQuery = `query PlayerUserQuery($id: ID) {
		player(id: $id) {
			user {
				id
				bio
				birthday
				discriminator
				genderPronoun
				name
				slug
			}
		}
	}`
)

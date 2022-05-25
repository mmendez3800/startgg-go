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
)

// Package queries stores the list of GraphQL queries that can be run on the Start GG
// API and indicates the possible values that can be passed.
package queries

const (
	UserQuery = `query UserQuery($id: ID, $slug: String) {
		user(id: $id, slug: $slug) {
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

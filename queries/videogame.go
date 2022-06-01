// Package queries stores the list of GraphQL queries that can be run on the StartGG
// API and indicates the possible values that can be passed.
package queries

const (
	VideogameQuery = `query VideogameQuery($id: ID, $slug: String) {
		videogame(id: $id, slug: $slug) {
			id
			displayName
			name
			slug
		}
	}`
)

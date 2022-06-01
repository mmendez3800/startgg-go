// Package queries stores the list of GraphQL queries that can be run on the StartGG
// API and indicates the possible values that can be passed.
package queries

const (
	ShopQuery = `query ShopQuery($id: ID, $slug: String) {
		shop(id: $id, slug: $slug) {
			id
			name
			slug
			url
		}
	}`
)

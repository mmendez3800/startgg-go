// Package queries stores the list of GraphQL queries that can be run on the Start GG
// API and indicates the possible values that can be passed.
package queries

const (
	PlayerQuery = `query PlayerQuery($id: ID!) {
		player(id: $id) {
			id
			gamerTag
			prefix
		}
	}`
)

package queries

const (
	PlayerQuery = `query PlayerQuery($id: ID) {
		player(id: $id) {
			id
			gamerTag
			prefix
		}
	}`
)

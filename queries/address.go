package queries

const (
	AddressQuery = `query AddressQuery($id: ID) {
		address(id: $id) {
			id
			city
			country
			countryId
			state
			stateId
		}
	}`
)

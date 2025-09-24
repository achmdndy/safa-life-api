package query

// ToGetAllRecitersQuery creates a query for getting all reciters.
func ToGetAllRecitersQuery() GetAllRecitersQuery {
	return GetAllRecitersQuery{}
}

// ToGetReciterByIDQuery creates a query for getting a reciter by ID.
func ToGetReciterByIDQuery(id string) GetReciterByIDQuery {
	return GetReciterByIDQuery{ID: id}
}

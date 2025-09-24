package query

func ToGetAllResourcesQuery() GetAllResourcesQuery {
	return GetAllResourcesQuery{}
}

func ToGetResourceByIDQuery(id string) GetResourceByIDQuery {
	return GetResourceByIDQuery{ID: id}
}

package query

func ToGetAllStoriesQuery() GetAllStoriesQuery {
	return GetAllStoriesQuery{}
}

func ToGetStoryByIDQuery(id string) GetStoryByIDQuery {
	return GetStoryByIDQuery{ID: id}
}

func ToGetAyahsForStoryQuery(storyID string) GetAyahsForStoryQuery {
	return GetAyahsForStoryQuery{StoryID: storyID}
}

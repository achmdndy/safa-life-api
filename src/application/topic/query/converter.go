package query

func ToGetAllTopicsQuery() GetAllTopicsQuery {
	return GetAllTopicsQuery{}
}

func ToGetTopicByIDQuery(id string) GetTopicByIDQuery {
	return GetTopicByIDQuery{ID: id}
}

func ToGetAyahsForTopicQuery(topicID string) GetAyahsForTopicQuery {
	return GetAyahsForTopicQuery{TopicID: topicID}
}

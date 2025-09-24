package query

import (
	"github.com/achmdndy/safa-life-api/src/domain/topic"
)

// QueryHandler handles all read operations for the Topic module.
type QueryHandler struct {
	topicService topic.TopicService
}

// NewQueryHandler creates a new query handler.
func NewQueryHandler(service topic.TopicService) *QueryHandler {
	return &QueryHandler{
		topicService: service,
	}
}

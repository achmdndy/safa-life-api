package query

import (
	"github.com/achmdndy/safa-life-api/src/domain/story"
)

// QueryHandler handles all read operations for the Story module.
type QueryHandler struct {
	storyService story.StoryService
}

// NewQueryHandler creates a new query handler.
func NewQueryHandler(service story.StoryService) *QueryHandler {
	return &QueryHandler{
		storyService: service,
	}
}

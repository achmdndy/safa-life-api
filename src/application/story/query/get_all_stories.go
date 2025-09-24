package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/story/dto"
)

type GetAllStoriesQuery struct{}

func (h *QueryHandler) GetAllStories(ctx context.Context, query GetAllStoriesQuery) ([]dto.StoryResponse, error) {
	domains, err := h.storyService.GetAllStories(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ToStoryResponseSlice(domains), nil
}

package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/story/dto"
)

type GetAyahsForStoryQuery struct {
	StoryID string
}

func (h *QueryHandler) GetAyahsForStory(ctx context.Context, query GetAyahsForStoryQuery) ([]dto.StoryAyahResponse, error) {
	domains, err := h.storyService.GetAyahsForStory(ctx, query.StoryID)
	if err != nil {
		return nil, err
	}
	return dto.ToStoryAyahResponseSlice(domains), nil
}

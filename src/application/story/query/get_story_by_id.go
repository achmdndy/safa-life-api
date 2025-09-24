package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/story/dto"
)

type GetStoryByIDQuery struct {
	ID string
}

func (h *QueryHandler) GetStoryByID(ctx context.Context, query GetStoryByIDQuery) (*dto.StoryResponse, error) {
	domain, err := h.storyService.GetStoryByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	resp := dto.ToStoryResponse(*domain)
	return &resp, nil
}

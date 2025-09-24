package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/topic/dto"
)

type GetTopicByIDQuery struct {
	ID string
}

func (h *QueryHandler) GetTopicByID(ctx context.Context, query GetTopicByIDQuery) (*dto.TopicResponse, error) {
	domain, err := h.topicService.GetTopicByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	resp := dto.ToTopicResponse(*domain)
	return &resp, nil
}

package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/topic/dto"
)

type GetAllTopicsQuery struct{}

func (h *QueryHandler) GetAllTopics(ctx context.Context, query GetAllTopicsQuery) ([]dto.TopicResponse, error) {
	domains, err := h.topicService.GetAllTopics(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ToTopicResponseSlice(domains), nil
}

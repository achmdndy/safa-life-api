package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/topic/dto"
)

type GetAyahsForTopicQuery struct {
	TopicID string
}

func (h *QueryHandler) GetAyahsForTopic(ctx context.Context, query GetAyahsForTopicQuery) ([]dto.TopicAyahResponse, error) {
	domains, err := h.topicService.GetAyahsForTopic(ctx, query.TopicID)
	if err != nil {
		return nil, err
	}
	return dto.ToTopicAyahResponseSlice(domains), nil
}

package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type GetLastReadByIdQuery struct {
	ID      string
	Include string
}

func (h *QueryHandler) GetLastReadById(ctx context.Context, query GetLastReadByIdQuery) (interface{}, error) {
	id, err := h.uuidGenerator.Parse(query.ID)
	if err != nil {
		return nil, err
	}

	if query.Include == "relations" {
		var lrWith *quran.LastReadWithRelations
		lrWith, err = h.lastReadService.GetLastReadByIdWithRelations(ctx, id)
		if err != nil {
			return nil, err
		}
		return dto.ToLastReadWithRelationsResponse(lrWith), nil
	}

	var lr *quran.LastRead
	lr, err = h.lastReadService.GetLastReadById(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToLastReadResponse(lr), nil
}

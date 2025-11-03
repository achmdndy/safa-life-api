package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetJuzByNumberQuery struct {
	Number  int
	Include string
}

func (h *QueryHandler) GetJuzByNumber(ctx context.Context, query GetJuzByNumberQuery) (interface{}, error) {
	// Check if eager loading is requested
	if query.Include == "relations" {
		juz, err := h.juzService.GetJuzByNumberWithRelations(ctx, query.Number)
		if err != nil {
			return nil, err
		}
		return dto.ToJuzWithRelationsResponse(juz), nil
	}

	// Default behavior without eager loading
	juz, err := h.juzService.GetJuzByNumber(ctx, query.Number)
	if err != nil {
		return nil, err
	}

	return dto.ToJuzResponse(juz), nil
}

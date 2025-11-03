package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type GetJuzByIdQuery struct {
	ID      string
	Include string
}

func (h *QueryHandler) GetJuzById(ctx context.Context, query GetJuzByIdQuery) (interface{}, error) {
	id, err := h.uuidGenerator.Parse(query.ID)
	if err != nil {
		return nil, err
	}

	// Check if eager loading is requested
	if query.Include == "relations" {
		var juzWithRelations *quran.JuzWithRelations
		juzWithRelations, err = h.juzService.GetJuzByIdWithRelations(ctx, id)
		if err != nil {
			return nil, err
		}
		return dto.ToJuzWithRelationsResponse(juzWithRelations), nil
	}

	// Default behavior without eager loading
	var juz *quran.Juz
	juz, err = h.juzService.GetJuzById(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.ToJuzResponse(juz), nil
}

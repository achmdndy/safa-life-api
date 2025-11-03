package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAyahByIdQuery struct {
	ID string
}

func (h *QueryHandler) GetAyahById(ctx context.Context, query GetAyahByIdQuery) (*dto.AyahResponse, error) {
	id, err := h.uuidGenerator.Parse(query.ID)
	if err != nil {
		return nil, err
	}

	ayah, err := h.ayahService.GetAyahById(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.ToAyahResponse(ayah), nil
}

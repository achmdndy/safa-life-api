package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetSurahByIdQuery struct {
	ID string
}

func (h *QueryHandler) GetSurahById(ctx context.Context, query GetSurahByIdQuery) (*dto.SurahResponse, error) {
	id, err := h.uuidGenerator.Parse(query.ID)
	if err != nil {
		return nil, err
	}

	surah, err := h.surahService.GetSurahById(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.ToSurahResponse(surah), nil
}

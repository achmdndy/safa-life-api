package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAyahTranslationsBySurahAndEditionQuery struct {
	SurahID   string
	EditionID string
	Limit     int
	Offset    int
}

func (h *QueryHandler) GetAyahTranslationsBySurahAndEdition(ctx context.Context, query GetAyahTranslationsBySurahAndEditionQuery) (*dto.AyahTranslationListResponse, error) {
	surahID, err := h.uuidGenerator.Parse(query.SurahID)
	if err != nil {
		return nil, err
	}

	editionID, err := h.uuidGenerator.Parse(query.EditionID)
	if err != nil {
		return nil, err
	}

	items, err := h.ayahTranslationService.GetAyahTranslationsBySurahAndEdition(ctx, surahID, editionID, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	count, err := h.ayahTranslationService.CountAyahTranslationsBySurahAndEdition(ctx, surahID, editionID)
	if err != nil {
		return nil, err
	}

	return &dto.AyahTranslationListResponse{
		Data: dto.ToAyahTranslationResponseSlice(items),
		Pagination: &dto.PaginationResponse{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  count,
		},
	}, nil
}

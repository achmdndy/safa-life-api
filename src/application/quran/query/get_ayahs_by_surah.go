package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAyahsBySurahQuery struct {
	SurahID string
	Limit   int
	Offset  int
}

func (h *QueryHandler) GetAyahsBySurah(ctx context.Context, query GetAyahsBySurahQuery) (*dto.AyahListResponse, error) {
	surahID, err := h.uuidGenerator.Parse(query.SurahID)
	if err != nil {
		return nil, err
	}

	ayahs, err := h.ayahService.GetAyahsBySurahId(ctx, surahID, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	count, err := h.ayahService.CountAyahsBySurahId(ctx, surahID)
	if err != nil {
		return nil, err
	}

	return &dto.AyahListResponse{
		Data: dto.ToAyahResponseSlice(ayahs),
		Pagination: &dto.PaginationResponse{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  count,
		},
	}, nil
}

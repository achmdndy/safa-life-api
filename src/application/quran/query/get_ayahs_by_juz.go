package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAyahsByJuzQuery struct {
	JuzNumber int
	Limit     int
	Offset    int
}

func (h *QueryHandler) GetAyahsByJuz(ctx context.Context, query GetAyahsByJuzQuery) (*dto.AyahListResponse, error) {
	ayahs, err := h.ayahService.GetAyahsByJuzNumber(ctx, query.JuzNumber, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	count, err := h.ayahService.CountAyahsByJuzNumber(ctx, query.JuzNumber)
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

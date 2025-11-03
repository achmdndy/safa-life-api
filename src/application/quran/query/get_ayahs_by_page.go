package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAyahsByPageQuery struct {
	PageNumber int
	Limit      int
	Offset     int
}

func (h *QueryHandler) GetAyahsByPage(ctx context.Context, query GetAyahsByPageQuery) (*dto.AyahListResponse, error) {
	ayahs, err := h.ayahService.GetAyahsByPageNumber(ctx, query.PageNumber, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	count, err := h.ayahService.CountAyahsByPageNumber(ctx, query.PageNumber)
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

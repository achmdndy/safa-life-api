package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAllAyahsQuery struct {
	Limit  int
	Offset int
}

func (h *QueryHandler) GetAllAyahs(ctx context.Context, query GetAllAyahsQuery) (*dto.AyahListResponse, error) {
	ayahs, err := h.ayahService.GetAllAyahs(ctx, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	count, err := h.ayahService.CountAyahs(ctx)
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

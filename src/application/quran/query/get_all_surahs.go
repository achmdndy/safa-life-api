package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAllSurahsQuery struct {
	Limit  int
	Offset int
}

func (h *QueryHandler) GetAllSurahs(ctx context.Context, query GetAllSurahsQuery) (*dto.SurahListResponse, error) {
	surahs, err := h.surahService.GetAllSurahs(ctx, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	count, err := h.surahService.CountSurahs(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.SurahListResponse{
		Data: dto.ToSurahResponseSlice(surahs),
		Pagination: &dto.PaginationResponse{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  count,
		},
	}, nil
}

package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAllRecitersQuery struct {
	Limit  int
	Offset int
}

func (h *QueryHandler) GetAllReciters(ctx context.Context, q GetAllRecitersQuery) (*dto.ReciterListResponse, error) {
	items, err := h.reciterService.GetAllReciters(ctx, q.Limit, q.Offset)
	if err != nil {
		return nil, err
	}
	count, err := h.reciterService.CountReciters(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.ReciterListResponse{
		Data: dto.ToReciterResponseSlice(items),
		Pagination: &dto.PaginationResponse{
			Total:  count,
			Limit:  q.Limit,
			Offset: q.Offset,
		},
	}, nil
}

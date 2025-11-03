package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAllJuzQuery struct {
	Limit   int
	Offset  int
	Include string
}

func (h *QueryHandler) GetAllJuz(ctx context.Context, query GetAllJuzQuery) (interface{}, error) {
	// Check if eager loading is requested
	if query.Include == "relations" {
		juzList, err := h.juzService.GetAllJuzWithRelations(ctx, query.Limit, query.Offset)
		if err != nil {
			return nil, err
		}

		count, err := h.juzService.CountJuz(ctx)
		if err != nil {
			return nil, err
		}

		return &dto.JuzWithRelationsListResponse{
			Data: dto.ToJuzWithRelationsResponseSlice(juzList),
			Pagination: &dto.PaginationResponse{
				Limit:  query.Limit,
				Offset: query.Offset,
				Total:  count,
			},
		}, nil
	}

	// Default behavior without eager loading
	juzList, err := h.juzService.GetAllJuz(ctx, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	count, err := h.juzService.CountJuz(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.JuzListResponse{
		Data: dto.ToJuzResponseSlice(juzList),
		Pagination: &dto.PaginationResponse{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  count,
		},
	}, nil
}

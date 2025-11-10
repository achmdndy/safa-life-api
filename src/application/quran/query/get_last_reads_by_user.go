package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type GetLastReadsByUserQuery struct {
	UserID  string
	Limit   int
	Offset  int
	Include string
}

func (h *QueryHandler) GetLastReadsByUser(ctx context.Context, query GetLastReadsByUserQuery) (interface{}, error) {
	var err error

	if query.Include == "relations" {
		var listWith []*quran.LastReadWithRelations
		var count int64
		listWith, err = h.lastReadService.GetLastReadsByUserWithRelations(ctx, query.UserID, query.Limit, query.Offset)
		if err != nil {
			return nil, err
		}
		count, err = h.lastReadService.CountLastReadsByUser(ctx, query.UserID)
		if err != nil {
			return nil, err
		}
		return &dto.LastReadWithRelationsListResponse{
			Data: dto.ToLastReadWithRelationsResponseSlice(listWith),
			Pagination: &dto.PaginationResponse{
				Limit:  query.Limit,
				Offset: query.Offset,
				Total:  count,
			},
		}, nil
	}

	var list []*quran.LastRead
	var count int64
	list, err = h.lastReadService.GetLastReadsByUser(ctx, query.UserID, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}
	count, err = h.lastReadService.CountLastReadsByUser(ctx, query.UserID)
	if err != nil {
		return nil, err
	}
	return &dto.LastReadListResponse{
		Data: dto.ToLastReadResponseSlice(list),
		Pagination: &dto.PaginationResponse{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  count,
		},
	}, nil
}

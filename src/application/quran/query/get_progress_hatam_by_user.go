package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type GetProgressHatamByUserQuery struct {
	UserID  string
	Limit   int
	Offset  int
	Include string
}

func (h *QueryHandler) GetProgressHatamByUser(ctx context.Context, query GetProgressHatamByUserQuery) (interface{}, error) {
	userID, err := h.uuidGenerator.Parse(query.UserID)
	if err != nil {
		return nil, err
	}

	if query.Include == "relations" {
		var listWith []*quran.ProgressHatamWithRelations
		var count int64
		listWith, err = h.progressService.GetProgressHatamByUserWithRelations(ctx, userID, query.Limit, query.Offset)
		if err != nil {
			return nil, err
		}
		count, err = h.progressService.CountProgressHatamByUser(ctx, userID)
		if err != nil {
			return nil, err
		}
		return &dto.ProgressHatamWithRelationsListResponse{
			Data: dto.ToProgressHatamWithRelationsResponseSlice(listWith),
			Pagination: &dto.PaginationResponse{
				Limit:  query.Limit,
				Offset: query.Offset,
				Total:  count,
			},
		}, nil
	}

	var list []*quran.ProgressHatam
	var count int64
	list, err = h.progressService.GetProgressHatamByUser(ctx, userID, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}
	count, err = h.progressService.CountProgressHatamByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &dto.ProgressHatamListResponse{
		Data: dto.ToProgressHatamResponseSlice(list),
		Pagination: &dto.PaginationResponse{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  count,
		},
	}, nil
}

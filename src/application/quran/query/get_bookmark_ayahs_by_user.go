package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetBookmarkAyahsByUserQuery struct {
	UserID string
	Limit  int
	Offset int
}

func (h *QueryHandler) GetBookmarkAyahsByUser(ctx context.Context, query GetBookmarkAyahsByUserQuery) (*dto.BookmarkAyahListResponse, error) {
	userID, err := h.uuidGenerator.Parse(query.UserID)
	if err != nil {
		return nil, err
	}

	bookmarks, err := h.bookmarkService.GetBookmarkAyahsByUser(ctx, userID, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	count, err := h.bookmarkService.CountBookmarkAyahsByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dto.BookmarkAyahListResponse{
		Data: dto.ToBookmarkAyahResponseSlice(bookmarks),
		Pagination: &dto.PaginationResponse{
			Limit:  query.Limit,
			Offset: query.Offset,
			Total:  count,
		},
	}, nil
}

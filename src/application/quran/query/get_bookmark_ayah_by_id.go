package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type GetBookmarkAyahByIdQuery struct {
	ID      string
	Include string
}

// GetBookmarkAyahById retrieves a bookmark by ID, optionally with related ayah
func (h *QueryHandler) GetBookmarkAyahById(ctx context.Context, query GetBookmarkAyahByIdQuery) (interface{}, error) {
	id, err := h.uuidGenerator.Parse(query.ID)
	if err != nil {
		return nil, err
	}

	if query.Include == "ayah" {
		var bWith *quran.BookmarkAyahWithAyah
		bWith, err = h.bookmarkService.GetBookmarkAyahByIdWithAyah(ctx, id)
		if err != nil {
			return nil, err
		}
		return dto.ToBookmarkAyahWithAyahResponse(bWith), nil
	}

	var b *quran.BookmarkAyah
	b, err = h.bookmarkService.GetBookmarkAyahById(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToBookmarkAyahResponse(b), nil
}

package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type GetBookmarkAyahByUserAndAyahQuery struct {
	UserID  string
	AyahID  string
	Include string
}

// GetBookmarkAyahByUserAndAyah retrieves a bookmark by user and ayah; if include=ayah, it returns with relation
func (h *QueryHandler) GetBookmarkAyahByUserAndAyah(ctx context.Context, query GetBookmarkAyahByUserAndAyahQuery) (interface{}, error) {
	userID, err := h.uuidGenerator.Parse(query.UserID)
	if err != nil {
		return nil, err
	}
	ayahID, err := h.uuidGenerator.Parse(query.AyahID)
	if err != nil {
		return nil, err
	}

	var b *quran.BookmarkAyah
	b, err = h.bookmarkService.GetBookmarkAyahByUserAndAyah(ctx, userID, ayahID)
	if err != nil {
		return nil, err
	}

	if query.Include == "ayah" && b != nil {
		var bWith *quran.BookmarkAyahWithAyah
		bWith, err = h.bookmarkService.GetBookmarkAyahByIdWithAyah(ctx, b.ID)
		if err != nil {
			return nil, err
		}
		return dto.ToBookmarkAyahWithAyahResponse(bWith), nil
	}

	return dto.ToBookmarkAyahResponse(b), nil
}

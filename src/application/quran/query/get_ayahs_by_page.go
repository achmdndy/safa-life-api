package query

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// GetAyahsByPageQuery represents the query to get ayahs by page
type GetAyahsByPageQuery struct {
	Page int
}

// GetAyahsByPage retrieves ayahs by page number
func (h *QueryHandler) GetAyahsByPage(ctx context.Context, query GetAyahsByPageQuery) (*dto.PaginatedResponse[dto.AyahResponse], error) {
	if query.Page < 1 || query.Page > 604 {
		return nil, fmt.Errorf("invalid page number: %d", query.Page)
	}

	ayahs, err := h.quranService.GetAyahsByPage(ctx, query.Page)
	if err != nil {
		return nil, err
	}

	// Convert to response DTOs
	ayahResponses := dto.ToAyahResponseSlice(ayahs)

	return &dto.PaginatedResponse[dto.AyahResponse]{
		Data: ayahResponses,
		Pagination: dto.PaginationMetadata{
			Page:       query.Page,
			Limit:      len(ayahResponses),
			Total:      len(ayahResponses),
			TotalPages: 1,
		},
	}, nil
}
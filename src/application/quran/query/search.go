package query

import (
	"context"
	"fmt"
	"math"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// SearchQuery represents the query to search for ayahs
type SearchQuery struct {
	Query string
	Page  int
	Limit int
}

// Search searches for ayahs based on the query string
func (h *QueryHandler) Search(ctx context.Context, query SearchQuery) (*dto.PaginatedResponse[dto.AyahResponse], error) {
	if query.Query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	ayahs, err := h.quranService.SearchAyahs(ctx, query.Query)
	if err != nil {
		return nil, err
	}

	// Convert to response DTOs
	ayahResponses := dto.ToAyahResponseSlice(ayahs)

	// Apply pagination if specified
	if query.Page > 0 && query.Limit > 0 {
		start := (query.Page - 1) * query.Limit
		end := start + query.Limit

		if start >= len(ayahResponses) {
			ayahResponses = []dto.AyahResponse{}
		} else {
			if end > len(ayahResponses) {
				end = len(ayahResponses)
			}
			ayahResponses = ayahResponses[start:end]
		}

		totalPages := int(math.Ceil(float64(len(ayahs)) / float64(query.Limit)))

		return &dto.PaginatedResponse[dto.AyahResponse]{
			Data: ayahResponses,
			Pagination: dto.PaginationMetadata{
				Page:       query.Page,
				Limit:      query.Limit,
				Total:      len(ayahs),
				TotalPages: totalPages,
			},
		}, nil
	}

	return &dto.PaginatedResponse[dto.AyahResponse]{
		Data: ayahResponses,
		Pagination: dto.PaginationMetadata{
			Page:       1,
			Limit:      len(ayahResponses),
			Total:      len(ayahResponses),
			TotalPages: 1,
		},
	}, nil
}
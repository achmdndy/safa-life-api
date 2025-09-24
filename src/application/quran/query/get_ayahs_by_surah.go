package query

import (
	"context"
	"fmt"
	"math"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// GetAyahsBySurahQuery represents the query to get ayahs by surah
type GetAyahsBySurahQuery struct {
	SurahID int
	Page    int
	Limit   int
}

// GetAyahsBySurah retrieves ayahs by surah ID with optional pagination
func (h *QueryHandler) GetAyahsBySurah(ctx context.Context, query GetAyahsBySurahQuery) (*dto.PaginatedResponse[dto.AyahResponse], error) {
	if query.SurahID < 1 || query.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", query.SurahID)
	}

	ayahs, err := h.quranService.GetAyahsBySurah(ctx, query.SurahID)
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
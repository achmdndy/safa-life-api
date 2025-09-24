package query

import (
	"context"
	"math"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// GetAllSurahsQuery represents the query to get all surahs
type GetAllSurahsQuery struct {
	Page  int
	Limit int
}

// GetAllSurahs retrieves all surahs with optional pagination
func (h *QueryHandler) GetAllSurahs(ctx context.Context, query GetAllSurahsQuery) (*dto.PaginatedResponse[dto.SurahResponse], error) {
	surahs, err := h.quranService.GetAllSurahs(ctx)
	if err != nil {
		return nil, err
	}

	// Convert to response DTOs
	surahResponses := dto.ToSurahResponseSlice(surahs)

	// Apply pagination if specified
	if query.Page > 0 && query.Limit > 0 {
		start := (query.Page - 1) * query.Limit
		end := start + query.Limit

		if start >= len(surahResponses) {
			surahResponses = []dto.SurahResponse{}
		} else {
			if end > len(surahResponses) {
				end = len(surahResponses)
			}
			surahResponses = surahResponses[start:end]
		}

		totalPages := int(math.Ceil(float64(len(surahs)) / float64(query.Limit)))

		return &dto.PaginatedResponse[dto.SurahResponse]{
			Data: surahResponses,
			Pagination: dto.PaginationMetadata{
				Page:       query.Page,
				Limit:      query.Limit,
				Total:      len(surahs),
				TotalPages: totalPages,
			},
		}, nil
	}

	return &dto.PaginatedResponse[dto.SurahResponse]{
		Data: surahResponses,
		Pagination: dto.PaginationMetadata{
			Page:       1,
			Limit:      len(surahResponses),
			Total:      len(surahResponses),
			TotalPages: 1,
		},
	}, nil
}
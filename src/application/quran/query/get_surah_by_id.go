package query

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// GetSurahByIDQuery represents the query to get a surah by ID
type GetSurahByIDQuery struct {
	ID int
}

// GetSurahByID retrieves a surah by its ID
func (h *QueryHandler) GetSurahByID(ctx context.Context, query GetSurahByIDQuery) (*dto.SurahResponse, error) {
	if query.ID < 1 || query.ID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", query.ID)
	}

	surah, err := h.quranService.GetSurahByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	response := dto.ToSurahResponse(*surah)
	return &response, nil
}
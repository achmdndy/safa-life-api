package query

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
)

// GetAyahTajweedQuery represents the query to get ayah tajweed.
type GetAyahTajweedQuery struct {
	TajweedID string
	SurahID   int
	AyahID    int
}

// GetAyahTajweed retrieves ayah tajweed.
func (h *QueryHandler) GetAyahTajweed(ctx context.Context, query GetAyahTajweedQuery) (*dto.AyahTajweedResponse, error) {
	if query.TajweedID == "" {
		return nil, fmt.Errorf("tajweed ID cannot be empty")
	}
	if query.SurahID < 1 || query.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", query.SurahID)
	}
	if query.AyahID < 1 {
		return nil, fmt.Errorf("invalid ayah ID: %d", query.AyahID)
	}

	ayahTajweed, err := h.tajweedService.GetAyahTajweed(ctx, query.TajweedID, query.SurahID, query.AyahID)
	if err != nil {
		return nil, err
	}

	response := dto.ToAyahTajweedResponse(*ayahTajweed)
	return &response, nil
}

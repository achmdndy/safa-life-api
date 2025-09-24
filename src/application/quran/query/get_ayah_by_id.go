package query

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// GetAyahByIDQuery represents the query to get an ayah by ID
type GetAyahByIDQuery struct {
	SurahID int
	AyahID  int
}

// GetAyahByID retrieves an ayah by its surah and ayah ID
func (h *QueryHandler) GetAyahByID(ctx context.Context, query GetAyahByIDQuery) (*dto.AyahResponse, error) {
	if query.SurahID < 1 || query.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", query.SurahID)
	}

	if query.AyahID < 1 {
		return nil, fmt.Errorf("invalid ayah ID: %d", query.AyahID)
	}

	ayah, err := h.quranService.GetAyahByID(ctx, query.SurahID, query.AyahID)
	if err != nil {
		return nil, err
	}

	response := dto.ToAyahResponse(*ayah)
	return &response, nil
}
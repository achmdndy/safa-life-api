package query

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// GetSurahWithAyahsQuery represents the query to get a surah with its ayahs
type GetSurahWithAyahsQuery struct {
	ID int
}

// GetSurahWithAyahs retrieves a surah with all its ayahs
func (h *QueryHandler) GetSurahWithAyahs(ctx context.Context, query GetSurahWithAyahsQuery) (*dto.SurahWithAyahsResponse, error) {
	if query.ID < 1 || query.ID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", query.ID)
	}

	surahWithAyahs, err := h.quranService.GetSurahWithAyahs(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	response := dto.ToSurahWithAyahsResponse(*surahWithAyahs)
	return &response, nil
}
package query

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// GetAyahsByJuzQuery represents the query to get ayahs by juz
type GetAyahsByJuzQuery struct {
	JuzID int
}

// GetAyahsByJuz retrieves ayahs by juz ID
func (h *QueryHandler) GetAyahsByJuz(ctx context.Context, query GetAyahsByJuzQuery) (*dto.JuzWithContentResponse, error) {
	if query.JuzID < 1 || query.JuzID > 30 {
		return nil, fmt.Errorf("invalid juz ID: %d", query.JuzID)
	}

	juzWithContent, err := h.quranService.GetJuzWithContent(ctx, query.JuzID)
	if err != nil {
		return nil, err
	}

	response := dto.ToJuzWithContentResponse(*juzWithContent)
	return &response, nil
}
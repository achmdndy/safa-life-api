package query

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// GetJuzWithContentQuery represents the query to get a juz with its content
type GetJuzWithContentQuery struct {
	ID int
}

// GetJuzWithContent retrieves a juz with all its content (ayahs)
func (h *QueryHandler) GetJuzWithContent(ctx context.Context, query GetJuzWithContentQuery) (*dto.JuzWithContentResponse, error) {
	if query.ID < 1 || query.ID > 30 {
		return nil, fmt.Errorf("invalid juz ID: %d", query.ID)
	}

	juzWithContent, err := h.quranService.GetJuzWithContent(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	response := dto.ToJuzWithContentResponse(*juzWithContent)
	return &response, nil
}
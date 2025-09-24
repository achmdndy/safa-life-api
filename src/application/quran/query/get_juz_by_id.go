package query

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// GetJuzByIDQuery represents the query to get a juz by ID
type GetJuzByIDQuery struct {
	ID int
}

// GetJuzByID retrieves a juz by its ID
func (h *QueryHandler) GetJuzByID(ctx context.Context, query GetJuzByIDQuery) (*dto.JuzResponse, error) {
	if query.ID < 1 || query.ID > 30 {
		return nil, fmt.Errorf("invalid juz ID: %d", query.ID)
	}

	juz, err := h.quranService.GetJuzByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	response := dto.ToJuzResponse(*juz)
	return &response, nil
}
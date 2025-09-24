package query

import (
	"context"
	"github.com/achmdndy/safa-life-api/src/application/quran/dto"
)

// GetAllJuzQuery represents the query to get all juz
type GetAllJuzQuery struct{}

// GetAllJuz retrieves all juz
func (h *QueryHandler) GetAllJuz(ctx context.Context, query GetAllJuzQuery) ([]dto.JuzResponse, error) {
	juzList, err := h.quranService.GetAllJuz(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToJuzResponseSlice(juzList), nil
}
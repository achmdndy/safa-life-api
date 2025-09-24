package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
)

// GetAyahTranslationQuery defines the query for getting an ayah translation.
type GetAyahTranslationQuery struct {
	TranslationID string
	SurahID       int
	AyahID        int
}

// GetAyahTranslation handles the retrieval of a specific ayah translation.
func (h *QueryHandler) GetAyahTranslation(ctx context.Context, query GetAyahTranslationQuery) (*dto.AyahTranslationResponse, error) {
	domain, err := h.translationService.GetAyahTranslation(ctx, query.TranslationID, query.SurahID, query.AyahID)
	if err != nil {
		return nil, err
	}
	resp := dto.ToAyahTranslationResponse(*domain)
	return &resp, nil
}

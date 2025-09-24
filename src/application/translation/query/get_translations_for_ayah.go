package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
)

// GetTranslationsForAyahQuery defines the query for getting all translations for an ayah.
type GetTranslationsForAyahQuery struct {
	SurahID int
	AyahID  int
}

// GetTranslationsForAyah handles the retrieval of all translations for a specific ayah.
func (h *QueryHandler) GetTranslationsForAyah(ctx context.Context, query GetTranslationsForAyahQuery) ([]dto.AyahTranslationResponse, error) {
	domains, err := h.translationService.GetTranslationsForAyah(ctx, query.SurahID, query.AyahID)
	if err != nil {
		return nil, err
	}
	return dto.ToAyahTranslationResponseSlice(domains), nil
}

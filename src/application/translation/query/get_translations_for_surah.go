package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
)

// GetTranslationsForSurahQuery defines the query for getting all translations for a surah.
type GetTranslationsForSurahQuery struct {
	TranslationID string
	SurahID       int
}

// GetTranslationsForSurah handles the retrieval of all ayah translations for a specific surah.
func (h *QueryHandler) GetTranslationsForSurah(ctx context.Context, query GetTranslationsForSurahQuery) ([]dto.AyahTranslationResponse, error) {
	domains, err := h.translationService.GetTranslationsForSurah(ctx, query.TranslationID, query.SurahID)
	if err != nil {
		return nil, err
	}
	return dto.ToAyahTranslationResponseSlice(domains), nil
}

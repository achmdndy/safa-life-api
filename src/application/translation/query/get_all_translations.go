package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
)

// GetAllTranslationsQuery defines the query for getting all translation editions.
type GetAllTranslationsQuery struct{}

// GetAllTranslations handles the retrieval of all translation editions.
func (h *QueryHandler) GetAllTranslations(ctx context.Context, query GetAllTranslationsQuery) ([]dto.TranslationResponse, error) {
	domains, err := h.translationService.GetAllTranslations(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ToTranslationResponseSlice(domains), nil
}

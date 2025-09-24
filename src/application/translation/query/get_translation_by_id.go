package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
)

// GetTranslationByIDQuery defines the query for getting a translation edition by ID.
type GetTranslationByIDQuery struct {
	ID string
}

// GetTranslationByID handles the retrieval of a translation edition by its ID.
func (h *QueryHandler) GetTranslationByID(ctx context.Context, query GetTranslationByIDQuery) (*dto.TranslationResponse, error) {
	domain, err := h.translationService.GetTranslationByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	resp := dto.ToTranslationResponse(*domain)
	return &resp, nil
}

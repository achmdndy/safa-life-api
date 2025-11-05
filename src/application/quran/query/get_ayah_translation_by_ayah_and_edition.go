package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAyahTranslationByAyahAndEditionQuery struct {
	AyahID    string
	EditionID string
}

func (h *QueryHandler) GetAyahTranslationByAyahAndEdition(ctx context.Context, query GetAyahTranslationByAyahAndEditionQuery) (*dto.AyahTranslationResponse, error) {
	ayahID, err := h.uuidGenerator.Parse(query.AyahID)
	if err != nil {
		return nil, err
	}

	editionID, err := h.uuidGenerator.Parse(query.EditionID)
	if err != nil {
		return nil, err
	}

	translation, err := h.ayahTranslationService.GetAyahTranslationByAyahAndEdition(ctx, ayahID, editionID)
	if err != nil {
		return nil, err
	}

	return dto.ToAyahTranslationResponse(translation), nil
}

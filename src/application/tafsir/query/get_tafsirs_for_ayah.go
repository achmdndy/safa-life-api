package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
)

type GetTafsirsForAyahQuery struct {
	SurahID int
	AyahID  int
}

func (h *QueryHandler) GetTafsirsForAyah(ctx context.Context, query GetTafsirsForAyahQuery) ([]dto.AyahTafsirResponse, error) {
	domains, err := h.tafsirService.GetTafsirsForAyah(ctx, query.SurahID, query.AyahID)
	if err != nil {
		return nil, err
	}
	return dto.ToAyahTafsirResponseSlice(domains), nil
}

package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
)

type GetTafsirsForSurahQuery struct {
	TafsirID string
	SurahID  int
}

func (h *QueryHandler) GetTafsirsForSurah(ctx context.Context, query GetTafsirsForSurahQuery) ([]dto.AyahTafsirResponse, error) {
	domains, err := h.tafsirService.GetTafsirsForSurah(ctx, query.TafsirID, query.SurahID)
	if err != nil {
		return nil, err
	}
	return dto.ToAyahTafsirResponseSlice(domains), nil
}

package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
)

type GetAyahTafsirQuery struct {
	TafsirID string
	SurahID  int
	AyahID   int
}

func (h *QueryHandler) GetAyahTafsir(ctx context.Context, query GetAyahTafsirQuery) (*dto.AyahTafsirResponse, error) {
	domain, err := h.tafsirService.GetAyahTafsir(ctx, query.TafsirID, query.SurahID, query.AyahID)
	if err != nil {
		return nil, err
	}
	resp := dto.ToAyahTafsirResponse(*domain)
	return &resp, nil
}

package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/audio/dto"
)

type GetAyahAudioQuery struct {
	ReciterID string
	SurahID   int
	AyahID    int
}

func (h *QueryHandler) GetAyahAudio(ctx context.Context, query GetAyahAudioQuery) (*dto.AyahAudioResponse, error) {
	domain, err := h.audioService.GetAyahAudio(ctx, query.ReciterID, query.SurahID, query.AyahID)
	if err != nil {
		return nil, err
	}
	resp := dto.ToAyahAudioResponse(*domain)
	return &resp, nil
}

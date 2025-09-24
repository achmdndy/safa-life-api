package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/audio/dto"
)

type GetAudioForSurahQuery struct {
	ReciterID string
	SurahID   int
}

func (h *QueryHandler) GetAudioForSurah(ctx context.Context, query GetAudioForSurahQuery) ([]dto.AyahAudioResponse, error) {
	domains, err := h.audioService.GetAudioFilesForSurah(ctx, query.ReciterID, query.SurahID)
	if err != nil {
		return nil, err
	}
	return dto.ToAyahAudioResponseSlice(domains), nil
}

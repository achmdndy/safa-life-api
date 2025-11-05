package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAyahAudioFileByAyahAndReciterQuery struct {
	AyahID    string
	ReciterID string
}

func (h *QueryHandler) GetAyahAudioFileByAyahAndReciter(ctx context.Context, q GetAyahAudioFileByAyahAndReciterQuery) (*dto.AyahAudioFileResponse, error) {
	ayahID, err := h.uuidGenerator.Parse(q.AyahID)
	if err != nil {
		return nil, err
	}
	reciterID, err := h.uuidGenerator.Parse(q.ReciterID)
	if err != nil {
		return nil, err
	}
	audio, err := h.audioService.GetAyahAudioFileByAyahAndReciter(ctx, ayahID, reciterID)
	if err != nil {
		return nil, err
	}
	return dto.ToAyahAudioFileResponse(audio), nil
}

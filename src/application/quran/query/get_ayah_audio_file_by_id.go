package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAyahAudioFileByIdQuery struct {
	ID string
}

func (h *QueryHandler) GetAyahAudioFileById(ctx context.Context, q GetAyahAudioFileByIdQuery) (*dto.AyahAudioFileResponse, error) {
	id, err := h.uuidGenerator.Parse(q.ID)
	if err != nil {
		return nil, err
	}
	audio, err := h.audioService.GetAyahAudioFileById(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToAyahAudioFileResponse(audio), nil
}

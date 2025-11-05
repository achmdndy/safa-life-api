package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetAyahAudioFilesBySurahAndReciterQuery struct {
	SurahID   string
	ReciterID string
	Limit     int
	Offset    int
}

func (h *QueryHandler) GetAyahAudioFilesBySurahAndReciter(ctx context.Context, q GetAyahAudioFilesBySurahAndReciterQuery) (*dto.AyahAudioFileListResponse, error) {
	surahID, err := h.uuidGenerator.Parse(q.SurahID)
	if err != nil {
		return nil, err
	}
	reciterID, err := h.uuidGenerator.Parse(q.ReciterID)
	if err != nil {
		return nil, err
	}
	items, err := h.audioService.GetAyahAudioFilesBySurahAndReciter(ctx, surahID, reciterID, q.Limit, q.Offset)
	if err != nil {
		return nil, err
	}
	count, err := h.audioService.CountAyahAudioFilesBySurahAndReciter(ctx, surahID, reciterID)
	if err != nil {
		return nil, err
	}
	return &dto.AyahAudioFileListResponse{
		Data: dto.ToAyahAudioFileResponseSlice(items),
		Pagination: &dto.PaginationResponse{
			Total:  count,
			Limit:  q.Limit,
			Offset: q.Offset,
		},
	}, nil
}

package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetSurahByNumberQuery struct {
	Number int
}

func (h *QueryHandler) GetSurahByNumber(ctx context.Context, query GetSurahByNumberQuery) (*dto.SurahResponse, error) {
	surah, err := h.surahService.GetSurahByNumber(ctx, query.Number)
	if err != nil {
		return nil, err
	}

	return dto.ToSurahResponse(surah), nil
}

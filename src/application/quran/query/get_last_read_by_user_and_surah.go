package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type GetLastReadByUserAndSurahQuery struct {
	UserID  string
	SurahID string
	Include string
}

func (h *QueryHandler) GetLastReadByUserAndSurah(ctx context.Context, query GetLastReadByUserAndSurahQuery) (interface{}, error) {
	surahID, err := h.uuidGenerator.Parse(query.SurahID)
	if err != nil {
		return nil, err
	}

	var lr *quran.LastRead
	lr, err = h.lastReadService.GetLastReadByUserAndSurah(ctx, query.UserID, surahID)
	if err != nil {
		return nil, err
	}

	if query.Include == "relations" && lr != nil {
		var lrWith *quran.LastReadWithRelations
		lrWith, err = h.lastReadService.GetLastReadByIdWithRelations(ctx, lr.ID)
		if err != nil {
			return nil, err
		}
		return dto.ToLastReadWithRelationsResponse(lrWith), nil
	}

	return dto.ToLastReadResponse(lr), nil
}

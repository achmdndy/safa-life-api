package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetReciterByIdQuery struct {
	ID string
}

func (h *QueryHandler) GetReciterById(ctx context.Context, q GetReciterByIdQuery) (*dto.ReciterResponse, error) {
	id, err := h.uuidGenerator.Parse(q.ID)
	if err != nil {
		return nil, err
	}
	rec, err := h.reciterService.GetReciterById(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToReciterResponse(rec), nil
}

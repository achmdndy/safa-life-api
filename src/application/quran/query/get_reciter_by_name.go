package query

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
)

type GetReciterByNameQuery struct {
	Name string
}

func (h *QueryHandler) GetReciterByName(ctx context.Context, q GetReciterByNameQuery) (*dto.ReciterResponse, error) {
	rec, err := h.reciterService.GetReciterByName(ctx, q.Name)
	if err != nil {
		return nil, err
	}
	return dto.ToReciterResponse(rec), nil
}

package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
)

type GetTafsirByIDQuery struct {
	ID string
}

func (h *QueryHandler) GetTafsirByID(ctx context.Context, query GetTafsirByIDQuery) (*dto.TafsirResponse, error) {
	domain, err := h.tafsirService.GetTafsirByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	resp := dto.ToTafsirResponse(*domain)
	return &resp, nil
}

package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
)

type GetAllTafsirsQuery struct{}

func (h *QueryHandler) GetAllTafsirs(ctx context.Context, query GetAllTafsirsQuery) ([]dto.TafsirResponse, error) {
	domains, err := h.tafsirService.GetAllTafsirs(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ToTafsirResponseSlice(domains), nil
}

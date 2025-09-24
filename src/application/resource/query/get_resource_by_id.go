package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/resource/dto"
)

type GetResourceByIDQuery struct {
	ID string
}

func (h *QueryHandler) GetResourceByID(ctx context.Context, query GetResourceByIDQuery) (*dto.ResourceResponse, error) {
	domain, err := h.resourceService.GetResourceByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	resp := dto.ToResourceResponse(*domain)
	return &resp, nil
}

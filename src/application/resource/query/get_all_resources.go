package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/resource/dto"
)

type GetAllResourcesQuery struct{}

func (h *QueryHandler) GetAllResources(ctx context.Context, query GetAllResourcesQuery) ([]dto.ResourceResponse, error) {
	domains, err := h.resourceService.GetAllResources(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ToResourceResponseSlice(domains), nil
}

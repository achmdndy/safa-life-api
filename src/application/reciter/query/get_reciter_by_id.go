package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/reciter/dto"
)

// GetReciterByIDQuery defines the query for getting a reciter by ID.
type GetReciterByIDQuery struct {
	ID string
}

// GetReciterByID handles the retrieval of a reciter by its ID.
func (h *QueryHandler) GetReciterByID(ctx context.Context, query GetReciterByIDQuery) (*dto.ReciterResponse, error) {
	domain, err := h.reciterService.GetReciterByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	resp := dto.ToReciterResponse(*domain)
	return &resp, nil
}

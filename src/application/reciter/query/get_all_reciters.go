package query

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/application/reciter/dto"
)

// GetAllRecitersQuery defines the query for getting all reciters.
type GetAllRecitersQuery struct{}

// GetAllReciters handles the retrieval of all reciters.
func (h *QueryHandler) GetAllReciters(ctx context.Context, query GetAllRecitersQuery) ([]dto.ReciterResponse, error) {
	domains, err := h.reciterService.GetAllReciters(ctx)
	if err != nil {
		return nil, err
	}
	return dto.ToReciterResponseSlice(domains), nil
}

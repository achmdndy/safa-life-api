package queries

import (
	"context"
	"errors"

	"github.com/achmdndy/safa-life-api/src/application/health/dto"
	"github.com/achmdndy/safa-life-api/src/domain/health"
)

type GetHealthQuery struct{}

type GetHealthQueryHandler struct {
	healthService health.HealthServiceInterface
}

func NewGetHealthQueryHandler(healthService health.HealthServiceInterface) *GetHealthQueryHandler {
	return &GetHealthQueryHandler{
		healthService: healthService,
	}
}

func (h *GetHealthQueryHandler) HandleHealth(ctx context.Context, query GetHealthQuery) (dto.HealthResponse, error) {
	domainStatus := h.healthService.GetHealth()
	
	// Convert domain entity to application DTO
	response := dto.HealthResponse{
		Status:   domainStatus.Status,
		Services: domainStatus.Services,
	}
	
	return response, nil
}

// Handle implements the domain QueryHandlerInterface
func (h *GetHealthQueryHandler) Handle(query interface{}) (interface{}, error) {
	ctx := context.Background()
	if healthQuery, ok := query.(GetHealthQuery); ok {
		return h.HandleHealth(ctx, healthQuery)
	}
	return nil, errors.New("invalid query type")
}
package query

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
)

// GetTajweedRuleByIDQuery represents the query to get a tajweed rule by ID.
type GetTajweedRuleByIDQuery struct {
	ID string
}

// GetTajweedRuleByID retrieves a tajweed rule by its ID.
func (h *QueryHandler) GetTajweedRuleByID(ctx context.Context, query GetTajweedRuleByIDQuery) (*dto.TajweedRuleResponse, error) {
	if query.ID == "" {
		return nil, fmt.Errorf("rule ID cannot be empty")
	}

	rule, err := h.tajweedService.GetTajweedRuleByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	response := dto.ToTajweedRuleResponse(*rule)
	return &response, nil
}

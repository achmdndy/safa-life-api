package query

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
)

// GetTajweedRuleByNameQuery represents the query to get a tajweed rule by name.
type GetTajweedRuleByNameQuery struct {
	RuleName string
}

// GetTajweedRuleByName retrieves a tajweed rule by its name.
func (h *QueryHandler) GetTajweedRuleByName(ctx context.Context, query GetTajweedRuleByNameQuery) (*dto.TajweedRuleResponse, error) {
	if query.RuleName == "" {
		return nil, fmt.Errorf("rule name cannot be empty")
	}

	rule, err := h.tajweedService.GetTajweedRuleByName(ctx, query.RuleName)
	if err != nil {
		return nil, err
	}

	response := dto.ToTajweedRuleResponse(*rule)
	return &response, nil
}

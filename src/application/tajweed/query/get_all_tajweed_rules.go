package query

import (
	"context"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/dto"
)

// GetAllTajweedRulesQuery represents the query to get all tajweed rules.
type GetAllTajweedRulesQuery struct{}

// GetAllTajweedRules retrieves all tajweed rules.
func (h *QueryHandler) GetAllTajweedRules(ctx context.Context, query GetAllTajweedRulesQuery) ([]dto.TajweedRuleResponse, error) {
	rules, err := h.tajweedService.GetAllTajweedRules(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToTajweedRuleResponseSlice(rules), nil
}

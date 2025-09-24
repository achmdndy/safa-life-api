package command

import (
	"context"
	"fmt"

	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
)

// UpdateTajweedRuleCommand represents the command to update a tajweed rule.
type UpdateTajweedRuleCommand struct {
	ID          string
	Rule        *string
	Explanation *string
	Color       *string
}

// UpdateTajweedRule updates an existing tajweed rule.
func (h *CommandHandler) UpdateTajweedRule(ctx context.Context, cmd UpdateTajweedRuleCommand) (*tajweed.TajweedRule, error) {
	if cmd.ID == "" {
		return nil, fmt.Errorf("rule ID cannot be empty")
	}

	existing, err := h.tajweedService.GetTajweedRuleByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if cmd.Rule != nil {
		existing.Rule = *cmd.Rule
	}
	if cmd.Explanation != nil {
		existing.Explanation = *cmd.Explanation
	}
	if cmd.Color != nil {
		existing.Color = *cmd.Color
	}

	var result *tajweed.TajweedRule
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.tajweedService.UpdateTajweedRule(ctx, existing)
		return err
	})

	return result, err
}

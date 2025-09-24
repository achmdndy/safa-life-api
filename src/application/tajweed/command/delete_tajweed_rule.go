package command

import (
	"context"
	"fmt"
)

// DeleteTajweedRuleCommand represents the command to delete a tajweed rule.
type DeleteTajweedRuleCommand struct {
	ID string
}

// DeleteTajweedRule deletes a tajweed rule.
func (h *CommandHandler) DeleteTajweedRule(ctx context.Context, cmd DeleteTajweedRuleCommand) error {
	if cmd.ID == "" {
		return fmt.Errorf("rule ID cannot be empty")
	}

	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.tajweedService.DeleteTajweedRule(ctx, cmd.ID)
	})
}

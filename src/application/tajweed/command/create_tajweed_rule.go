package command

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
)

// CreateTajweedRuleCommand represents the command to create a new tajweed rule.
type CreateTajweedRuleCommand struct {
	ID          string
	Rule        string
	Explanation string
	Color       string
}

// CreateTajweedRule creates a new tajweed rule.
func (h *CommandHandler) CreateTajweedRule(ctx context.Context, cmd CreateTajweedRuleCommand) (*tajweed.TajweedRule, error) {
	if cmd.ID == "" {
		return nil, fmt.Errorf("rule ID cannot be empty")
	}
	if cmd.Rule == "" {
		return nil, fmt.Errorf("rule name cannot be empty")
	}

	rule := &tajweed.TajweedRule{
		ID:          cmd.ID,
		Rule:        cmd.Rule,
		Explanation: cmd.Explanation,
		Color:       cmd.Color,
	}

	var result *tajweed.TajweedRule
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.tajweedService.CreateTajweedRule(ctx, rule)
		return err
	})

	return result, err
}

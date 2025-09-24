package command

import (
	"context"
	"fmt"
)

// DeleteJuzCommand represents the command to delete a juz
type DeleteJuzCommand struct {
	ID int
}

// DeleteJuz deletes a juz
func (h *CommandHandler) DeleteJuz(ctx context.Context, cmd DeleteJuzCommand) error {
	// Validate ID
	if cmd.ID < 1 || cmd.ID > 30 {
		return fmt.Errorf("invalid juz ID: %d", cmd.ID)
	}

	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.quranService.DeleteJuz(ctx, cmd.ID)
	})
}
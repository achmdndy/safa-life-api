package command

import (
	"context"
)

// DeleteReciterCommand defines the command for deleting a reciter.
type DeleteReciterCommand struct {
	ID string
}

// DeleteReciter handles the deletion of a reciter.
func (h *CommandHandler) DeleteReciter(ctx context.Context, cmd DeleteReciterCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.reciterService.DeleteReciter(ctx, cmd.ID)
	})
}

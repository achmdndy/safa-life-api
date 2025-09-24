package command

import (
	"context"
)

type DeleteResourceCommand struct {
	ID string
}

func (h *CommandHandler) DeleteResource(ctx context.Context, cmd DeleteResourceCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.resourceService.DeleteResource(ctx, cmd.ID)
	})
}

package command

import (
	"context"
)

type DeleteLastReadCommand struct {
	ID string
}

func (h *CommandHandler) DeleteLastRead(ctx context.Context, command DeleteLastReadCommand) error {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return err
	}

	return h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		return h.lastReadService.DeleteLastRead(ctx, id)
	})
}

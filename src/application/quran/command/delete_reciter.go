package command

import (
	"context"
)

type DeleteReciterCommand struct {
	ID string
}

func (h *CommandHandler) DeleteReciter(ctx context.Context, command DeleteReciterCommand) error {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return err
	}

	return h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		return h.reciterService.DeleteReciter(ctx, id)
	})
}

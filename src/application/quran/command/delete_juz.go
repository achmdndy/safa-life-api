package command

import (
	"context"
)

type DeleteJuzCommand struct {
	ID string
}

func (h *CommandHandler) DeleteJuz(ctx context.Context, command DeleteJuzCommand) error {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return err
	}

	return h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		return h.juzService.DeleteJuz(ctx, id)
	})
}
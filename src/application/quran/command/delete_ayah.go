package command

import (
	"context"
)

type DeleteAyahCommand struct {
	ID string
}

func (h *CommandHandler) DeleteAyah(ctx context.Context, command DeleteAyahCommand) error {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return err
	}

	return h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		return h.ayahService.DeleteAyah(ctx, id)
	})
}
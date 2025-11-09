package command

import (
	"context"
)

type DeleteProgressHatamCommand struct {
	ID string
}

func (h *CommandHandler) DeleteProgressHatam(ctx context.Context, command DeleteProgressHatamCommand) error {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return err
	}

	return h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		return h.progressService.DeleteProgressHatam(ctx, id)
	})
}

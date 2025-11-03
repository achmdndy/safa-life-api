package command

import (
	"context"
)

type DeleteSurahCommand struct {
	ID string
}

func (h *CommandHandler) DeleteSurah(ctx context.Context, command DeleteSurahCommand) error {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return err
	}

	return h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		return h.surahService.DeleteSurah(ctx, id)
	})
}
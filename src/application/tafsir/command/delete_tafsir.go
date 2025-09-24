package command

import (
	"context"
)

type DeleteTafsirCommand struct {
	ID string
}

func (h *CommandHandler) DeleteTafsir(ctx context.Context, cmd DeleteTafsirCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.tafsirService.DeleteTafsir(ctx, cmd.ID)
	})
}

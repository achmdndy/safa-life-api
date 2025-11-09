package command

import (
	"context"
)

type DeleteBookmarkAyahCommand struct {
	ID string
}

func (h *CommandHandler) DeleteBookmarkAyah(ctx context.Context, command DeleteBookmarkAyahCommand) error {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return err
	}

	return h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		return h.bookmarkService.DeleteBookmarkAyah(ctx, id)
	})
}

package command

import (
	"context"
)

type DeleteAyahAudioFileCommand struct {
	ID string
}

func (h *CommandHandler) DeleteAyahAudioFile(ctx context.Context, command DeleteAyahAudioFileCommand) error {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return err
	}

	return h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		return h.audioService.DeleteAyahAudioFile(ctx, id)
	})
}

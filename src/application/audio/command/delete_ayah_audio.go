package command

import (
	"context"
)

type DeleteAyahAudioCommand struct {
	ReciterID string
	SurahID   int
	AyahID    int
}

func (h *CommandHandler) DeleteAyahAudio(ctx context.Context, cmd DeleteAyahAudioCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.audioService.DeleteAyahAudio(ctx, cmd.ReciterID, cmd.SurahID, cmd.AyahID)
	})
}

package command

import (
	"context"
)

type DeleteAyahTafsirCommand struct {
	TafsirID string
	SurahID  int
	AyahID   int
}

func (h *CommandHandler) DeleteAyahTafsir(ctx context.Context, cmd DeleteAyahTafsirCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.tafsirService.DeleteAyahTafsir(ctx, cmd.TafsirID, cmd.SurahID, cmd.AyahID)
	})
}

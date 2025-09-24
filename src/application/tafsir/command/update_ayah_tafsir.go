package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/tafsir"
)

type UpdateAyahTafsirCommand struct {
	TafsirID string
	SurahID  int
	AyahID   int
	Text     *string
}

func (h *CommandHandler) UpdateAyahTafsir(ctx context.Context, cmd UpdateAyahTafsirCommand) (*tafsir.AyahTafsir, error) {
	existing, err := h.tafsirService.GetAyahTafsir(ctx, cmd.TafsirID, cmd.SurahID, cmd.AyahID)
	if err != nil {
		return nil, err
	}

	if cmd.Text != nil {
		existing.Text = *cmd.Text
	}

	var result *tafsir.AyahTafsir
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.tafsirService.UpdateAyahTafsir(ctx, existing)
		return err
	})

	return result, err
}

package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/tafsir"
)

type CreateAyahTafsirCommand struct {
	TafsirID string
	SurahID  int
	AyahID   int
	Text     string
}

func (h *CommandHandler) CreateAyahTafsir(ctx context.Context, cmd CreateAyahTafsirCommand) (*tafsir.AyahTafsir, error) {
	domain := &tafsir.AyahTafsir{
		TafsirID: cmd.TafsirID,
		SurahID:  cmd.SurahID,
		AyahID:   cmd.AyahID,
		Text:     cmd.Text,
	}

	var result *tafsir.AyahTafsir
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.tafsirService.CreateAyahTafsir(ctx, domain)
		return err
	})

	return result, err
}

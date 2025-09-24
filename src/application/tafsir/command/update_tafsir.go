package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/tafsir"
)

type UpdateTafsirCommand struct {
	ID       string
	Name     *string
	Author   *string
	Language *string
}

func (h *CommandHandler) UpdateTafsir(ctx context.Context, cmd UpdateTafsirCommand) (*tafsir.Tafsir, error) {
	existing, err := h.tafsirService.GetTafsirByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if cmd.Name != nil {
		existing.Name = *cmd.Name
	}
	if cmd.Author != nil {
		existing.Author = *cmd.Author
	}
	if cmd.Language != nil {
		existing.Language = *cmd.Language
	}

	var result *tafsir.Tafsir
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.tafsirService.UpdateTafsir(ctx, existing)
		return err
	})

	return result, err
}

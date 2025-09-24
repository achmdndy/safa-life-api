package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/tafsir"
)

type CreateTafsirCommand struct {
	ID       string
	Name     string
	Author   string
	Language string
}

func (h *CommandHandler) CreateTafsir(ctx context.Context, cmd CreateTafsirCommand) (*tafsir.Tafsir, error) {
	domain := &tafsir.Tafsir{
		ID:       cmd.ID,
		Name:     cmd.Name,
		Author:   cmd.Author,
		Language: cmd.Language,
	}

	var result *tafsir.Tafsir
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.tafsirService.CreateTafsir(ctx, domain)
		return err
	})

	return result, err
}

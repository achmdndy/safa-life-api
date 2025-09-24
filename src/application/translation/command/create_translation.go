package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/translation"
)

// CreateTranslationCommand defines the command for creating a new translation edition.
type CreateTranslationCommand struct {
	ID       string
	Name     string
	Author   string
	Language string
}

// CreateTranslation handles the creation of a translation edition.
func (h *CommandHandler) CreateTranslation(ctx context.Context, cmd CreateTranslationCommand) (*translation.Translation, error) {
	domain := &translation.Translation{
		ID:       cmd.ID,
		Name:     cmd.Name,
		Author:   cmd.Author,
		Language: cmd.Language,
	}

	var result *translation.Translation
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.translationService.CreateTranslation(ctx, domain)
		return err
	})

	return result, err
}

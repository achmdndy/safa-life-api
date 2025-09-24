package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/translation"
)

// UpdateTranslationCommand defines the command for updating a translation edition.
type UpdateTranslationCommand struct {
	ID       string
	Name     *string
	Author   *string
	Language *string
}

// UpdateTranslation handles the update of a translation edition.
func (h *CommandHandler) UpdateTranslation(ctx context.Context, cmd UpdateTranslationCommand) (*translation.Translation, error) {
	existing, err := h.translationService.GetTranslationByID(ctx, cmd.ID)
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

	var result *translation.Translation
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.translationService.UpdateTranslation(ctx, existing)
		return err
	})

	return result, err
}

package command

import (
	"context"
)

// DeleteTranslationCommand defines the command for deleting a translation edition.
type DeleteTranslationCommand struct {
	ID string
}

// DeleteTranslation handles the deletion of a translation edition.
func (h *CommandHandler) DeleteTranslation(ctx context.Context, cmd DeleteTranslationCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.translationService.DeleteTranslation(ctx, cmd.ID)
	})
}

package command

import (
	"context"
)

// DeleteAyahTranslationCommand defines the command for deleting an ayah translation.
type DeleteAyahTranslationCommand struct {
	TranslationID string
	SurahID       int
	AyahID        int
}

// DeleteAyahTranslation handles the deletion of an ayah translation.
func (h *CommandHandler) DeleteAyahTranslation(ctx context.Context, cmd DeleteAyahTranslationCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.translationService.DeleteAyahTranslation(ctx, cmd.TranslationID, cmd.SurahID, cmd.AyahID)
	})
}

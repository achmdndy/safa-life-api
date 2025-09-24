package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/translation"
)

// UpdateAyahTranslationCommand defines the command for updating an ayah translation.
type UpdateAyahTranslationCommand struct {
	TranslationID string
	SurahID       int
	AyahID        int
	Text          *string
}

// UpdateAyahTranslation handles the update of an ayah translation.
func (h *CommandHandler) UpdateAyahTranslation(ctx context.Context, cmd UpdateAyahTranslationCommand) (*translation.AyahTranslation, error) {
	existing, err := h.translationService.GetAyahTranslation(ctx, cmd.TranslationID, cmd.SurahID, cmd.AyahID)
	if err != nil {
		return nil, err
	}

	if cmd.Text != nil {
		existing.Text = *cmd.Text
	}

	var result *translation.AyahTranslation
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.translationService.UpdateAyahTranslation(ctx, existing)
		return err
	})

	return result, err
}

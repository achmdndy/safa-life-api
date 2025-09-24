package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/translation"
)

// CreateAyahTranslationCommand defines the command for creating a new ayah translation.
type CreateAyahTranslationCommand struct {
	TranslationID string
	SurahID       int
	AyahID        int
	Text          string
}

// CreateAyahTranslation handles the creation of an ayah translation.
func (h *CommandHandler) CreateAyahTranslation(ctx context.Context, cmd CreateAyahTranslationCommand) (*translation.AyahTranslation, error) {
	domain := &translation.AyahTranslation{
		TranslationID: cmd.TranslationID,
		SurahID:       cmd.SurahID,
		AyahID:        cmd.AyahID,
		Text:          cmd.Text,
	}

	var result *translation.AyahTranslation
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.translationService.CreateAyahTranslation(ctx, domain)
		return err
	})

	return result, err
}

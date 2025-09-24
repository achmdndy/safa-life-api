package command

import (
	"context"
	"fmt"
)

// DeleteAyahCommand represents the command to delete an ayah
type DeleteAyahCommand struct {
	SurahID int
	AyahID  int
}

// DeleteAyah deletes an ayah
func (h *CommandHandler) DeleteAyah(ctx context.Context, cmd DeleteAyahCommand) error {
	// Validate surah ID
	if cmd.SurahID < 1 || cmd.SurahID > 114 {
		return fmt.Errorf("invalid surah ID: %d", cmd.SurahID)
	}

	// Validate ayah ID
	if cmd.AyahID < 1 {
		return fmt.Errorf("invalid ayah ID: %d", cmd.AyahID)
	}

	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.quranService.DeleteAyah(ctx, cmd.SurahID, cmd.AyahID)
	})
}
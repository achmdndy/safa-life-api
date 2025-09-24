package command

import (
	"context"
	"fmt"
)

// DeleteAyahTajweedCommand represents the command to delete an ayah tajweed.
type DeleteAyahTajweedCommand struct {
	TajweedID string
	SurahID   int
	AyahID    int
}

// DeleteAyahTajweed deletes an ayah tajweed.
func (h *CommandHandler) DeleteAyahTajweed(ctx context.Context, cmd DeleteAyahTajweedCommand) error {
	if cmd.TajweedID == "" {
		return fmt.Errorf("tajweed ID cannot be empty")
	}
	if cmd.SurahID < 1 || cmd.SurahID > 114 {
		return fmt.Errorf("invalid surah ID: %d", cmd.SurahID)
	}
	if cmd.AyahID < 1 {
		return fmt.Errorf("invalid ayah ID: %d", cmd.AyahID)
	}

	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.tajweedService.DeleteAyahTajweed(ctx, cmd.TajweedID, cmd.SurahID, cmd.AyahID)
	})
}

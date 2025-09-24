package command

import (
	"context"
	"fmt"
)

// DeleteSurahCommand represents the command to delete a surah
type DeleteSurahCommand struct {
	ID int
}

// DeleteSurah deletes a surah
func (h *CommandHandler) DeleteSurah(ctx context.Context, cmd DeleteSurahCommand) error {
	// Validate ID
	if cmd.ID < 1 || cmd.ID > 114 {
		return fmt.Errorf("invalid surah ID: %d", cmd.ID)
	}

	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.quranService.DeleteSurah(ctx, cmd.ID)
	})
}
package command

import (
	"context"
	"fmt"

	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
)

// UpdateAyahTajweedCommand represents the command to update an ayah tajweed.
type UpdateAyahTajweedCommand struct {
	TajweedID string
	SurahID   int
	AyahID    int
	Words     []tajweed.TajweedWord
}

// UpdateAyahTajweed updates an existing ayah tajweed.
func (h *CommandHandler) UpdateAyahTajweed(ctx context.Context, cmd UpdateAyahTajweedCommand) (*tajweed.AyahTajweed, error) {
	if cmd.TajweedID == "" {
		return nil, fmt.Errorf("tajweed ID cannot be empty")
	}
	if cmd.SurahID < 1 || cmd.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", cmd.SurahID)
	}
	if cmd.AyahID < 1 {
		return nil, fmt.Errorf("invalid ayah ID: %d", cmd.AyahID)
	}

	existing, err := h.tajweedService.GetAyahTajweed(ctx, cmd.TajweedID, cmd.SurahID, cmd.AyahID)
	if err != nil {
		return nil, err
	}

	if len(cmd.Words) > 0 {
		existing.Words = cmd.Words
	}

	var result *tajweed.AyahTajweed
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.tajweedService.UpdateAyahTajweed(ctx, existing)
		return err
	})

	return result, err
}

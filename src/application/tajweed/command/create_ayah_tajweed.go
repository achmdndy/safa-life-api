package command

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
)

// CreateAyahTajweedCommand represents the command to create a new ayah tajweed.
type CreateAyahTajweedCommand struct {
	TajweedID string
	SurahID   int
	AyahID    int
	Words     []tajweed.TajweedWord
}

// CreateAyahTajweed creates a new ayah tajweed.
func (h *CommandHandler) CreateAyahTajweed(ctx context.Context, cmd CreateAyahTajweedCommand) (*tajweed.AyahTajweed, error) {
	if cmd.TajweedID == "" {
		return nil, fmt.Errorf("tajweed ID cannot be empty")
	}
	if cmd.SurahID < 1 || cmd.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", cmd.SurahID)
	}
	if cmd.AyahID < 1 {
		return nil, fmt.Errorf("invalid ayah ID: %d", cmd.AyahID)
	}
	if len(cmd.Words) == 0 {
		return nil, fmt.Errorf("words cannot be empty")
	}

	ayahTajweed := &tajweed.AyahTajweed{
		TajweedID: cmd.TajweedID,
		SurahID:   cmd.SurahID,
		AyahID:    cmd.AyahID,
		Words:     cmd.Words,
	}

	var result *tajweed.AyahTajweed
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.tajweedService.CreateAyahTajweed(ctx, ayahTajweed)
		return err
	})

	return result, err
}

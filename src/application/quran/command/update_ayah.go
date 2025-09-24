package command

import (
	"context"
	"fmt"

	"github.com/achmdndy/safa-life-api/src/domain/quran"
)

// UpdateAyahCommand represents the command to update an ayah
type UpdateAyahCommand struct {
	SurahID      int
	AyahID       int
	Text         *string
	PageNumber   *int
	JuzNumber    *int
	HizbNumber   *int
	ManzilNumber *int
}

// UpdateAyah updates an existing ayah
func (h *CommandHandler) UpdateAyah(ctx context.Context, cmd UpdateAyahCommand) (*quran.Ayah, error) {
	// Validate surah ID
	if cmd.SurahID < 1 || cmd.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", cmd.SurahID)
	}

	// Validate ayah ID
	if cmd.AyahID < 1 {
		return nil, fmt.Errorf("invalid ayah ID: %d", cmd.AyahID)
	}

	// Validate page number if provided
	if cmd.PageNumber != nil && (*cmd.PageNumber < 1 || *cmd.PageNumber > 604) {
		return nil, fmt.Errorf("invalid page number: %d", *cmd.PageNumber)
	}

	// Validate juz number if provided
	if cmd.JuzNumber != nil && (*cmd.JuzNumber < 1 || *cmd.JuzNumber > 30) {
		return nil, fmt.Errorf("invalid juz number: %d", *cmd.JuzNumber)
	}

	// Validate hizb number if provided
	if cmd.HizbNumber != nil && (*cmd.HizbNumber < 1 || *cmd.HizbNumber > 60) {
		return nil, fmt.Errorf("invalid hizb number: %d", *cmd.HizbNumber)
	}

	// Validate manzil number if provided
	if cmd.ManzilNumber != nil && (*cmd.ManzilNumber < 1 || *cmd.ManzilNumber > 7) {
		return nil, fmt.Errorf("invalid manzil number: %d", *cmd.ManzilNumber)
	}

	// Get existing ayah
	existingAyah, err := h.quranService.GetAyahByID(ctx, cmd.SurahID, cmd.AyahID)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if cmd.Text != nil {
		existingAyah.Text = *cmd.Text
	}
	if cmd.PageNumber != nil {
		existingAyah.PageNumber = *cmd.PageNumber
	}
	if cmd.JuzNumber != nil {
		existingAyah.JuzNumber = *cmd.JuzNumber
	}
	if cmd.HizbNumber != nil {
		existingAyah.HizbNumber = *cmd.HizbNumber
	}
	if cmd.ManzilNumber != nil {
		existingAyah.ManzilNumber = *cmd.ManzilNumber
	}

	var result *quran.Ayah
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.quranService.UpdateAyah(ctx, existingAyah)
		return err
	})

	return result, err
}

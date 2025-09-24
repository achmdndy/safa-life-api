package command

import (
	"context"
	"fmt"

	"github.com/achmdndy/safa-life-api/src/domain/quran"
)

// UpdateJuzCommand represents the command to update a juz
type UpdateJuzCommand struct {
	ID         int
	StartSurah *int
	StartAyah  *int
	EndSurah   *int
	EndAyah    *int
}

// UpdateJuz updates an existing juz
func (h *CommandHandler) UpdateJuz(ctx context.Context, cmd UpdateJuzCommand) (*quran.Juz, error) {
	// Validate ID
	if cmd.ID < 1 || cmd.ID > 30 {
		return nil, fmt.Errorf("invalid juz ID: %d", cmd.ID)
	}

	// Validate start surah if provided
	if cmd.StartSurah != nil && (*cmd.StartSurah < 1 || *cmd.StartSurah > 114) {
		return nil, fmt.Errorf("invalid start surah: %d", *cmd.StartSurah)
	}

	// Validate end surah if provided
	if cmd.EndSurah != nil && (*cmd.EndSurah < 1 || *cmd.EndSurah > 114) {
		return nil, fmt.Errorf("invalid end surah: %d", *cmd.EndSurah)
	}

	// Validate start ayah if provided
	if cmd.StartAyah != nil && *cmd.StartAyah < 1 {
		return nil, fmt.Errorf("invalid start ayah: %d", *cmd.StartAyah)
	}

	// Validate end ayah if provided
	if cmd.EndAyah != nil && *cmd.EndAyah < 1 {
		return nil, fmt.Errorf("invalid end ayah: %d", *cmd.EndAyah)
	}

	// Get existing juz
	existingJuz, err := h.quranService.GetJuzByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if cmd.StartSurah != nil {
		existingJuz.StartSurah = *cmd.StartSurah
	}
	if cmd.StartAyah != nil {
		existingJuz.StartAyah = *cmd.StartAyah
	}
	if cmd.EndSurah != nil {
		existingJuz.EndSurah = *cmd.EndSurah
	}
	if cmd.EndAyah != nil {
		existingJuz.EndAyah = *cmd.EndAyah
	}

	var result *quran.Juz
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.quranService.UpdateJuz(ctx, existingJuz)
		return err
	})

	return result, err
}

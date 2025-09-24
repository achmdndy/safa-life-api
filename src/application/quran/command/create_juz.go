package command

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/domain/quran"
)

// CreateJuzCommand represents the command to create a new juz
type CreateJuzCommand struct {
	StartSurah int
	StartAyah  int
	EndSurah   int
	EndAyah    int
}

// CreateJuz creates a new juz
func (h *CommandHandler) CreateJuz(ctx context.Context, cmd CreateJuzCommand) (*quran.Juz, error) {
	// Validate start surah
	if cmd.StartSurah < 1 || cmd.StartSurah > 114 {
		return nil, fmt.Errorf("invalid start surah: %d", cmd.StartSurah)
	}

	// Validate end surah
	if cmd.EndSurah < 1 || cmd.EndSurah > 114 {
		return nil, fmt.Errorf("invalid end surah: %d", cmd.EndSurah)
	}

	// Validate start ayah
	if cmd.StartAyah < 1 {
		return nil, fmt.Errorf("invalid start ayah: %d", cmd.StartAyah)
	}

	// Validate end ayah
	if cmd.EndAyah < 1 {
		return nil, fmt.Errorf("invalid end ayah: %d", cmd.EndAyah)
	}

	juz := &quran.Juz{
		StartSurah: cmd.StartSurah,
		StartAyah:  cmd.StartAyah,
		EndSurah:   cmd.EndSurah,
		EndAyah:    cmd.EndAyah,
	}

	var result *quran.Juz
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.quranService.CreateJuz(ctx, juz)
		return err
	})

	return result, err
}
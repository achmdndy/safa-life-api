package command

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/domain/quran"
)

// CreateSurahCommand represents the command to create a new surah
type CreateSurahCommand struct {
	NameArabic      string
	NameEnglish     string
	RevelationPlace string
	RevelationOrder int
	AyahCount       int
}

// CreateSurah creates a new surah
func (h *CommandHandler) CreateSurah(ctx context.Context, cmd CreateSurahCommand) (*quran.Surah, error) {
	// Validate revelation place
	if cmd.RevelationPlace != "Mecca" && cmd.RevelationPlace != "Medina" {
		return nil, fmt.Errorf("invalid revelation place: %s", cmd.RevelationPlace)
	}

	// Validate revelation order (1-114)
	if cmd.RevelationOrder < 1 || cmd.RevelationOrder > 114 {
		return nil, fmt.Errorf("invalid revelation order: %d", cmd.RevelationOrder)
	}

	// Validate ayah count
	if cmd.AyahCount < 1 {
		return nil, fmt.Errorf("invalid ayah count: %d", cmd.AyahCount)
	}

	surah := &quran.Surah{
		NameArabic:      cmd.NameArabic,
		NameEnglish:     cmd.NameEnglish,
		RevelationPlace: cmd.RevelationPlace,
		RevelationOrder: cmd.RevelationOrder,
		AyahCount:       cmd.AyahCount,
	}

	var result *quran.Surah
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.quranService.CreateSurah(ctx, surah)
		return err
	})

	return result, err
}
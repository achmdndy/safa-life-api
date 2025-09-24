package command

import (
	"context"
	"fmt"

	"github.com/achmdndy/safa-life-api/src/domain/quran"
)

// UpdateSurahCommand represents the command to update a surah
type UpdateSurahCommand struct {
	ID              int
	NameArabic      *string
	NameEnglish     *string
	RevelationPlace *string
	RevelationOrder *int
	AyahCount       *int
}

// UpdateSurah updates an existing surah
func (h *CommandHandler) UpdateSurah(ctx context.Context, cmd UpdateSurahCommand) (*quran.Surah, error) {
	// Validate ID
	if cmd.ID < 1 || cmd.ID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", cmd.ID)
	}

	// Validate revelation place if provided
	if cmd.RevelationPlace != nil && *cmd.RevelationPlace != "Mecca" && *cmd.RevelationPlace != "Medina" {
		return nil, fmt.Errorf("invalid revelation place: %s", *cmd.RevelationPlace)
	}

	// Validate revelation order if provided
	if cmd.RevelationOrder != nil && (*cmd.RevelationOrder < 1 || *cmd.RevelationOrder > 114) {
		return nil, fmt.Errorf("invalid revelation order: %d", *cmd.RevelationOrder)
	}

	// Validate ayah count if provided
	if cmd.AyahCount != nil && *cmd.AyahCount < 1 {
		return nil, fmt.Errorf("invalid ayah count: %d", *cmd.AyahCount)
	}

	// Get existing surah
	existingSurah, err := h.quranService.GetSurahByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if cmd.NameArabic != nil {
		existingSurah.NameArabic = *cmd.NameArabic
	}
	if cmd.NameEnglish != nil {
		existingSurah.NameEnglish = *cmd.NameEnglish
	}
	if cmd.RevelationPlace != nil {
		existingSurah.RevelationPlace = *cmd.RevelationPlace
	}
	if cmd.RevelationOrder != nil {
		existingSurah.RevelationOrder = *cmd.RevelationOrder
	}
	if cmd.AyahCount != nil {
		existingSurah.AyahCount = *cmd.AyahCount
	}

	var result *quran.Surah
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.quranService.UpdateSurah(ctx, existingSurah)
		return err
	})

	return result, err
}

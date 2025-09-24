package command

import (
	"context"
	"fmt"
	"github.com/achmdndy/safa-life-api/src/domain/quran"
)

// CreateAyahCommand represents the command to create a new ayah
type CreateAyahCommand struct {
	SurahID      int
	AyahID       int
	Text         string
	PageNumber   int
	JuzNumber    int
	HizbNumber   int
	ManzilNumber int
}

// CreateAyah creates a new ayah
func (h *CommandHandler) CreateAyah(ctx context.Context, cmd CreateAyahCommand) (*quran.Ayah, error) {
	// Validate surah ID
	if cmd.SurahID < 1 || cmd.SurahID > 114 {
		return nil, fmt.Errorf("invalid surah ID: %d", cmd.SurahID)
	}

	// Validate ayah ID
	if cmd.AyahID < 1 {
		return nil, fmt.Errorf("invalid ayah ID: %d", cmd.AyahID)
	}

	// Validate page number
	if cmd.PageNumber < 1 || cmd.PageNumber > 604 {
		return nil, fmt.Errorf("invalid page number: %d", cmd.PageNumber)
	}

	// Validate juz number
	if cmd.JuzNumber < 1 || cmd.JuzNumber > 30 {
		return nil, fmt.Errorf("invalid juz number: %d", cmd.JuzNumber)
	}

	// Validate hizb number
	if cmd.HizbNumber < 1 || cmd.HizbNumber > 60 {
		return nil, fmt.Errorf("invalid hizb number: %d", cmd.HizbNumber)
	}

	// Validate manzil number
	if cmd.ManzilNumber < 1 || cmd.ManzilNumber > 7 {
		return nil, fmt.Errorf("invalid manzil number: %d", cmd.ManzilNumber)
	}

	ayah := &quran.Ayah{
		SurahID:      cmd.SurahID,
		AyahID:       cmd.AyahID,
		Text:         cmd.Text,
		PageNumber:   cmd.PageNumber,
		JuzNumber:    cmd.JuzNumber,
		HizbNumber:   cmd.HizbNumber,
		ManzilNumber: cmd.ManzilNumber,
	}

	var result *quran.Ayah
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.quranService.CreateAyah(ctx, ayah)
		return err
	})

	return result, err
}
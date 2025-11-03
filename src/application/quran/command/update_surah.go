package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type UpdateSurahCommand struct {
	ID              string
	NameArabic      string
	NameEnglish     string
	RevelationPlace string
	RevelationOrder int
	AyahCount       int
	UpdatedBy       string
}

func (h *CommandHandler) UpdateSurah(ctx context.Context, command UpdateSurahCommand) (*dto.SurahResponse, error) {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return nil, err
	}

	// Get existing surah
	existingSurah, err := h.surahService.GetSurahById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update the surah
	existingSurah.Update(
		command.NameArabic,
		command.NameEnglish,
		command.RevelationPlace,
		command.RevelationOrder,
		command.AyahCount,
		command.UpdatedBy,
	)

	var updatedSurah *quran.Surah
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		updatedSurah, txErr = h.surahService.UpdateSurah(ctx, existingSurah)
		return txErr
	})

	if err != nil {
		return nil, err
	}

	return dto.ToSurahResponse(updatedSurah), nil
}

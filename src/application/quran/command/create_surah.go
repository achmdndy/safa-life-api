package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type CreateSurahCommand struct {
	NameArabic      string
	NameEnglish     string
	RevelationPlace string
	RevelationOrder int
	AyahCount       int
	CreatedBy       string
}

func (h *CommandHandler) CreateSurah(ctx context.Context, command CreateSurahCommand) (*dto.SurahResponse, error) {
	id := h.uuidGenerator.New()

	surah := quran.NewSurah(
		id,
		command.NameArabic,
		command.NameEnglish,
		command.RevelationPlace,
		command.RevelationOrder,
		command.AyahCount,
		command.CreatedBy,
	)

	var createdSurah *quran.Surah
	err := h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		createdSurah, err = h.surahService.CreateSurah(ctx, surah)
		return err
	})

	if err != nil {
		return nil, err
	}

	return dto.ToSurahResponse(createdSurah), nil
}

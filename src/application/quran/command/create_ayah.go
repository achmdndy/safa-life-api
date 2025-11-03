package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type CreateAyahCommand struct {
	SurahID      string
	Text         string
	PageNumber   int
	JuzNumber    int
	HizbNumber   int
	ManzilNumber int
	CreatedBy    string
}

func (h *CommandHandler) CreateAyah(ctx context.Context, command CreateAyahCommand) (*dto.AyahResponse, error) {
	id := h.uuidGenerator.New()

	surahID, err := h.uuidGenerator.Parse(command.SurahID)
	if err != nil {
		return nil, err
	}

	ayah := quran.NewAyah(
		id,
		surahID,
		command.Text,
		command.PageNumber,
		command.JuzNumber,
		command.HizbNumber,
		command.ManzilNumber,
		command.CreatedBy,
	)

	var createdAyah *quran.Ayah
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		createdAyah, txErr = h.ayahService.CreateAyah(ctx, ayah)
		return txErr
	})

	if err != nil {
		return nil, err
	}

	return dto.ToAyahResponse(createdAyah), nil
}
package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type UpdateAyahCommand struct {
	ID           string
	Text         string
	PageNumber   int
	JuzNumber    int
	HizbNumber   int
	ManzilNumber int
	UpdatedBy    string
}

func (h *CommandHandler) UpdateAyah(ctx context.Context, command UpdateAyahCommand) (*dto.AyahResponse, error) {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return nil, err
	}

	// Get existing ayah
	existingAyah, err := h.ayahService.GetAyahById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update the ayah
	existingAyah.Update(
		command.Text,
		command.PageNumber,
		command.JuzNumber,
		command.HizbNumber,
		command.ManzilNumber,
		command.UpdatedBy,
	)

	var updatedAyah *quran.Ayah
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		updatedAyah, txErr = h.ayahService.UpdateAyah(ctx, existingAyah)
		return txErr
	})

	if err != nil {
		return nil, err
	}

	return dto.ToAyahResponse(updatedAyah), nil
}
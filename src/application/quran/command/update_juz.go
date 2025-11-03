package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type UpdateJuzCommand struct {
	ID           string
	StartSurahID string
	EndSurahID   string
	StartAyahID  string
	EndAyahID    string
	UpdatedBy    string
}

func (h *CommandHandler) UpdateJuz(ctx context.Context, command UpdateJuzCommand) (*dto.JuzResponse, error) {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return nil, err
	}

	// Get existing juz
	existingJuz, err := h.juzService.GetJuzById(ctx, id)
	if err != nil {
		return nil, err
	}

	startSurahID, err := h.uuidGenerator.Parse(command.StartSurahID)
	if err != nil {
		return nil, err
	}

	endSurahID, err := h.uuidGenerator.Parse(command.EndSurahID)
	if err != nil {
		return nil, err
	}

	startAyahID, err := h.uuidGenerator.Parse(command.StartAyahID)
	if err != nil {
		return nil, err
	}

	endAyahID, err := h.uuidGenerator.Parse(command.EndAyahID)
	if err != nil {
		return nil, err
	}

	// Update the juz
	existingJuz.Update(
		startSurahID,
		endSurahID,
		startAyahID,
		endAyahID,
		command.UpdatedBy,
	)

	var updatedJuz *quran.Juz
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		updatedJuz, txErr = h.juzService.UpdateJuz(ctx, existingJuz)
		return txErr
	})

	if err != nil {
		return nil, err
	}

	return dto.ToJuzResponse(updatedJuz), nil
}
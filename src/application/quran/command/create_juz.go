package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type CreateJuzCommand struct {
	StartSurahID string
	EndSurahID   string
	StartAyahID  string
	EndAyahID    string
	CreatedBy    string
}

func (h *CommandHandler) CreateJuz(ctx context.Context, command CreateJuzCommand) (*dto.JuzResponse, error) {
	id := h.uuidGenerator.New()

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

	juz := quran.NewJuz(
		id,
		startSurahID,
		endSurahID,
		startAyahID,
		endAyahID,
		command.CreatedBy,
	)

	var createdJuz *quran.Juz
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		createdJuz, txErr = h.juzService.CreateJuz(ctx, juz)
		return txErr
	})

	if err != nil {
		return nil, err
	}

	return dto.ToJuzResponse(createdJuz), nil
}

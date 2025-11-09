package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type CreateLastReadCommand struct {
	UserID      string
	SurahID     string
	AyahID      string
	AyahNumber  int
	ProgressPct float64
	CreatedBy   string
}

func (h *CommandHandler) CreateLastRead(ctx context.Context, command CreateLastReadCommand) (*dto.LastReadResponse, error) {
	id := h.uuidGenerator.New()

	userID, err := h.uuidGenerator.Parse(command.UserID)
	if err != nil {
		return nil, err
	}
	surahID, err := h.uuidGenerator.Parse(command.SurahID)
	if err != nil {
		return nil, err
	}
	ayahID, err := h.uuidGenerator.Parse(command.AyahID)
	if err != nil {
		return nil, err
	}

	lr := quran.NewLastRead(
		id,
		userID,
		surahID,
		ayahID,
		command.AyahNumber,
		command.ProgressPct,
		command.CreatedBy,
	)

	var created *quran.LastRead
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		created, txErr = h.lastReadService.CreateLastRead(ctx, lr)
		return txErr
	})
	if err != nil {
		return nil, err
	}

	return dto.ToLastReadResponse(created), nil
}

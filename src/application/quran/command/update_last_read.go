package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type UpdateLastReadCommand struct {
	ID          string
	AyahID      string
	AyahNumber  int
	ProgressPct float64
	UpdatedBy   string
}

func (h *CommandHandler) UpdateLastRead(ctx context.Context, command UpdateLastReadCommand) (*dto.LastReadResponse, error) {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return nil, err
	}

	existing, err := h.lastReadService.GetLastReadById(ctx, id)
	if err != nil {
		return nil, err
	}

	ayahID, err := h.uuidGenerator.Parse(command.AyahID)
	if err != nil {
		return nil, err
	}

	existing.Update(ayahID, command.AyahNumber, command.ProgressPct, command.UpdatedBy)

	var updated *quran.LastRead
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		updated, txErr = h.lastReadService.UpdateLastRead(ctx, existing)
		return txErr
	})
	if err != nil {
		return nil, err
	}

	return dto.ToLastReadResponse(updated), nil
}

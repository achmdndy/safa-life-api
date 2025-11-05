package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type UpdateReciterCommand struct {
	ID        string
	Name      string
	Style     string
	UpdatedBy string
}

func (h *CommandHandler) UpdateReciter(ctx context.Context, command UpdateReciterCommand) (*dto.ReciterResponse, error) {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return nil, err
	}

	existing, err := h.reciterService.GetReciterById(ctx, id)
	if err != nil {
		return nil, err
	}

	existing.Update(command.Name, command.Style, command.UpdatedBy)

	var updated *quran.Reciter
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		updated, txErr = h.reciterService.UpdateReciter(ctx, existing)
		return txErr
	})
	if err != nil {
		return nil, err
	}

	return dto.ToReciterResponse(updated), nil
}

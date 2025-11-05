package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type CreateReciterCommand struct {
	Name      string
	Style     string
	CreatedBy string
}

func (h *CommandHandler) CreateReciter(ctx context.Context, command CreateReciterCommand) (*dto.ReciterResponse, error) {
	id := h.uuidGenerator.New()

	reciter := quran.NewReciter(
		id,
		command.Name,
		command.Style,
		command.CreatedBy,
	)

	var createdReciter *quran.Reciter
	err := h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		createdReciter, txErr = h.reciterService.CreateReciter(ctx, reciter)
		return txErr
	})

	if err != nil {
		return nil, err
	}

	return dto.ToReciterResponse(createdReciter), nil
}

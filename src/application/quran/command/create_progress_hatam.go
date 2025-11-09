package command

import (
	"context"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type CreateProgressHatamCommand struct {
	UserID      string
	JuzID       string
	StartAyahID string
	ProgressPct float64
	CreatedBy   string
}

func (h *CommandHandler) CreateProgressHatam(ctx context.Context, command CreateProgressHatamCommand) (*dto.ProgressHatamResponse, error) {
	id := h.uuidGenerator.New()

	userID, err := h.uuidGenerator.Parse(command.UserID)
	if err != nil {
		return nil, err
	}
	juzID, err := h.uuidGenerator.Parse(command.JuzID)
	if err != nil {
		return nil, err
	}
	startAyahID, err := h.uuidGenerator.Parse(command.StartAyahID)
	if err != nil {
		return nil, err
	}

	p := quran.NewProgressHatam(
		id,
		userID,
		juzID,
		startAyahID,
		command.ProgressPct,
		command.CreatedBy,
	)

	var created *quran.ProgressHatam
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		created, txErr = h.progressService.CreateProgressHatam(ctx, p)
		return txErr
	})
	if err != nil {
		return nil, err
	}

	return dto.ToProgressHatamResponse(created), nil
}

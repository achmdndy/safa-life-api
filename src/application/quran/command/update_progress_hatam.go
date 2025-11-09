package command

import (
	"context"
	"time"

	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/domain/quran"
)

type UpdateProgressHatamCommand struct {
	ID          string
	LastAyahID  string
	ProgressPct float64
	IsCompleted bool
	CompletedAt string
	UpdatedBy   string
}

func (h *CommandHandler) UpdateProgressHatam(ctx context.Context, command UpdateProgressHatamCommand) (*dto.ProgressHatamResponse, error) {
	id, err := h.uuidGenerator.Parse(command.ID)
	if err != nil {
		return nil, err
	}

	lastAyahID, err := h.uuidGenerator.Parse(command.LastAyahID)
	if err != nil {
		return nil, err
	}

	p, err := h.progressService.GetProgressHatamById(ctx, id)
	if err != nil {
		return nil, err
	}

	var completedAt *time.Time
	if command.CompletedAt != "" {
		// Try to parse RFC3339 timestamp; if parsing fails, leave nil and let service validate as needed
		if t, parseErr := time.Parse(time.RFC3339, command.CompletedAt); parseErr == nil {
			completedAt = &t
		} else {
			// If parsing fails, return error to the caller to keep consistency
			return nil, parseErr
		}
	}

	p.Update(lastAyahID, command.ProgressPct, command.IsCompleted, completedAt, command.UpdatedBy)

	var updated *quran.ProgressHatam
	err = h.transactionMgr.WithTransaction(ctx, func(ctx context.Context) error {
		var txErr error
		updated, txErr = h.progressService.UpdateProgressHatam(ctx, p)
		return txErr
	})
	if err != nil {
		return nil, err
	}

	return dto.ToProgressHatamResponse(updated), nil
}

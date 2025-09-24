package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/reciter"
)

// UpdateReciterCommand defines the command for updating a reciter.
type UpdateReciterCommand struct {
	ID    string
	Name  *string
	Style *string
}

// UpdateReciter handles the update of a reciter.
func (h *CommandHandler) UpdateReciter(ctx context.Context, cmd UpdateReciterCommand) (*reciter.Reciter, error) {
	existing, err := h.reciterService.GetReciterByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if cmd.Name != nil {
		existing.Name = *cmd.Name
	}
	if cmd.Style != nil {
		existing.Style = *cmd.Style
	}

	var result *reciter.Reciter
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.reciterService.UpdateReciter(ctx, existing)
		return err
	})

	return result, err
}

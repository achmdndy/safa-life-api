package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/reciter"
)

// CreateReciterCommand defines the command for creating a new reciter.
type CreateReciterCommand struct {
	ID    string
	Name  string
	Style string
}

// CreateReciter handles the creation of a reciter.
func (h *CommandHandler) CreateReciter(ctx context.Context, cmd CreateReciterCommand) (*reciter.Reciter, error) {
	domain := &reciter.Reciter{
		ID:    cmd.ID,
		Name:  cmd.Name,
		Style: cmd.Style,
	}

	var result *reciter.Reciter
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.reciterService.CreateReciter(ctx, domain)
		return err
	})

	return result, err
}

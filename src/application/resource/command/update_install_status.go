package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/resource"
)

type UpdateInstallStatusCommand struct {
	ID          string
	IsInstalled bool
}

func (h *CommandHandler) UpdateInstallStatus(ctx context.Context, cmd UpdateInstallStatusCommand) (*resource.Resource, error) {
	var result *resource.Resource
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.resourceService.UpdateInstallStatus(ctx, cmd.ID, cmd.IsInstalled)
		return err
	})

	return result, err
}

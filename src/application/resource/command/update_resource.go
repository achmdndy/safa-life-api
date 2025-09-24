package command

import (
	"context"
	"time"

	"github.com/achmdndy/safa-life-api/src/domain/resource"
)

type UpdateResourceCommand struct {
	ID            string
	Type          *string
	Name          *string
	Language      *string
	Version       *string
	LastUpdatedAt *time.Time
	DownloadURL   *string
	Size          *int64
}

func (h *CommandHandler) UpdateResource(ctx context.Context, cmd UpdateResourceCommand) (*resource.Resource, error) {
	existing, err := h.resourceService.GetResourceByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if cmd.Type != nil {
		existing.Type = *cmd.Type
	}
	if cmd.Name != nil {
		existing.Name = *cmd.Name
	}
	if cmd.Language != nil {
		existing.Language = *cmd.Language
	}
	if cmd.Version != nil {
		existing.Version = *cmd.Version
	}
	if cmd.LastUpdatedAt != nil {
		existing.LastUpdatedAt = *cmd.LastUpdatedAt
	}
	if cmd.DownloadURL != nil {
		existing.DownloadURL = *cmd.DownloadURL
	}
	if cmd.Size != nil {
		existing.Size = *cmd.Size
	}

	var result *resource.Resource
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.resourceService.UpdateResource(ctx, existing)
		return err
	})

	return result, err
}

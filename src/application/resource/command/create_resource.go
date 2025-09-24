package command

import (
	"context"
	"time"

	"github.com/achmdndy/safa-life-api/src/domain/resource"
)

type CreateResourceCommand struct {
	ID            string
	Type          string
	Name          string
	Language      string
	Version       string
	LastUpdatedAt time.Time
	DownloadURL   string
	Size          int64
}

func (h *CommandHandler) CreateResource(ctx context.Context, cmd CreateResourceCommand) (*resource.Resource, error) {
	domain := &resource.Resource{
		ID:            cmd.ID,
		Type:          cmd.Type,
		Name:          cmd.Name,
		Language:      cmd.Language,
		Version:       cmd.Version,
		LastUpdatedAt: cmd.LastUpdatedAt,
		DownloadURL:   cmd.DownloadURL,
		Size:          cmd.Size,
	}

	var result *resource.Resource
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.resourceService.CreateResource(ctx, domain)
		return err
	})

	return result, err
}

package command

import (
	"context"
)

type DeleteStoryCommand struct {
	ID string
}

func (h *CommandHandler) DeleteStory(ctx context.Context, cmd DeleteStoryCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.storyService.DeleteStory(ctx, cmd.ID)
	})
}

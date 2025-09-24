package command

import (
	"context"
)

type DeleteTopicCommand struct {
	ID string
}

func (h *CommandHandler) DeleteTopic(ctx context.Context, cmd DeleteTopicCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.topicService.DeleteTopic(ctx, cmd.ID)
	})
}

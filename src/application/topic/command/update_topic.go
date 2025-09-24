package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/topic"
)

type UpdateTopicCommand struct {
	ID          string
	Name        *string
	Description *string
}

func (h *CommandHandler) UpdateTopic(ctx context.Context, cmd UpdateTopicCommand) (*topic.Topic, error) {
	existing, err := h.topicService.GetTopicByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if cmd.Name != nil {
		existing.Name = *cmd.Name
	}
	if cmd.Description != nil {
		existing.Description = *cmd.Description
	}

	var result *topic.Topic
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.topicService.UpdateTopic(ctx, existing)
		return err
	})

	return result, err
}

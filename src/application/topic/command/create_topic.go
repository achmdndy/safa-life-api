package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/topic"
)

type CreateTopicCommand struct {
	ID          string
	Name        string
	Description string
}

func (h *CommandHandler) CreateTopic(ctx context.Context, cmd CreateTopicCommand) (*topic.Topic, error) {
	domain := &topic.Topic{
		ID:          cmd.ID,
		Name:        cmd.Name,
		Description: cmd.Description,
	}

	var result *topic.Topic
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.topicService.CreateTopic(ctx, domain)
		return err
	})

	return result, err
}

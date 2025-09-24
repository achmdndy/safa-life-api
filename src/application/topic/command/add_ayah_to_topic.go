package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/topic"
)

type AddAyahToTopicCommand struct {
	TopicID string
	SurahID int
	AyahID  int
}

func (h *CommandHandler) AddAyahToTopic(ctx context.Context, cmd AddAyahToTopicCommand) (*topic.TopicAyah, error) {
	domain := &topic.TopicAyah{
		TopicID: cmd.TopicID,
		SurahID: cmd.SurahID,
		AyahID:  cmd.AyahID,
	}

	var result *topic.TopicAyah
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.topicService.AddAyahToTopic(ctx, domain)
		return err
	})

	return result, err
}

package command

import (
	"context"
)

type RemoveAyahFromTopicCommand struct {
	TopicID string
	SurahID int
	AyahID  int
}

func (h *CommandHandler) RemoveAyahFromTopic(ctx context.Context, cmd RemoveAyahFromTopicCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.topicService.RemoveAyahFromTopic(ctx, cmd.TopicID, cmd.SurahID, cmd.AyahID)
	})
}

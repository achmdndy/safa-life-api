package command

import (
	"context"
)

type RemoveAyahFromStoryCommand struct {
	StoryID string
	SurahID int
	AyahID  int
}

func (h *CommandHandler) RemoveAyahFromStory(ctx context.Context, cmd RemoveAyahFromStoryCommand) error {
	return h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		return h.storyService.RemoveAyahFromStory(ctx, cmd.StoryID, cmd.SurahID, cmd.AyahID)
	})
}

package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/story"
)

type AddAyahToStoryCommand struct {
	StoryID string
	SurahID int
	AyahID  int
}

func (h *CommandHandler) AddAyahToStory(ctx context.Context, cmd AddAyahToStoryCommand) (*story.StoryAyah, error) {
	domain := &story.StoryAyah{
		StoryID: cmd.StoryID,
		SurahID: cmd.SurahID,
		AyahID:  cmd.AyahID,
	}

	var result *story.StoryAyah
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.storyService.AddAyahToStory(ctx, domain)
		return err
	})

	return result, err
}

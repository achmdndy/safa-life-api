package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/story"
)

type CreateStoryCommand struct {
	ID         string
	Title      string
	Summary    string
	Characters []string
}

func (h *CommandHandler) CreateStory(ctx context.Context, cmd CreateStoryCommand) (*story.Story, error) {
	domain := &story.Story{
		ID:         cmd.ID,
		Title:      cmd.Title,
		Summary:    cmd.Summary,
		Characters: cmd.Characters,
	}

	var result *story.Story
	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		var err error
		result, err = h.storyService.CreateStory(ctx, domain)
		return err
	})

	return result, err
}

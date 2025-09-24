package command

import (
	"context"

	"github.com/achmdndy/safa-life-api/src/domain/story"
)

type UpdateStoryCommand struct {
	ID         string
	Title      *string
	Summary    *string
	Characters *[]string
}

func (h *CommandHandler) UpdateStory(ctx context.Context, cmd UpdateStoryCommand) (*story.Story, error) {
	existing, err := h.storyService.GetStoryByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if cmd.Title != nil {
		existing.Title = *cmd.Title
	}
	if cmd.Summary != nil {
		existing.Summary = *cmd.Summary
	}
	if cmd.Characters != nil {
		existing.Characters = *cmd.Characters
	}

	var result *story.Story
	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		result, err = h.storyService.UpdateStory(ctx, existing)
		return err
	})

	return result, err
}

package command

import (
	"github.com/achmdndy/safa-life-api/src/domain/story"
)

// CommandHandler handles all write operations for the Story module.
type CommandHandler struct {
	storyService story.StoryService
	txManager    story.TransactionManager
}

// NewCommandHandler creates a new command handler.
func NewCommandHandler(service story.StoryService, txManager story.TransactionManager) *CommandHandler {
	return &CommandHandler{
		storyService: service,
		txManager:    txManager,
	}
}

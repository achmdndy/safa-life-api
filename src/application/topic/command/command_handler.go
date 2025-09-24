package command

import (
	"github.com/achmdndy/safa-life-api/src/domain/topic"
)

// CommandHandler handles all write operations for the Topic module.
type CommandHandler struct {
	topicService topic.TopicService
	txManager    topic.TransactionManager
}

// NewCommandHandler creates a new command handler.
func NewCommandHandler(service topic.TopicService, txManager topic.TransactionManager) *CommandHandler {
	return &CommandHandler{
		topicService: service,
		txManager:    txManager,
	}
}

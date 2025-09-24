package command

import (
	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
)

// CommandHandler handles all write operations for Tajweed module
type CommandHandler struct {
	tajweedService tajweed.TajweedService
	txManager      tajweed.TransactionManager
}

// NewCommandHandler creates a new command handler
func NewCommandHandler(tajweedService tajweed.TajweedService, txManager tajweed.TransactionManager) *CommandHandler {
	return &CommandHandler{
		tajweedService: tajweedService,
		txManager:      txManager,
	}
}

package command

import (
	"github.com/achmdndy/safa-life-api/src/domain/reciter"
)

// CommandHandler handles all write operations for the Reciter module.
type CommandHandler struct {
	reciterService reciter.ReciterService
	txManager      reciter.TransactionManager
}

// NewCommandHandler creates a new command handler.
func NewCommandHandler(service reciter.ReciterService, txManager reciter.TransactionManager) *CommandHandler {
	return &CommandHandler{
		reciterService: service,
		txManager:      txManager,
	}
}

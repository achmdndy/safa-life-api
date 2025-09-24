package command

import (
	"github.com/achmdndy/safa-life-api/src/domain/audio"
)

// CommandHandler handles all write operations for the Audio module.
type CommandHandler struct {
	audioService audio.AudioService
	txManager    audio.TransactionManager
}

// NewCommandHandler creates a new command handler.
func NewCommandHandler(service audio.AudioService, txManager audio.TransactionManager) *CommandHandler {
	return &CommandHandler{
		audioService: service,
		txManager:    txManager,
	}
}

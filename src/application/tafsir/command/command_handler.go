package command

import (
	"github.com/achmdndy/safa-life-api/src/domain/tafsir"
)

// CommandHandler handles all write operations for the Tafsir module.
type CommandHandler struct {
	tafsirService tafsir.TafsirService
	txManager     tafsir.TransactionManager
}

// NewCommandHandler creates a new command handler.
func NewCommandHandler(service tafsir.TafsirService, txManager tafsir.TransactionManager) *CommandHandler {
	return &CommandHandler{
		tafsirService: service,
		txManager:     txManager,
	}
}

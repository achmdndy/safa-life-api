package command

import (
	"github.com/achmdndy/safa-life-api/src/domain/translation"
)

// CommandHandler handles all write operations for the Translation module.
type CommandHandler struct {
	translationService translation.TranslationService
	txManager          translation.TransactionManager
}

// NewCommandHandler creates a new command handler.
func NewCommandHandler(service translation.TranslationService, txManager translation.TransactionManager) *CommandHandler {
	return &CommandHandler{
		translationService: service,
		txManager:          txManager,
	}
}

package command

import (
	"github.com/achmdndy/safa-life-api/src/domain/quran"
)

// CommandHandler handles all write operations for Quran module
type CommandHandler struct {
	quranService quran.QuranService
	txManager    quran.TransactionManager
}

// NewCommandHandler creates a new command handler
func NewCommandHandler(quranService quran.QuranService, txManager quran.TransactionManager) *CommandHandler {
	return &CommandHandler{
		quranService: quranService,
		txManager:    txManager,
	}
}

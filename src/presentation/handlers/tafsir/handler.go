package tafsir

import (
	"github.com/achmdndy/safa-life-api/src/application/tafsir/command"
	"github.com/achmdndy/safa-life-api/src/application/tafsir/query"
)

// Handler handles HTTP requests for the Tafsir module.
type Handler struct {
	queryHandler   *query.QueryHandler
	commandHandler *command.CommandHandler
}

// NewHandler creates a new Tafsir handler.
func NewHandler(queryHandler *query.QueryHandler, commandHandler *command.CommandHandler) *Handler {
	return &Handler{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}
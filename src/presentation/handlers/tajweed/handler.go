package tajweed

import (
	"github.com/achmdndy/safa-life-api/src/application/tajweed/command"
	"github.com/achmdndy/safa-life-api/src/application/tajweed/query"
)

// Handler handles HTTP requests for Tajweed module
type Handler struct {
	queryHandler   *query.QueryHandler
	commandHandler *command.CommandHandler
}

// NewHandler creates a new Tajweed handler
func NewHandler(queryHandler *query.QueryHandler, commandHandler *command.CommandHandler) *Handler {
	return &Handler{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}

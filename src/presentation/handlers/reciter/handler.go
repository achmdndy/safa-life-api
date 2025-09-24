package reciter

import (
	"github.com/achmdndy/safa-life-api/src/application/reciter/command"
	"github.com/achmdndy/safa-life-api/src/application/reciter/query"
)

// Handler handles HTTP requests for the Reciter module.
type Handler struct {
	queryHandler   *query.QueryHandler
	commandHandler *command.CommandHandler
}

// NewHandler creates a new Reciter handler.
func NewHandler(queryHandler *query.QueryHandler, commandHandler *command.CommandHandler) *Handler {
	return &Handler{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}

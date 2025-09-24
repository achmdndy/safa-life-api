package audio

import (
	"github.com/achmdndy/safa-life-api/src/application/audio/command"
	"github.com/achmdndy/safa-life-api/src/application/audio/query"
)

// Handler handles HTTP requests for the Audio module.
type Handler struct {
	queryHandler   *query.QueryHandler
	commandHandler *command.CommandHandler
}

// NewHandler creates a new Audio handler.
func NewHandler(queryHandler *query.QueryHandler, commandHandler *command.CommandHandler) *Handler {
	return &Handler{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}
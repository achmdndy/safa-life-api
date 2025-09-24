package resource

import (
	"github.com/achmdndy/safa-life-api/src/application/resource/command"
	"github.com/achmdndy/safa-life-api/src/application/resource/query"
)

// Handler handles HTTP requests for the Resource module.
type Handler struct {
	queryHandler   *query.QueryHandler
	commandHandler *command.CommandHandler
}

// NewHandler creates a new Resource handler.
func NewHandler(queryHandler *query.QueryHandler, commandHandler *command.CommandHandler) *Handler {
	return &Handler{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}
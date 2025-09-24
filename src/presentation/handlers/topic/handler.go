package topic

import (
	"github.com/achmdndy/safa-life-api/src/application/topic/command"
	"github.com/achmdndy/safa-life-api/src/application/topic/query"
)

// Handler handles HTTP requests for the Topic module.
type Handler struct {
	queryHandler   *query.QueryHandler
	commandHandler *command.CommandHandler
}

// NewHandler creates a new Topic handler.
func NewHandler(queryHandler *query.QueryHandler, commandHandler *command.CommandHandler) *Handler {
	return &Handler{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}
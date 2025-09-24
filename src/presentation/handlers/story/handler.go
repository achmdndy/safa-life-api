package story

import (
	"github.com/achmdndy/safa-life-api/src/application/story/command"
	"github.com/achmdndy/safa-life-api/src/application/story/query"
)

// Handler handles HTTP requests for the Story module.
type Handler struct {
	queryHandler   *query.QueryHandler
	commandHandler *command.CommandHandler
}

// NewHandler creates a new Story handler.
func NewHandler(queryHandler *query.QueryHandler, commandHandler *command.CommandHandler) *Handler {
	return &Handler{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}
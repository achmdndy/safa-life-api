package translation

import (
	"github.com/achmdndy/safa-life-api/src/application/translation/command"
	"github.com/achmdndy/safa-life-api/src/application/translation/query"
)

// Handler handles HTTP requests for the Translation module.
type Handler struct {
	queryHandler   *query.QueryHandler
	commandHandler *command.CommandHandler
}

// NewHandler creates a new Translation handler.
func NewHandler(queryHandler *query.QueryHandler, commandHandler *command.CommandHandler) *Handler {
	return &Handler{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}

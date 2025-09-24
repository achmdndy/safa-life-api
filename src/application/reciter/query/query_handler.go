package query

import (
	"github.com/achmdndy/safa-life-api/src/domain/reciter"
)

// QueryHandler handles all read operations for the Reciter module.
type QueryHandler struct {
	reciterService reciter.ReciterService
}

// NewQueryHandler creates a new query handler.
func NewQueryHandler(service reciter.ReciterService) *QueryHandler {
	return &QueryHandler{
		reciterService: service,
	}
}

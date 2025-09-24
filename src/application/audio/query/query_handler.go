package query

import (
	"github.com/achmdndy/safa-life-api/src/domain/audio"
)

// QueryHandler handles all read operations for the Audio module.
type QueryHandler struct {
	audioService audio.AudioService
}

// NewQueryHandler creates a new query handler.
func NewQueryHandler(service audio.AudioService) *QueryHandler {
	return &QueryHandler{
		audioService: service,
	}
}

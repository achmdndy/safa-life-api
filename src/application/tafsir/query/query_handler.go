package query

import (
	"github.com/achmdndy/safa-life-api/src/domain/tafsir"
)

// QueryHandler handles all read operations for the Tafsir module.
type QueryHandler struct {
	tafsirService tafsir.TafsirService
}

// NewQueryHandler creates a new query handler.
func NewQueryHandler(service tafsir.TafsirService) *QueryHandler {
	return &QueryHandler{
		tafsirService: service,
	}
}

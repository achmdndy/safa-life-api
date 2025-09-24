package query

import (
	"github.com/achmdndy/safa-life-api/src/domain/translation"
)

// QueryHandler handles all read operations for the Translation module.
type QueryHandler struct {
	translationService translation.TranslationService
}

// NewQueryHandler creates a new query handler.
func NewQueryHandler(service translation.TranslationService) *QueryHandler {
	return &QueryHandler{
		translationService: service,
	}
}

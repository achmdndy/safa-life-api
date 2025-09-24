package query

import (
	"github.com/achmdndy/safa-life-api/src/domain/resource"
)

// QueryHandler handles all read operations for the Resource module.
type QueryHandler struct {
	resourceService resource.ResourceService
}

// NewQueryHandler creates a new query handler.
func NewQueryHandler(service resource.ResourceService) *QueryHandler {
	return &QueryHandler{
		resourceService: service,
	}
}

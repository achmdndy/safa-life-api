package query

import (
	"github.com/achmdndy/safa-life-api/src/domain/tajweed"
)

type QueryHandler struct {
	tajweedService tajweed.TajweedService
}

func NewQueryHandler(tajweedService tajweed.TajweedService) *QueryHandler {
	return &QueryHandler{
		tajweedService: tajweedService,
	}
}

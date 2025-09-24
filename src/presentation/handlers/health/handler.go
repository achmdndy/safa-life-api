package health

import (
	"github.com/achmdndy/safa-life-api/src/application/health/interfaces"
)

type Handler struct {
	healthQueryHandler interfaces.QueryHandler
}

func NewHandler(healthQueryHandler interfaces.QueryHandler) *Handler {
	return &Handler{
		healthQueryHandler: healthQueryHandler,
	}
}
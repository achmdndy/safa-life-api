package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/achmdndy/safa-life-api/src/application/health/dto"
	"github.com/achmdndy/safa-life-api/src/application/health/interfaces"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

type Handler struct {
	healthQueryHandler interfaces.QueryHandler
}

func NewHandler(healthQueryHandler interfaces.QueryHandler) *Handler {
	return &Handler{
		healthQueryHandler: healthQueryHandler,
	}
}

func (h *Handler) GetHealth(c *gin.Context) {
	start := time.Now()

	// Create a simple query object
	query := struct{}{}

	result, err := h.healthQueryHandler.Handle(query)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get health status", &core.ErrorDetail{
			Reason: err.Error(),
		}, start)
		return
	}

	// Type assertion to get health status
	healthStatus, ok := result.(dto.HealthResponse)
	if !ok {
		core.Error(c, http.StatusInternalServerError, "Invalid health status format", &core.ErrorDetail{
			Reason: "Failed to parse health status",
		}, start)
		return
	}

	// Determine HTTP status based on health status
	statusCode := http.StatusOK
	switch healthStatus.Status {
	case "degraded":
		statusCode = http.StatusServiceUnavailable
	case "unhealthy":
		statusCode = http.StatusServiceUnavailable
	}

	core.Success(c, statusCode, "Health check completed", healthStatus, start)
}

func (h *Handler) GetHealthSimple(c *gin.Context) {
	start := time.Now()

	simpleResponse := gin.H{
		"status": "ok",
		"time":   time.Now().UTC(),
	}

	core.Success(c, http.StatusOK, "Simple health check", simpleResponse, start)
}

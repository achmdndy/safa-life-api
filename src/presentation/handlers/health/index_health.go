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

// GetHealth handles GET /health
// @Summary Get health status
// @Description Get detailed health status of the application including database and Redis connections
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=dto.HealthResponse} "Health status retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Failed to get health status"
// @Router /health [get]
func (h *Handler) GetHealth(c *gin.Context) {
	start := time.Now()

	// Create a simple query object
	query := struct{}{}

	result, err := h.healthQueryHandler.Handle(query)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get health status", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	// Convert result to HealthResponse
	healthResponse, ok := result.(dto.HealthResponse)
	if !ok {
		core.Error(c, http.StatusInternalServerError, "Failed to convert health response", &core.ErrorDetail{Reason: "Type assertion failed"}, start)
		return
	}

	// Determine status code based on health status
	statusCode := http.StatusOK
	if healthResponse.Status != "healthy" {
		statusCode = http.StatusServiceUnavailable
	}

	core.Success(c, statusCode, "Health status retrieved successfully", healthResponse, start)
}

// GetHealthSimple handles GET /health/simple
// @Summary Get simple health status
// @Description Get a simple health check response
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse "Application is healthy"
// @Router /health/simple [get]
func (h *Handler) GetHealthSimple(c *gin.Context) {
	start := time.Now()
	core.Success(c, http.StatusOK, "Application is healthy", core.EmptyData{}, start)
}

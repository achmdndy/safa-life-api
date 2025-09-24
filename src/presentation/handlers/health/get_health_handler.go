package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/achmdndy/safa-life-api/src/application/health/dto"
	"github.com/achmdndy/safa-life-api/src/application/health/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

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

	// Create GetHealthQuery
	healthQuery := query.GetHealthQuery{}

	result, err := h.healthQueryHandler.Handle(healthQuery)
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

	// Determine status code and message based on health status
	statusCode := http.StatusOK
	message := "Health status retrieved successfully"
	
	// Check if any service is down or status is not "ok"
	if healthResponse.Status != "ok" {
		statusCode = http.StatusServiceUnavailable
		message = "Service is unhealthy"
	}

	core.Success(c, statusCode, message, healthResponse, start)
}
package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/achmdndy/safa-life-api/src/application/health/dto"
	"github.com/achmdndy/safa-life-api/src/application/health/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// GetHealthRedis handles GET /health/redis
// @Summary Get Redis health status
// @Description Get health status specifically for Redis connection
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=dto.HealthResponse} "Redis health status retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Failed to get Redis health status"
// @Router /health/redis [get]
func (h *Handler) GetHealthRedis(c *gin.Context) {
	start := time.Now()

	// Create GetHealthQuery
	healthQuery := query.GetHealthQuery{}

	result, err := h.healthQueryHandler.Handle(healthQuery)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get Redis health status", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	// Convert result to HealthResponse
	healthResponse, ok := result.(dto.HealthResponse)
	if !ok {
		core.Error(c, http.StatusInternalServerError, "Failed to convert health response", &core.ErrorDetail{Reason: "Type assertion failed"}, start)
		return
	}

	// Filter to only show Redis status
	redisHealthResponse := dto.HealthResponse{
		Status: "ok",
		Services: map[string]string{
			"redis": healthResponse.Services["redis"],
		},
	}

	// Check if Redis is unhealthy
	if redisStatus, exists := healthResponse.Services["redis"]; !exists || redisStatus != "ok" {
		redisHealthResponse.Status = "degraded"
	}

	// Determine status code and message based on Redis health status
	statusCode := http.StatusOK
	message := "Redis health status retrieved successfully"
	
	if redisHealthResponse.Status != "ok" {
		statusCode = http.StatusServiceUnavailable
		message = "Redis is unhealthy"
	}

	core.Success(c, statusCode, message, redisHealthResponse, start)
}
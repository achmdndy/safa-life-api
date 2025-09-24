package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/achmdndy/safa-life-api/src/application/health/dto"
	"github.com/achmdndy/safa-life-api/src/application/health/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

// GetHealthDB handles GET /health/db
// @Summary Get database health status
// @Description Get health status specifically for database connection
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=dto.HealthResponse} "Database health status retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Failed to get database health status"
// @Router /health/db [get]
func (h *Handler) GetHealthDB(c *gin.Context) {
	start := time.Now()

	// Create GetHealthQuery
	healthQuery := query.GetHealthQuery{}

	result, err := h.healthQueryHandler.Handle(healthQuery)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get database health status", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	// Convert result to HealthResponse
	healthResponse, ok := result.(dto.HealthResponse)
	if !ok {
		core.Error(c, http.StatusInternalServerError, "Failed to convert health response", &core.ErrorDetail{Reason: "Type assertion failed"}, start)
		return
	}

	// Filter to only show database status
	dbHealthResponse := dto.HealthResponse{
		Status: "ok",
		Services: map[string]string{
			"database": healthResponse.Services["database"],
		},
	}

	// Check if database is unhealthy
	if dbStatus, exists := healthResponse.Services["database"]; !exists || dbStatus != "ok" {
		dbHealthResponse.Status = "degraded"
	}

	// Determine status code and message based on database health status
	statusCode := http.StatusOK
	message := "Database health status retrieved successfully"
	
	if dbHealthResponse.Status != "ok" {
		statusCode = http.StatusServiceUnavailable
		message = "Database is unhealthy"
	}

	core.Success(c, statusCode, message, dbHealthResponse, start)
}

package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/achmdndy/safa-life-api/src/presentation/core"
)

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
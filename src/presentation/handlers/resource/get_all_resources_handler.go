package resource

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/resource/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAllResources handles GET /resources
// @Summary Get All Resources
// @Description Retrieve a list of all downloadable resources.
// @Tags Resources
// @Accept json
// @Produce json
// @Success 200 {object} core.SuccessResponse{data=[]dto.ResourceResponse} "Resources retrieved successfully"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /resources [get]
func (h *Handler) GetAllResources(c *gin.Context) {
	start := time.Now()
	queryReq := query.ToGetAllResourcesQuery()
	result, err := h.queryHandler.GetAllResources(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get resources", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Resources retrieved successfully", result, start)
}

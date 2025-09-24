package resource

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/resource/dto"
	"github.com/achmdndy/safa-life-api/src/application/resource/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetResourceByID handles GET /resources/:id
// @Summary Get Resource By ID
// @Description Retrieve a specific resource by its ID.
// @Tags Resources
// @Accept json
// @Produce json
// @Param id path string true "Resource ID"
// @Success 200 {object} core.SuccessResponse{data=dto.ResourceResponse} "Resource retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid resource ID"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /resources/{id} [get]
func (h *Handler) GetResourceByID(c *gin.Context) {
	start := time.Now()
	var req dto.GetResourceByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetResourceByIDQuery(req.ID)
	result, err := h.queryHandler.GetResourceByID(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get resource", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Resource retrieved successfully", result, start)
}

package resource

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/resource/command"
	"github.com/achmdndy/safa-life-api/src/application/resource/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateResource handles PUT /resources/:id
// @Summary Update Resource
// @Description Update an existing resource.
// @Tags Resources
// @Accept json
// @Produce json
// @Param id path string true "Resource ID"
// @Param resource body dto.UpdateResourceRequest true "Resource Data"
// @Success 200 {object} core.SuccessResponse{data=dto.ResourceResponse} "Resource updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters or request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /resources/{id} [put]
func (h *Handler) UpdateResource(c *gin.Context) {
	start := time.Now()
	var req dto.UpdateResourceRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromUpdateResourceRequest(req)
	result, err := h.commandHandler.UpdateResource(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to update resource", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Resource updated successfully", dto.ToResourceResponse(*result), start)
}

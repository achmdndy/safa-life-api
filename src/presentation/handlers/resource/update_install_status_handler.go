package resource

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/resource/command"
	"github.com/achmdndy/safa-life-api/src/application/resource/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateInstallStatus handles PATCH /resources/:id/status
// @Summary Update Resource Install Status
// @Description Update the installation status of a resource.
// @Tags Resources
// @Accept json
// @Produce json
// @Param id path string true "Resource ID"
// @Param status body dto.UpdateInstallStatusRequest true "Install Status"
// @Success 200 {object} core.SuccessResponse{data=dto.ResourceResponse} "Resource status updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters or request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /resources/{id}/status [patch]
func (h *Handler) UpdateInstallStatus(c *gin.Context) {
	start := time.Now()
	var req dto.UpdateInstallStatusRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromUpdateInstallStatusRequest(req)
	result, err := h.commandHandler.UpdateInstallStatus(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to update resource status", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Resource status updated successfully", dto.ToResourceResponse(*result), start)
}

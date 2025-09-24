package resource

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/resource/command"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// DeleteResource handles DELETE /resources/:id
// @Summary Delete Resource
// @Description Delete a specific resource.
// @Tags Resources
// @Accept json
// @Produce json
// @Param id path string true "Resource ID"
// @Success 200 {object} core.SuccessResponse "Resource deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid resource ID"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /resources/{id} [delete]
func (h *Handler) DeleteResource(c *gin.Context) {
	start := time.Now()
	id := c.Param("id")

	cmd := command.DeleteResourceCommand{ID: id}
	if err := h.commandHandler.DeleteResource(c.Request.Context(), cmd); err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to delete resource", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Resource deleted successfully", core.EmptyData{}, start)
}

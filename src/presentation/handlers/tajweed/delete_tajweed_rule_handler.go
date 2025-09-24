package tajweed

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tajweed/command"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// DeleteTajweedRule handles DELETE /tajweed/rules/:id
// DeleteTajweedRule handles DELETE /tajweed/rules/:id
// @Summary Delete Tajweed Rule
// @Description Delete a specific tajweed rule from the system.
// @Tags Tajweed Rules
// @Accept json
// @Produce json
// @Param id path string true "Rule ID"
// @Success 200 {object} core.SuccessResponse "Tajweed rule deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid rule ID"
// @Failure 404 {object} core.ErrorResponse "Tajweed rule not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /tajweed/rules/{id} [delete]
func (h *Handler) DeleteTajweedRule(c *gin.Context) {
	start := time.Now()
	id := c.Param("id")

	cmd := command.DeleteTajweedRuleCommand{ID: id}
	err := h.commandHandler.DeleteTajweedRule(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Tajweed rule not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to delete tajweed rule", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Tajweed rule deleted successfully", core.EmptyData{}, start)
}

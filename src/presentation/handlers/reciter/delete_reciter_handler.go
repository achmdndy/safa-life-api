package reciter

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/reciter/command"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// DeleteReciter handles DELETE /reciters/:id
// DeleteReciter handles DELETE /reciters/:id
// @Summary Delete Reciter
// @Description Delete a specific reciter from the system.
// @Tags Reciters
// @Accept json
// @Produce json
// @Param id path string true "Reciter ID"
// @Success 200 {object} core.SuccessResponse "Reciter deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid reciter ID"
// @Failure 404 {object} core.ErrorResponse "Reciter not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /reciters/{id} [delete]
func (h *Handler) DeleteReciter(c *gin.Context) {
	start := time.Now()
	id := c.Param("id")

	cmd := command.DeleteReciterCommand{ID: id}
	if err := h.commandHandler.DeleteReciter(c.Request.Context(), cmd); err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Reciter not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to delete reciter", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Reciter deleted successfully", core.EmptyData{}, start)
}

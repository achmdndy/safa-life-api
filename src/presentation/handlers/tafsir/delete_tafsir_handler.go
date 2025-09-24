package tafsir

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/command"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// DeleteTafsir handles DELETE /tafsirs/editions/:id
// @Summary Delete Tafsir Edition
// @Description Delete a specific tafsir edition.
// @Tags Tafsir Editions
// @Accept json
// @Produce json
// @Param id path string true "Tafsir ID"
// @Success 200 {object} core.SuccessResponse
// @Failure 400 {object} core.ErrorResponse
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/editions/{id} [delete]
func (h *Handler) DeleteTafsir(c *gin.Context) {
	start := time.Now()
	id := c.Param("id")

	cmd := command.DeleteTafsirCommand{ID: id}
	if err := h.commandHandler.DeleteTafsir(c.Request.Context(), cmd); err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to delete tafsir edition", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Tafsir edition deleted successfully", core.EmptyData{}, start)
}

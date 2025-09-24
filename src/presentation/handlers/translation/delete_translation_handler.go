package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/command"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// DeleteTranslation handles DELETE /translations/editions/:id
// DeleteTranslation handles DELETE /translations/editions/:id
// @Summary Delete Translation Edition
// @Description Delete a specific translation edition from the system.
// @Tags Translation Editions
// @Accept json
// @Produce json
// @Param id path string true "Edition ID"
// @Success 200 {object} core.SuccessResponse "Translation edition deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid edition ID"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/editions/{id} [delete]
func (h *Handler) DeleteTranslation(c *gin.Context) {
	start := time.Now()
	id := c.Param("id")

	cmd := command.DeleteTranslationCommand{ID: id}
	if err := h.commandHandler.DeleteTranslation(c.Request.Context(), cmd); err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to delete translation edition", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Translation edition deleted successfully", core.EmptyData{}, start)
}

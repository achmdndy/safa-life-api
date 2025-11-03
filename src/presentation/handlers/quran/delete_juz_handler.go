package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// DeleteJuzHandler handles the delete juz request
type DeleteJuzHandler struct {
	commandHandler *command.CommandHandler
}

// NewDeleteJuzHandler creates a new delete juz handler
func NewDeleteJuzHandler(commandHandler *command.CommandHandler) *DeleteJuzHandler {
	return &DeleteJuzHandler{
		commandHandler: commandHandler,
	}
}

// Handle processes the delete juz request
// @Summary Delete a juz
// @Description Delete an existing juz by ID
// @Tags Juz
// @Accept json
// @Produce json
// @Param id path string true "Juz ID"
// @Success 200 {object} JuzSuccessResponse "Juz deleted successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Juz not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /api/v1/quran/juz/{id} [delete]
func (h *DeleteJuzHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.DeleteJuzRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusBadRequest, "Invalid juz ID", errorDetail, start)
		return
	}

	cmd := command.DeleteJuzCommand{
		ID: req.ID,
	}

	err := h.commandHandler.DeleteJuz(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{
			Reason: err.Error(),
		}
		core.Error(c, http.StatusInternalServerError, "Failed to delete juz", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Juz deleted successfully", map[string]interface{}{}, start)
}
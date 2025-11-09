package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// DeleteLastReadHandler handles deleting a last read by ID
type DeleteLastReadHandler struct {
	commandHandler *command.CommandHandler
}

// NewDeleteLastReadHandler creates a new handler instance
func NewDeleteLastReadHandler(commandHandler *command.CommandHandler) *DeleteLastReadHandler {
	return &DeleteLastReadHandler{commandHandler: commandHandler}
}

// Handle processes the request
// @Summary Delete last read
// @Description Delete a last read entry by its ID
// @Tags LastRead
// @Accept json
// @Produce json
// @Param id path string true "LastRead ID"
// @Success 200 {object} DeleteLastReadSuccessResponse "Last read deleted successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Last read not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/last-reads/{id} [delete]
func (h *DeleteLastReadHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.DeleteLastReadRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid last read ID", errorDetail, start)
		return
	}

	cmd := command.DeleteLastReadCommand{ID: req.ID}

	err := h.commandHandler.DeleteLastRead(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to delete last read", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Last read deleted successfully", map[string]interface{}{}, start)
}

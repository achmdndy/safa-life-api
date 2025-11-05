package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// DeleteReciterHandler handles the delete reciter request
type DeleteReciterHandler struct {
	commandHandler *command.CommandHandler
}

// NewDeleteReciterHandler creates a new delete reciter handler
func NewDeleteReciterHandler(commandHandler *command.CommandHandler) *DeleteReciterHandler {
	return &DeleteReciterHandler{commandHandler: commandHandler}
}

// Handle processes the delete reciter request
// @Summary Delete a reciter
// @Description Delete a reciter by ID
// @Tags Reciters
// @Accept json
// @Produce json
// @Param id path string true "Reciter ID"
// @Success 200 {object} DeleteReciterSuccessResponse "Reciter deleted successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/reciters/{id} [delete]
func (h *DeleteReciterHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.DeleteReciterRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", errorDetail, start)
		return
	}

	cmd := command.DeleteReciterCommand{ID: req.ID}
	if err := h.commandHandler.DeleteReciter(ctx, cmd); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to delete reciter", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Reciter deleted successfully", map[string]interface{}{}, start)
}

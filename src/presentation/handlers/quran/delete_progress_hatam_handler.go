package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// DeleteProgressHatamHandler handles deleting a progress hatam by ID
type DeleteProgressHatamHandler struct {
	commandHandler *command.CommandHandler
}

// NewDeleteProgressHatamHandler creates a new handler instance
func NewDeleteProgressHatamHandler(commandHandler *command.CommandHandler) *DeleteProgressHatamHandler {
	return &DeleteProgressHatamHandler{commandHandler: commandHandler}
}

// Handle processes the request
// @Summary Delete progress hatam
// @Description Delete a progress hatam entry by its ID
// @Tags ProgressHatam
// @Accept json
// @Produce json
// @Param id path string true "ProgressHatam ID"
// @Success 200 {object} DeleteProgressHatamSuccessResponse "Progress hatam deleted successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 404 {object} QuranErrorResponse "Progress hatam not found"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/progress-hatam/{id} [delete]
func (h *DeleteProgressHatamHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.DeleteProgressHatamRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid progress hatam ID", errorDetail, start)
		return
	}

	cmd := command.DeleteProgressHatamCommand{ID: req.ID}

	err := h.commandHandler.DeleteProgressHatam(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to delete progress hatam", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Progress hatam deleted successfully", map[string]interface{}{}, start)
}

package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// DeleteAyahAudioFileHandler handles the delete ayah audio file request
type DeleteAyahAudioFileHandler struct {
	commandHandler *command.CommandHandler
}

// NewDeleteAyahAudioFileHandler creates a new delete ayah audio file handler
func NewDeleteAyahAudioFileHandler(commandHandler *command.CommandHandler) *DeleteAyahAudioFileHandler {
	return &DeleteAyahAudioFileHandler{commandHandler: commandHandler}
}

// Handle processes the delete ayah audio file request
// @Summary Delete an ayah audio file
// @Description Delete an ayah audio file by ID
// @Tags Audio
// @Accept json
// @Produce json
// @Param id path string true "Ayah audio file ID"
// @Success 200 {object} DeleteAyahAudioFileSuccessResponse "Ayah audio file deleted successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/audio/ayahs/{id} [delete]
func (h *DeleteAyahAudioFileHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.DeleteAyahAudioFileRequest
	if err := c.ShouldBindUri(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", errorDetail, start)
		return
	}

	cmd := command.DeleteAyahAudioFileCommand{ID: req.ID}
	if err := h.commandHandler.DeleteAyahAudioFile(ctx, cmd); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to delete ayah audio file", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah audio file deleted successfully", map[string]interface{}{}, start)
}

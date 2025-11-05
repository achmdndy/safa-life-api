package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
)

// UpdateAyahAudioFileHandler handles the update ayah audio file request
type UpdateAyahAudioFileHandler struct {
	commandHandler *command.CommandHandler
}

// NewUpdateAyahAudioFileHandler creates a new update ayah audio file handler
func NewUpdateAyahAudioFileHandler(commandHandler *command.CommandHandler) *UpdateAyahAudioFileHandler {
	return &UpdateAyahAudioFileHandler{commandHandler: commandHandler}
}

// Handle processes the update ayah audio file request
// @Summary Update an ayah audio file
// @Description Update an existing ayah audio file
// @Tags Audio
// @Accept json
// @Produce json
// @Param id path string true "Ayah audio file ID"
// @Param request body dto.UpdateAyahAudioFileRequest true "Update ayah audio file request"
// @Success 200 {object} UpdateAyahAudioFileSuccessResponse "Ayah audio file updated successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/audio/ayahs/{id} [put]
func (h *UpdateAyahAudioFileHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var uri dto.UpdateAyahAudioFileRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", errorDetail, start)
		return
	}

	var body dto.UpdateAyahAudioFileRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.UpdateAyahAudioFileCommand{
		ID:        uri.ID,
		FilePath:  body.FilePath,
		Duration:  body.Duration,
		ByteSize:  body.ByteSize,
		UpdatedBy: body.UpdatedBy,
	}

	result, err := h.commandHandler.UpdateAyahAudioFile(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to update ayah audio file", errorDetail, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah audio file updated successfully", result, start)
}

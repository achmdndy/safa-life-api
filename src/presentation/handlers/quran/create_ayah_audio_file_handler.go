package quran

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/safalife/core-api/src/application/quran/command"
	"github.com/safalife/core-api/src/application/quran/dto"
	"github.com/safalife/core-api/src/presentation/core"
	"github.com/safalife/core-api/src/presentation/middlewares"
)

// CreateAyahAudioFileHandler handles the create ayah audio file request
type CreateAyahAudioFileHandler struct {
	commandHandler *command.CommandHandler
}

// NewCreateAyahAudioFileHandler creates a new create ayah audio file handler
func NewCreateAyahAudioFileHandler(commandHandler *command.CommandHandler) *CreateAyahAudioFileHandler {
	return &CreateAyahAudioFileHandler{commandHandler: commandHandler}
}

// Handle processes the create ayah audio file request
// @Summary Create a new ayah audio file
// @Description Create a new ayah audio file
// @Tags Audio
// @Accept json
// @Produce json
// @Param request body dto.CreateAyahAudioFileRequest true "Create ayah audio file request"
// @Success 201 {object} CreateAyahAudioFileSuccessResponse "Ayah audio file created successfully"
// @Failure 400 {object} QuranErrorResponse "Bad request"
// @Failure 500 {object} QuranErrorResponse "Internal server error"
// @Router /quran/audio/ayahs [post]
func (h *CreateAyahAudioFileHandler) Handle(c *gin.Context) {
	start := time.Now()
	ctx := c.Request.Context()

	var req dto.CreateAyahAudioFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusBadRequest, "Invalid request body", errorDetail, start)
		return
	}

	cmd := command.CreateAyahAudioFileCommand{
		ReciterID: req.ReciterID,
		SurahID:   req.SurahID,
		AyahID:    req.AyahID,
		FilePath:  req.FilePath,
		Duration:  req.Duration,
		ByteSize:  req.ByteSize,
		CreatedBy: middlewares.GetUserID(c),
	}

	result, err := h.commandHandler.CreateAyahAudioFile(ctx, cmd)
	if err != nil {
		errorDetail := &core.ErrorDetail{Reason: err.Error()}
		core.Error(c, http.StatusInternalServerError, "Failed to create ayah audio file", errorDetail, start)
		return
	}

	core.Success(c, http.StatusCreated, "Ayah audio file created successfully", result, start)
}

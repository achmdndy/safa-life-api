package audio

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/audio/command"
	"github.com/achmdndy/safa-life-api/src/application/audio/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// CreateAyahAudio handles POST /audio/ayahs
// @Summary Create Ayah Audio
// @Description Add a new audio file for a specific ayah.
// @Tags Audio
// @Accept json
// @Produce json
// @Param ayah_audio body dto.CreateAyahAudioRequest true "Ayah Audio Data"
// @Success 201 {object} core.SuccessResponse{data=dto.AyahAudioResponse} "Ayah audio created successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /audio/ayahs [post]
func (h *Handler) CreateAyahAudio(c *gin.Context) {
	start := time.Now()
	var req dto.CreateAyahAudioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromCreateAyahAudioRequest(req)
	result, err := h.commandHandler.CreateAyahAudio(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to create ayah audio", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusCreated, "Ayah audio created successfully", dto.ToAyahAudioResponse(*result), start)
}
package audio

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/audio/command"
	"github.com/achmdndy/safa-life-api/src/application/audio/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateAyahAudio handles PUT /audio/ayahs/:reciter_id/:surah_id/:ayah_id
// @Summary Update Ayah Audio
// @Description Update an existing ayah audio file's details.
// @Tags Audio
// @Accept json
// @Produce json
// @Param reciter_id path string true "Reciter ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Param ayah_audio body dto.UpdateAyahAudioRequest true "Ayah Audio Data"
// @Success 200 {object} core.SuccessResponse{data=dto.AyahAudioResponse} "Ayah audio updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters or request body"
// @Failure 404 {object} core.ErrorResponse "Ayah audio not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /audio/ayahs/{reciter_id}/{surah_id}/{ayah_id} [put]
func (h *Handler) UpdateAyahAudio(c *gin.Context) {
	start := time.Now()
	var req dto.UpdateAyahAudioRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromUpdateAyahAudioRequest(req)
	result, err := h.commandHandler.UpdateAyahAudio(c.Request.Context(), cmd)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Ayah audio not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to update ayah audio", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Ayah audio not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah audio updated successfully", dto.ToAyahAudioResponse(*result), start)
}

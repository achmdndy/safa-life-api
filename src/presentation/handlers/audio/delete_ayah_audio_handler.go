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

// DeleteAyahAudio handles DELETE /audio/ayahs/:reciter_id/:surah_id/:ayah_id
// @Summary Delete Ayah Audio
// @Description Delete a specific ayah audio file.
// @Tags Audio
// @Accept json
// @Produce json
// @Param reciter_id path string true "Reciter ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse "Ayah audio deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 404 {object} core.ErrorResponse "Ayah audio not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /audio/ayahs/{reciter_id}/{surah_id}/{ayah_id} [delete]
func (h *Handler) DeleteAyahAudio(c *gin.Context) {
	start := time.Now()
	var req dto.GetAyahAudioRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.DeleteAyahAudioCommand(req)
	if err := h.commandHandler.DeleteAyahAudio(c.Request.Context(), cmd); err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Ayah audio not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to delete ayah audio", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah audio deleted successfully", core.EmptyData{}, start)
}

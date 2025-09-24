package audio

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/audio/dto"
	"github.com/achmdndy/safa-life-api/src/application/audio/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAyahAudio handles GET /audio/ayahs/:reciter_id/:surah_id/:ayah_id
// @Summary Get Ayah Audio
// @Description Retrieve a specific ayah audio file.
// @Tags Audio
// @Accept json
// @Produce json
// @Param reciter_id path string true "Reciter ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse{data=dto.AyahAudioResponse} "Ayah audio retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 404 {object} core.ErrorResponse "Ayah audio not found"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /audio/ayahs/{reciter_id}/{surah_id}/{ayah_id} [get]
func (h *Handler) GetAyahAudio(c *gin.Context) {
	start := time.Now()
	var req dto.GetAyahAudioRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetAyahAudioQuery(req.ReciterID, req.SurahID, req.AyahID)
	result, err := h.queryHandler.GetAyahAudio(c.Request.Context(), queryReq)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Ayah audio not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah audio", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Ayah audio not found", &core.ErrorDetail{Reason: "Ayah audio not found"}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah audio retrieved successfully", result, start)
}

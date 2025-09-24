package audio

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/audio/dto"
	"github.com/achmdndy/safa-life-api/src/application/audio/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAudioForSurah handles GET /audio/surahs/:reciter_id/:surah_id
// @Summary Get Audio Files for a Surah
// @Description Retrieve all audio files for a specific surah by a specific reciter.
// @Tags Audio
// @Accept json
// @Produce json
// @Param reciter_id path string true "Reciter ID"
// @Param surah_id path int true "Surah ID"
// @Success 200 {object} core.SuccessResponse{data=[]dto.AyahAudioResponse} "Audio files retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /audio/surahs/{reciter_id}/{surah_id} [get]
func (h *Handler) GetAudioForSurah(c *gin.Context) {
	start := time.Now()
	var req dto.GetAudioForSurahRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetAudioForSurahQuery(req.ReciterID, req.SurahID)
	result, err := h.queryHandler.GetAudioForSurah(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get audio files", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Audio files retrieved successfully", result, start)
}
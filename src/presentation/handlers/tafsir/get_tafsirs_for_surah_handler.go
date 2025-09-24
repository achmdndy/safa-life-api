package tafsir

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
	"github.com/achmdndy/safa-life-api/src/application/tafsir/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetTafsirsForSurah handles GET /tafsirs/ayahs/by-surah/:tafsir_id/:surah_id
// @Summary Get Tafsirs for Surah
// @Description Get all tafsirs for a specific surah.
// @Tags Ayah Tafsirs
// @Accept json
// @Produce json
// @Param tafsir_id path string true "Tafsir ID"
// @Param surah_id path int true "Surah ID"
// @Success 200 {object} core.SuccessResponse{data=[]dto.AyahTafsirResponse}
// @Failure 400 {object} core.ErrorResponse
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/ayahs/by-surah/{tafsir_id}/{surah_id} [get]
func (h *Handler) GetTafsirsForSurah(c *gin.Context) {
	start := time.Now()
	var req dto.GetTafsirsForSurahRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetTafsirsForSurahQuery(req.TafsirID, req.SurahID)
	result, err := h.queryHandler.GetTafsirsForSurah(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get tafsirs for surah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Tafsirs for surah retrieved successfully", result, start)
}

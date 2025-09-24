package tafsir

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
	"github.com/achmdndy/safa-life-api/src/application/tafsir/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetTafsirsForAyah handles GET /tafsirs/ayahs/by-ayah/:surah_id/:ayah_id
// @Summary Get Tafsirs for Ayah
// @Description Get all tafsirs for a specific ayah.
// @Tags Ayah Tafsirs
// @Accept json
// @Produce json
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse{data=[]dto.AyahTafsirResponse}
// @Failure 400 {object} core.ErrorResponse
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/ayahs/by-ayah/{surah_id}/{ayah_id} [get]
func (h *Handler) GetTafsirsForAyah(c *gin.Context) {
	start := time.Now()
	var req dto.GetTafsirsForAyahRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetTafsirsForAyahQuery(req.SurahID, req.AyahID)
	result, err := h.queryHandler.GetTafsirsForAyah(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get tafsirs for ayah", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Tafsirs for ayah retrieved successfully", result, start)
}

package tafsir

import (
	"net/http"
	"strings"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/tafsir/dto"
	"github.com/achmdndy/safa-life-api/src/application/tafsir/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAyahTafsir handles GET /tafsirs/ayahs/:tafsir_id/:surah_id/:ayah_id
// @Summary Get Ayah Tafsir
// @Description Get a specific ayah tafsir.
// @Tags Ayah Tafsirs
// @Accept json
// @Produce json
// @Param tafsir_id path string true "Tafsir ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse{data=dto.AyahTafsirResponse}
// @Failure 400 {object} core.ErrorResponse
// @Failure 404 {object} core.ErrorResponse "Tafsir not found"
// @Failure 500 {object} core.ErrorResponse
// @Router /tafsirs/ayahs/{tafsir_id}/{surah_id}/{ayah_id} [get]
func (h *Handler) GetAyahTafsir(c *gin.Context) {
	start := time.Now()
	var req dto.GetAyahTafsirRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetAyahTafsirQuery(req.TafsirID, req.SurahID, req.AyahID)
	result, err := h.queryHandler.GetAyahTafsir(c.Request.Context(), queryReq)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			core.Error(c, http.StatusNotFound, "Tafsir not found", &core.ErrorDetail{Reason: err.Error()}, start)
			return
		}
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah tafsir", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	if result == nil {
		core.Error(c, http.StatusNotFound, "Tafsir not found", &core.ErrorDetail{}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah tafsir retrieved successfully", result, start)
}

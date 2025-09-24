package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
	"github.com/achmdndy/safa-life-api/src/application/translation/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetAyahTranslation handles GET /translations/ayahs/:translation_id/:surah_id/:ayah_id
// GetAyahTranslation handles GET /translations/ayahs/:translation_id/:surah_id/:ayah_id
// @Summary Get Ayah Translation
// @Description Retrieve a specific ayah translation.
// @Tags Ayah Translations
// @Accept json
// @Produce json
// @Param translation_id path string true "Translation ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse{data=dto.AyahTranslationResponse} "Ayah translation retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/ayahs/{translation_id}/{surah_id}/{ayah_id} [get]
func (h *Handler) GetAyahTranslation(c *gin.Context) {
	start := time.Now()
	var req dto.GetAyahTranslationRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetAyahTranslationQuery(req.TranslationID, req.SurahID, req.AyahID)
	result, err := h.queryHandler.GetAyahTranslation(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah translation", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah translation retrieved successfully", result, start)
}

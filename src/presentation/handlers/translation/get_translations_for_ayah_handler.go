package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
	"github.com/achmdndy/safa-life-api/src/application/translation/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetTranslationsForAyah handles GET /translations/ayahs/:surah_id/:ayah_id
// GetTranslationsForAyah handles GET /translations/ayahs/by-ayah/:surah_id/:ayah_id
// @Summary Get All Translations for an Ayah
// @Description Retrieve all available translations for a specific ayah.
// @Tags Ayah Translations
// @Accept json
// @Produce json
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse{data=[]dto.AyahTranslationResponse} "Ayah translations retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/ayahs/by-ayah/{surah_id}/{ayah_id} [get]
func (h *Handler) GetTranslationsForAyah(c *gin.Context) {
	start := time.Now()
	var req dto.GetTranslationsForAyahRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetTranslationsForAyahQuery(req.SurahID, req.AyahID)
	result, err := h.queryHandler.GetTranslationsForAyah(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get ayah translations", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah translations retrieved successfully", result, start)
}

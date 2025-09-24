package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
	"github.com/achmdndy/safa-life-api/src/application/translation/query"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// GetTranslationsForSurah handles GET /translations/surahs/:translation_id/:surah_id
// GetTranslationsForSurah handles GET /translations/ayahs/by-surah/:translation_id/:surah_id
// @Summary Get All Translations for a Surah
// @Description Retrieve all ayah translations for a specific surah and translation edition.
// @Tags Ayah Translations
// @Accept json
// @Produce json
// @Param translation_id path string true "Translation ID"
// @Param surah_id path int true "Surah ID"
// @Success 200 {object} core.SuccessResponse{data=[]dto.AyahTranslationResponse} "Surah translations retrieved successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/ayahs/by-surah/{translation_id}/{surah_id} [get]
func (h *Handler) GetTranslationsForSurah(c *gin.Context) {
	start := time.Now()
	var req dto.GetTranslationsForSurahRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	queryReq := query.ToGetTranslationsForSurahQuery(req.TranslationID, req.SurahID)
	result, err := h.queryHandler.GetTranslationsForSurah(c.Request.Context(), queryReq)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to get surah translations", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Surah translations retrieved successfully", result, start)
}

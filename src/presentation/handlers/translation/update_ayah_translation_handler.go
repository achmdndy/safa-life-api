package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/command"
	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// UpdateAyahTranslation handles PUT /translations/ayahs/:translation_id/:surah_id/:ayah_id
// UpdateAyahTranslation handles PUT /translations/ayahs/:translation_id/:surah_id/:ayah_id
// @Summary Update Ayah Translation
// @Description Update an existing ayah translation.
// @Tags Ayah Translations
// @Accept json
// @Produce json
// @Param translation_id path string true "Translation ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Param ayah_translation body dto.UpdateAyahTranslationRequest true "Ayah Translation Data"
// @Success 200 {object} core.SuccessResponse{data=dto.AyahTranslationResponse} "Ayah translation updated successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters or request body"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/ayahs/{translation_id}/{surah_id}/{ayah_id} [put]
func (h *Handler) UpdateAyahTranslation(c *gin.Context) {
	start := time.Now()
	var req dto.UpdateAyahTranslationRequest

	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid URI parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request body", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.FromUpdateAyahTranslationRequest(req)
	result, err := h.commandHandler.UpdateAyahTranslation(c.Request.Context(), cmd)
	if err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to update ayah translation", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah translation updated successfully", dto.ToAyahTranslationResponse(*result), start)
}

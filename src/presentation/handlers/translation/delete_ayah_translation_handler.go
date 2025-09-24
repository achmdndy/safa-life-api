package translation

import (
	"net/http"
	"time"

	"github.com/achmdndy/safa-life-api/src/application/translation/command"
	"github.com/achmdndy/safa-life-api/src/application/translation/dto"
	"github.com/achmdndy/safa-life-api/src/presentation/core"
	"github.com/gin-gonic/gin"
)

// DeleteAyahTranslation handles DELETE /translations/ayahs/:translation_id/:surah_id/:ayah_id
// DeleteAyahTranslation handles DELETE /translations/ayahs/:translation_id/:surah_id/:ayah_id
// @Summary Delete Ayah Translation
// @Description Delete a specific ayah translation.
// @Tags Ayah Translations
// @Accept json
// @Produce json
// @Param translation_id path string true "Translation ID"
// @Param surah_id path int true "Surah ID"
// @Param ayah_id path int true "Ayah ID"
// @Success 200 {object} core.SuccessResponse "Ayah translation deleted successfully"
// @Failure 400 {object} core.ErrorResponse "Invalid parameters"
// @Failure 500 {object} core.ErrorResponse "Internal server error"
// @Router /translations/ayahs/{translation_id}/{surah_id}/{ayah_id} [delete]
func (h *Handler) DeleteAyahTranslation(c *gin.Context) {
	start := time.Now()
	var req dto.GetAyahTranslationRequest
	if err := c.ShouldBindUri(&req); err != nil {
		core.Error(c, http.StatusBadRequest, "Invalid request parameters", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	cmd := command.DeleteAyahTranslationCommand(req)
	if err := h.commandHandler.DeleteAyahTranslation(c.Request.Context(), cmd); err != nil {
		core.Error(c, http.StatusInternalServerError, "Failed to delete ayah translation", &core.ErrorDetail{Reason: err.Error()}, start)
		return
	}

	core.Success(c, http.StatusOK, "Ayah translation deleted successfully", core.EmptyData{}, start)
}
